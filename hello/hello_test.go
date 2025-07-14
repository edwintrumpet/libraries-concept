package hello

import "testing"

func TestSayHello(t *testing.T) {
	expectedResponse := "Hello world!"

	res := SayHello()
	if res != expectedResponse {
		t.Errorf("response is not as expected\nreceived: %s\nexpected: %s", res, expectedResponse)
	}
}
