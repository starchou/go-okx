package main

import (
	"log"

	"github.com/starchou/go-okx/examples/rest"
	"github.com/starchou/go-okx/rest/api"
	"github.com/starchou/go-okx/rest/api/trade"
)

func main() {
	param := &trade.GetOrdersQueryParam{
		InstType: api.InstTypeSPOT,
	}
	req, resp := trade.NewGetOrdersHistory(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.GetOrderResponse))
}
