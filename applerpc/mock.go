package applerpc

import (
	"context"
	"github.com/pepeunlimited/apple-iap/apple"
)

type AppleIAPMock struct {
	mock 		apple.AppStoreMock
	productId 	string
}

func (a AppleIAPMock) VerifyReceipt(ctx context.Context, params *VerifyReceiptParams) (*VerifyReceiptResponse, error) {
	receipt, err := a.mock.VerifyReceipt(ctx, params.Receipt, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyReceiptResponse{
		Status:         VerifyReceiptResponse_Status(receipt.Status),
		Type:           VerifyReceiptResponse_CONSUMABLE,
		AppleProductId: a.productId,
	},nil
}

func NewAppleIAPMock(mock apple.AppStoreMock) AppleIAPMock {
	return AppleIAPMock{mock:mock}
}