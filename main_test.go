package main

import (
	"fmt"
	"testing"
)

func TestConnectToLXDserver(t *testing.T) {

	err := connectToLXDserver()

	if err != nil {
		t.Log("OK")
		fmt.Printf("no error")
	} else {
		t.Log("not OK")
		fmt.Printf("retured error")
	}

}
