package lunoService

import (
	"github.com/bhbosman/goTrader/internal/lunoApi/client"
	"net/http"
)

type lunoTransport struct {
	keyId     string
	keySecret string
	next      http.RoundTripper
}

func newLunoTransport(keyId string, keySecret string, next http.RoundTripper) (http.RoundTripper, error) {
	return &lunoTransport{keyId: keyId, keySecret: keySecret, next: next}, nil
}

func (self *lunoTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	request.SetBasicAuth(self.keyId, self.keySecret)
	return self.next.RoundTrip(request)
}

type lunoHttpRequestDoer struct {
	roundTripper http.RoundTripper
}

func newLunoHttpRequestDoer(roundTripper http.RoundTripper) (client.HttpRequestDoer, error) {
	return &lunoHttpRequestDoer{roundTripper: roundTripper}, nil
}

func (self *lunoHttpRequestDoer) Do(req *http.Request) (*http.Response, error) {
	return self.roundTripper.RoundTrip(req)
}
