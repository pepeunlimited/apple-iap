package twrip

import (
	"context"
	"github.com/pepeunlimited/apple-iap/internal/server/validator"
	"github.com/pepeunlimited/apple-iap/pkg/appstore"
	"github.com/pepeunlimited/apple-iap/pkg/rpc/appleiap"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"log"
)

type AppleIAPServer struct {
	validator validator.AppleIAPServerValidator
	appstore  appstore.AppStore
}

func (server AppleIAPServer) VerifyReceipt(ctx context.Context, params *appleiap.VerifyReceiptParams) (*appleiap.VerifyReceiptResponse, error) {
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
	return &appleiap.VerifyReceiptResponse{
		Status:         "OK",			// TODO: enum?
		Type:           "CONSUMABLE",	// TODO: enum?
		AppleProductId: "AppleProductID",
	}, nil
}

func (server AppleIAPServer) password(params *appleiap.VerifyReceiptParams) *string {
	if params.Password == nil || validator2.IsEmpty(params.Password.Value) {
		return nil
	}
	return &params.Password.Value
}

func NewAppleIAPServer(appstore appstore.AppStore) AppleIAPServer {
	return AppleIAPServer{
		validator: validator.NewAppleIAPServerValidator(),
		appstore:  appstore,
	}
}