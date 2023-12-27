package gitinfo

import (
	"fmt"
	"testing"
)

func Test_ID1(t *testing.T) {
	fmt.Println("Testing no file")
	msg, err := ID("./", "main")
	if err != nil {
		fmt.Println("Error: ", msg)
	} else {
		fmt.Println("Success: ", msg)
		t.Errorf("Error occured: %v , %s\n", err, msg)
	}
}
func Test_ID2(t *testing.T) {
	fmt.Println("Testing have file")
	msg, err := ID("../", "master")
	if err != nil {
		t.Errorf("Error occured: %v\n", err)
		fmt.Println("Error: ", msg)
	} else {
		fmt.Println("Success: ", msg)
	}
}
