package main

import (
	"log"

	"github.com/forwardalex/go-okx/examples/rest"
	"github.com/forwardalex/go-okx/rest/api"
	"github.com/forwardalex/go-okx/rest/api/account"
)

func main() {
	param := &account.GetMaxSizeParam{
		InstId: "BTC-USDT",
		TdMode: api.TdModeCross,
	}
	req, resp := account.NewGetMaxSize(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetMaxSizeResponse))
}
