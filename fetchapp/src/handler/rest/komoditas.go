package restserver

import (
	"net/http"
)

func (e *rest) GetKomoditas(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authInfo, err := e.auth.GetAuthInfo(r.Header.Get("Authorization"))
	if err != nil {
		if err != nil {
			Failed(err.Error(), http.StatusUnauthorized, w)
			return
		}
	}
	if authInfo.Role == "admin" {
		println(authInfo.Role)
	}
	results, err := e.uc.Komoditas.GetKomoditas(ctx)
	if err != nil {
		Failed(err.Error(), http.StatusInternalServerError, w)
		return
	}

	Success(&Meta{Count: len(results), Status: http.StatusOK, Message: "Get Komoditas"}, results, w)
}

// GetAggregate godoc
// @Summary Get Aggregate
// @Description Get Aggregate
// @Tags komoditas
// @Accept json
// @Produce json
// @Success 200 {object} restserver.Response
// @Failure 400 {object} restserver.Response
// @Failure 401 {object} restserver.Response
// @Failure 500 {object} restserver.Response
// @Router /aggregate [get]
func (e *rest) GetAggregate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authInfo, err := e.auth.GetAuthInfo(r.Header.Get("Authorization"))
	if err != nil {
		if err != nil {
			Failed(err.Error(), http.StatusUnauthorized, w)
			return
		}
	}
	if authInfo.Role != "admin" {
		Failed("Unautorized", http.StatusUnauthorized, w)
		return
	}

	results, err := e.uc.Komoditas.GetAggregate(ctx)
	if err != nil {
		Failed(err.Error(), http.StatusInternalServerError, w)
		return
	}

	Success(&Meta{Count: len(results), Status: http.StatusOK, Message: "Get Aggregate"}, results, w)
}
