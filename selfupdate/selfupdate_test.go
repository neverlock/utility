package selfupdate

import (
	"testing"
)

func TestSelfUpdate(t *testing.T) {
	err := SelfUpdate("1.0", "http://www.conf.in.th/auto-update/", "go-selfupdate-example")
	if err != nil {
		t.Errorf("Error occured: %#v", err)
	}
}
