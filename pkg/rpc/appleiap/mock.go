package appleiap

import (
	"context"
	"github.com/pepeunlimited/apple-iap/pkg/appstore"
	"log"
)

type AppleIAPMock struct {
	mock      appstore.AppStoreMock
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

func NewAppleIAPMock(mock appstore.AppStoreMock) AppleIAPMock {
	return AppleIAPMock{mock:mock}
}