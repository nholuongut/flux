package daemon

import (
	"testing"
)

func TestEndpointInference(t *testing.T) {
	wsEndpoint := "ws://cloud.https://github.com/nholuongut/api/flux"
	wssEndpoint := "wss://cloud.https://github.com/nholuongut/api/flux"
	httpEndpoint := "http://cloud.https://github.com/nholuongut/api/flux"
	httpsEndpoint := "https://cloud.https://github.com/nholuongut/api/flux"

	assertExpected(t, wsEndpoint, httpEndpoint, wsEndpoint)
	assertExpected(t, wssEndpoint, httpsEndpoint, wssEndpoint)
	assertExpected(t, httpEndpoint, httpEndpoint, wsEndpoint)
	assertExpected(t, httpsEndpoint, httpsEndpoint, wssEndpoint)
}

func assertExpected(t *testing.T, input, expectedHTTP, expectedWS string) {
	actualHTTP, actualWS, err := inferEndpoints(input)
	if err != nil {
		t.Error(err)
	}
	assertEquals(t, expectedHTTP, actualHTTP)
	assertEquals(t, expectedWS, actualWS)
}

func assertEquals(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected [%s], actual [%s]", expected, actual)
	}
}

func TestUnsupportedEndpoint(t *testing.T) {
	_, _, err := inferEndpoints("mailto://cloud.https://github.com/nholuongut/api/flux")
	if err == nil {
		t.Error("Expected err, got nil")
	}
}
