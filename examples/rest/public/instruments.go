package main

import (
	"log"

	"github.com/starchou/go-okx/examples/rest"
	"github.com/starchou/go-okx/rest/api"
	"github.com/starchou/go-okx/rest/api/public"
)

func main() {
	param := &public.GetInstrumentsParam{
		InstType: api.InstTypeSPOT,
	}
	req, resp := public.NewGetInstruments(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*public.GetInstrumentsResponse))
}
