package httpclient

import (
	"testing"
)

func TestCall(t *testing.T) {

	result, err := Call(t.Context(), &defaultUrls[0])
	if err != nil {
		t.Error(err)
	}

	if !result.Success {
		t.Errorf("Test failed. result.Success: %v, expected: %v", result, true)
	}
}
