// Unit Test code for gomixpanel.
//
// Use `go test` on the commandline to test this package
//
//
package gomixpanel

import (
	"testing"
)

func TestMixpanelWithToken(t *testing.T) {
	SetApiToken("y")
	ok, err := SendEvent("ok", map[string]interface{}{
		"name": "miao",
		"age":  12,
	})
	if err != nil {
		t.Errorf("should get no error")
	}
	if !ok {
		t.Errorf("should be ok")
	}

}

func TestMixpanelWithEmptyToken(t *testing.T) {
	SetApiToken("")
	ok, err := SendEvent("ok", map[string]interface{}{
		"name": "miao",
		"age":  12,
	})
	if err != nil {
		t.Errorf("should get no error")
	}
	if ok {
		t.Errorf("should not be ok")
	}
}
