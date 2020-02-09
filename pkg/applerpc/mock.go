package applerpc

import (
	"context"
	"github.com/pepeunlimited/apple-iap/pkg/appleiap"
	"log"
)

type AppleIAPMock struct {
	mock      appleiap.AppStoreMock
	productId string
}

func (a AppleIAPMock) VerifyReceipt(ctx context.Context, params *VerifyReceiptParams) (*VerifyReceiptResponse, error) {
	receipt, err := a.mock.VerifyReceipt(ctx, params.Receipt, nil)
	if err != nil {
		return nil, err
	}
	log.Print(receipt)
	return &VerifyReceiptResponse{
		Status:         "OK",
		Type:           "CONSUMABLE",
		AppleProductId: a.productId,
	},nil
}

func NewAppleIAPMock(mock appleiap.AppStoreMock) AppleIAPMock {
	return AppleIAPMock{mock:mock}
}