package validator

import (
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/rpc-starter-kit/applerpc"
	"github.com/twitchtv/twirp"
)

type AppleIAPServerValidator struct {}

func (v AppleIAPServerValidator) VerifyReceipt(params *applerpc.VerifyReceiptParams) error {
	if validator.IsEmpty(params.Receipt) {
		return twirp.RequiredArgumentError("receipt")
	}
	return nil
}

func NewAppleIAPServerValidator() AppleIAPServerValidator {
	return AppleIAPServerValidator{}
}