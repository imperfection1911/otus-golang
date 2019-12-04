package main

import "testing"

func TestHelloNow(t *testing.T) {
	WrongHost := "1qaz2wsx.org"
	if HelloNow(WrongHost) {
		t.Fatalf("bad return value for ntp server %s", WrongHost)
	}

}