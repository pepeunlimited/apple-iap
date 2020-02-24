package twrip

import (
	"context"
	"github.com/pepeunlimited/apple-iap/pkg/appstore"
	"github.com/pepeunlimited/apple-iap/pkg/rpc/appleiap"
	"github.com/twitchtv/twirp"
	"log"
	"testing"
)

func TestAppleIAPServer_VerifyReceipt(t *testing.T) {
	ctx    := context.TODO()
	server := NewAppleIAPServer(appstore.NewAppStore())
	receipt, err := server.VerifyReceipt(ctx, &appleiap.VerifyReceiptParams{
		Receipt: "receipt",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Print(receipt.Type)
	log.Print(receipt.Status)
}

func TestAppleIAPServer_VerifyReceiptMock(t *testing.T) {
	ctx    := context.TODO()
	mock := appstore.NewAppStoreMock([]int{0, 2000})
	server := NewAppleIAPServer(&mock)
	_,err := server.VerifyReceipt(ctx, &appleiap.VerifyReceiptParams{
		Receipt: "receipt",
	})
	if err == nil {
		t.FailNow()
	}
	if err.(twirp.Error).Msg() != "apple_iap_internal" {
		t.FailNow()
	}
	receipt,err := server.VerifyReceipt(ctx, &appleiap.VerifyReceiptParams{
		Receipt: "receipt",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if receipt == nil {
		t.FailNow()
	}
	receipt,err = server.VerifyReceipt(ctx, &appleiap.VerifyReceiptParams{
		Receipt: "receipt",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if receipt == nil {
		t.FailNow()
	}
}