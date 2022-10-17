package entity

import "time"

type HTTPKomoditasResp struct {
	UUID         *string `json:"uuid"`
	Komoditas    *string `json:"komoditas"`
	AreaProvinsi *string `json:"area_provinsi"`
	AreaKota     *string `json:"area_kota"`
	Size         *string `json:"size"`
	Price        *string `json:"price"`
	TglParsed    *string `json:"tgl_parsed"`
	TimeStamp    *string `json:"timestamp"`
}

type Komoditas struct {
	UUID         string    `json:"uuid"`
	Komoditas    string    `json:"komoditas"`
	AreaProvinsi string    `json:"area_provinsi"`
	AreaKota     string    `json:"area_kota"`
	Size         int       `json:"size"`
	Price        []Price   `json:"price,omitempty"`
	TglParsed    time.Time `json:"tgl_parsed"`
	TimeStamp    time.Time `json:"timestamp"`
}

type Price struct {
	UOM   string  `json:"uom"`
	Value float64 `json:"value"`
}

type AggregateKey struct {
	AreaProvinsi string    `json:"area_provinsi"`
	Week         int       `json:"week"`
	Aggregate    Aggregate `json:"aggregate"`
}

type Aggregate struct {
	Min    AggregateData `json:"min"`
	Max    AggregateData `json:"max"`
	Median AggregateData `json:"median"`
	Avg    AggregateData `json:"avg"`
}

type AggregateData struct {
	Size  float64 `json:"size"`
	Price float64 `json:"price"`
}

type MapAggregate struct {
	Size  []int
	Price [][]Price
}
