package restserver

import (
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	_ "github.com/taufiq-mohamad-41186/test-project/fetchapp/docs"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/usecase"
	"log"
	"net/http"
	"sync"
	"time"
)

var once = &sync.Once{}

type REST interface{}

type rest struct {
	auth   *Auth
	router *mux.Router
	uc     *usecase.Usecase
}

func Init(auth *Auth, router *mux.Router, uc *usecase.Usecase) REST {
	var e *rest
	once.Do(func() {
		e = &rest{
			auth:   auth,
			router: router,
			uc:     uc,
		}
		e.Serve()
	})
	return e
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func (e *rest) Serve() {
	e.router.HandleFunc("/komoditas", e.GetKomoditas).Methods("GET")
	e.router.HandleFunc("/aggregate", e.GetAggregate).Methods("GET")
	e.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://127.0.0.1:8000/swagger/index.html"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	)).Methods("GET")

	srv := &http.Server{
		Handler:      e.router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
