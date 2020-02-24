package main

import (
	"github.com/pepeunlimited/apple-iap/internal/server/twrip"
	"github.com/pepeunlimited/apple-iap/pkg/appstore"
	"github.com/pepeunlimited/apple-iap/pkg/rpc/appleiap"
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/microservice-kit/misc"
	"log"
	"net/http"
)

const (
	Version = "0.0.3"
)

func main() {
	log.Printf("Starting the AppleIAPServer... version=[%v]", Version)

	ts := appleiap.NewAppleIAPServiceServer(twrip.NewAppleIAPServer(appstore.NewAppStoreByMode(misc.GetEnv(appstore.IapMode, "MOCK"))), nil)

	mux := http.NewServeMux()
	mux.Handle(ts.PathPrefix(), middleware.Adapt(ts, headers.Username()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}