package apple

import (
	"context"
	"github.com/awa/go-iap/appstore"
)

type AppStoreMock struct {
	StatusCodes []int
	Receipt 	string
	Password 	*string
}

func (a *AppStoreMock) VerifyReceipt(ctx context.Context, receipt string, password *string) (*appstore.IAPResponse, error) {
	i := len(a.StatusCodes)
	resp := &appstore.IAPResponse{}
	if i > 1 {
		i -= 1
		statusCode := a.StatusCodes[i]
		a.StatusCodes = append(a.StatusCodes[:i], a.StatusCodes[i+1:]...)
		resp.Status = statusCode
	} else if i == 1 {
		statusCode := a.StatusCodes[i-1]
		resp.Status = statusCode
	}
	a.Receipt = receipt
	a.Password = password
	return read(resp)
}

func NewAppStoreMock(statusCodes []int) AppStoreMock {
	return AppStoreMock{StatusCodes:statusCodes}
}