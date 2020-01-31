package server

import (
	"context"
	"github.com/pepeunlimited/apple-iap/applerpc"
	"github.com/pepeunlimited/apple-iap/internal/app/app1/apple"
	"github.com/pepeunlimited/apple-iap/internal/app/app1/validator"
	validator2 "github.com/pepeunlimited/microservice-kit/validator"
	"log"
)

type AppleIAPServer struct {
	validator 			validator.AppleIAPServerValidator
	appstore			apple.AppStore
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
		Status:         applerpc.VerifyReceiptResponse_OK,
		Type:           applerpc.VerifyReceiptResponse_CONSUMABLE,
		AppleProductId: "AppleProductID",
	}, nil
}

func (server AppleIAPServer) password(params *applerpc.VerifyReceiptParams) *string {
	if params.Password == nil || validator2.IsEmpty(params.Password.Value) {
		return nil
	}
	return &params.Password.Value
}

func NewAppleIAPServer(appstore apple.AppStore) AppleIAPServer {
	return AppleIAPServer{
		validator: 			validator.NewAppleIAPServerValidator(),
		appstore:			appstore,
	}
}