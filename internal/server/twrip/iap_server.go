package twrip

import (
	"context"
	"github.com/pepeunlimited/apple-iap/internal/server/validator"
	"github.com/pepeunlimited/apple-iap/pkg/appleiap"
	"github.com/pepeunlimited/apple-iap/pkg/applerpc"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"log"
)

type AppleIAPServer struct {
	validator validator.AppleIAPServerValidator
	appstore  appleiap.AppStore
}

func (server AppleIAPServer) VerifyReceipt(ctx context.Context, params *applerpc.VerifyReceiptParams) (*applerpc.VerifyReceiptResponse, error) {
	err := server.validator.VerifyReceipt(params)
	if err != nil {
		return nil, err
	}
	password := server.password(params)
	verified, err := server.appstore.VerifyReceipt(ctx, params.Receipt, password)
	if err != nil {
		log.Print("apple-iap: issue during verify receipt: "+err.Error())
		return nil, err
	}
	log.Print(verified)
	return &applerpc.VerifyReceiptResponse{
		Status:         "OK",			// TODO: enum?
		Type:           "CONSUMABLE",	// TODO: enum?
		AppleProductId: "AppleProductID",
	}, nil
}

func (server AppleIAPServer) password(params *applerpc.VerifyReceiptParams) *string {
	if params.Password == nil || validator2.IsEmpty(params.Password.Value) {
		return nil
	}
	return &params.Password.Value
}

func NewAppleIAPServer(appstore appleiap.AppStore) AppleIAPServer {
	return AppleIAPServer{
		validator: validator.NewAppleIAPServerValidator(),
		appstore:  appstore,
	}
}