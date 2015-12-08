package logs

import (
	"testing"
)

func TestPrintl1(t *testing.T) {
	err := Printl("test log", "")
	if err != nil {
		t.Errorf("Error occured: %#v", err)
	}
}

func TestPrintl2(t *testing.T) {
	err := Printl("test log file", "./test.log")
	if err != nil {
		t.Errorf("Error occured: %#v", err)
	}

}
