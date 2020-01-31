package server

import (
	"context"
	"github.com/pepeunlimited/rpc-starter-kit/applerpc"
	"github.com/pepeunlimited/rpc-starter-kit/internal/app/app1/apple"
	"github.com/twitchtv/twirp"
	"log"
	"testing"
)

func TestAppleIAPServer_VerifyReceipt(t *testing.T) {
	ctx    := context.TODO()
	server := NewAppleIAPServer(apple.NewAppStore())
	receipt, err := server.VerifyReceipt(ctx, &applerpc.VerifyReceiptParams{
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
	mock := apple.NewAppStoreMock([]int{0, 2000})
	server := NewAppleIAPServer(&mock)
	_,err := server.VerifyReceipt(ctx, &applerpc.VerifyReceiptParams{
		Receipt: "receipt",
	})
	if err == nil {
		t.FailNow()
	}
	if err.(twirp.Error).Msg() != "internal" {
		t.FailNow()
	}
	receipt,err := server.VerifyReceipt(ctx, &applerpc.VerifyReceiptParams{
		Receipt: "receipt",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if receipt == nil {
		t.FailNow()
	}
	receipt,err = server.VerifyReceipt(ctx, &applerpc.VerifyReceiptParams{
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