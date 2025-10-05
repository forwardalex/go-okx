package main

import (
	"log"

	"github.com/forwardalex/go-okx/examples/rest"
	"github.com/forwardalex/go-okx/rest/api"
	"github.com/forwardalex/go-okx/rest/api/trade"
)

func main() {
	param := &trade.GetFillsParam{
		InstType: api.InstTypeSPOT,
	}
	req, resp := trade.NewGetFills(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.GetFillsResponse))
}
