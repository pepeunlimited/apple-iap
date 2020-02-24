package appstore

import (
	"context"
	"github.com/awa/go-iap/appstore"
)


type client struct {
	client 		*appstore.Client
}

func (apple client) VerifyReceipt(ctx context.Context, receipt string, password *string) (*appstore.IAPResponse, error) {
	req := appstore.IAPRequest{
		ReceiptData: receipt,
	}
	if password != nil {
		req.Password = *password
	}
	resp := &appstore.IAPResponse{}
	err := apple.client.Verify(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return read(resp)
}

func NewAppStore() AppStore {
	return client{client:appstore.New()}
}