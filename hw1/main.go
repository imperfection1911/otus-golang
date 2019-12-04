package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

// HelloNow prints current ntp time.
func HelloNow(host string) bool {
	time, err := ntp.Time(host)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return false
	}
	fmt.Println(time)
	return true
}

func main() {
	if HelloNow("0.beevik-ntp.pool.ntp.org") {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
