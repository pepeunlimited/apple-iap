# cURL

### Install
```$ brew install jq > curl ... | jq```

##### VerifyReceipt
```
$  curl -H "Content-Type: application/json" \
-X POST "api.dev.pepeunlimited.com/twirp/pepeunlimited.apple.AppleIAPService/VerifyReceipt" \
-d '{"receipt": "base64+encoded"}'
```