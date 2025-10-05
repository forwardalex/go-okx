package main

import (
	"log"

	"github.com/forwardalex/go-okx/examples/rest"
	"github.com/forwardalex/go-okx/rest/api/market"
)

func main() {
	req, resp := market.NewGetExchangeRate()
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetExchangeRateResponse))
}
