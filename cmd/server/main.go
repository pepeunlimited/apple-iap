package main

import (
	"github.com/pepeunlimited/apple-iap/internal/server/twrip"
	"github.com/pepeunlimited/apple-iap/pkg/appleiap"
	"github.com/pepeunlimited/apple-iap/pkg/applerpc"
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

	ts := applerpc.NewAppleIAPServiceServer(twrip.NewAppleIAPServer(appleiap.NewAppStoreByMode(misc.GetEnv(appleiap.IapMode, "MOCK"))), nil)

	mux := http.NewServeMux()
	mux.Handle(ts.PathPrefix(), middleware.Adapt(ts, headers.Username()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}