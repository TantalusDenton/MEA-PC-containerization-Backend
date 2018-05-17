package main

import (
	"testing"
)

func TestConnectToLXDserver(t *testing.T) {

	err := connectToLXDserver()

	if err != nil {
		t.Log("OK")
	} else {
		t.Log("not OK")
	}

}
