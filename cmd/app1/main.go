package main

import (
	"github.com/pepeunlimited/apple-iap/apple"
	"github.com/pepeunlimited/apple-iap/applerpc"
	"github.com/pepeunlimited/apple-iap/internal/app/app1/server"
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

	ts := applerpc.NewAppleIAPServiceServer(server.NewAppleIAPServer(apple.NewAppStoreByMode(misc.GetEnv(apple.IapMode, "MOCK"))), nil)

	mux := http.NewServeMux()
	mux.Handle(ts.PathPrefix(), middleware.Adapt(ts, headers.Username()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}