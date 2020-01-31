package apple

import (
	"context"
	"github.com/awa/go-iap/appstore"
	"github.com/twitchtv/twirp"
	"log"
)

const (
	IapMode = "IAP_MODE" // LIVE, MOCK
)

var (
	// 21000
	// the App Store could not read the JSON object you provided.
	ErrReadIssue 				= twirp.NewError(twirp.Aborted, "read_issue")
	// 21002
	// the data in the receipt-data property was malformed or missing.
	ErrMalformed 				= twirp.NewError(twirp.Malformed, "malformed")
	// 21003
	// the receipt could not be authenticated.
	ErrAuthentication 			= twirp.NewError(twirp.Unauthenticated, "authentication")
	// 21004
	// the shared secret you provided does not match the shared secret on file for your account.
	ErrMismatchedSecret 		= twirp.NewError(twirp.Aborted, "mismatched_secret")
	// 21005
	// the receipt server is not currently available.
	ErrServerDown 				= twirp.NewError(twirp.Unavailable, "server_down")
	// 21006 (Only returned for iOS 6 style transaction receipts for auto-renewable subscriptions.)
	// This receipt is valid but the subscription has expired. When this status code is returned to your server, the receipt data is also decoded and returned as part of the response.
	ErrSubscriptionExpired 		= twirp.NewError(twirp.Aborted, "subscription_expired")
	// 21007
	// this receipt is from the test environment, but it was sent to the production environment for verification. Send it to the test environment instead
	ErrFromTestToProduction 	= twirp.NewError(twirp.Aborted, "from_test_to_prod")
	// 21008
	// This receipt is from the production environment, but it was sent to the test environment for verification. Send it to the production environment instead.
	ErrFromProductionToTest 	= twirp.NewError(twirp.Aborted, "from_prod_to_test")
	// 21010
	// This receipt could not be authorized. Treat this the same as if a purchase was never made.
	ErrAuthorization 			= twirp.NewError(twirp.Unauthenticated, "authorization")

	// 21100-21199
	// Internal data access error
	// should do retry
	ErrInternal   				=  twirp.NewError(twirp.Internal, "internal")
)

type AppStore interface {
	VerifyReceipt(ctx context.Context, receipt string, password *string) (*appstore.IAPResponse, error)
}

func read(response *appstore.IAPResponse) (*appstore.IAPResponse, error) {
	switch response.Status {
	case 0:
		return response, nil
	case 21000:
		return nil, ErrReadIssue
	case 21002:
		return nil, ErrMalformed
	case 21003:
		return nil, ErrAuthentication
	case 21004:
		return nil, ErrMismatchedSecret
	case 21005:
		return nil, ErrServerDown
	case 21006:
		return response, ErrSubscriptionExpired
	case 21007:
		return nil, ErrFromTestToProduction
	case 21008:
		return nil, ErrFromProductionToTest
	case 21010:
		return nil, ErrAuthorization
	default:
		return nil, ErrInternal
	}
}

func NewAppStoreByMode(mode string) AppStore {
	log.Print("using AppleIAP: "+mode)
	if mode == "LIVE" {
		return NewAppStore()
	}
	mock := NewAppStoreMock([]int{0})
	return &mock
}