package random

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	r := Str(10)
	fmt.Println(r)
	if r == "" {
		t.Errorf("Error occured: %s", r)
	}
}

func TestInt(t *testing.T) {
	r := Int(1, 100)
	fmt.Println(r)
	if r > 100 {
		t.Errorf("Error occured: %d", r)
	}

}

func TestFint(t *testing.T) {
	r := Fint(5)
	fmt.Println(r)
	if r > 5 {
		t.Errorf("error occured: %d", r)
	}
}
