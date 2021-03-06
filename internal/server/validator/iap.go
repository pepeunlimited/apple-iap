package validator

import (
	"github.com/pepeunlimited/apple-iap/pkg/rpc/appleiap"
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/twitchtv/twirp"
)

type AppleIAPServerValidator struct {}

func (v AppleIAPServerValidator) VerifyReceipt(params *appleiap.VerifyReceiptParams) error {
	if validator.IsEmpty(params.Receipt) {
		return twirp.RequiredArgumentError("receipt")
	}
	return nil
}

func NewAppleIAPServerValidator() AppleIAPServerValidator {
	return AppleIAPServerValidator{}
}