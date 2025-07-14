package bye

import "testing"

func TestSayBye(t *testing.T) {
	expectedResponse := "Goodbye!"

	response := SayBye()
	if response != expectedResponse {
		t.Errorf(
			"response is not as expected\nreceived: %s\nexpected: %s",
			response,
			expectedResponse,
		)
	}
}
