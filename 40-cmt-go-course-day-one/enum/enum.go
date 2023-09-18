package main

import (
	"fmt"
)

type Wochentag int

const (
	Mo Wochentag = iota
	Di
	Mi
	Do
	Fr
	Sa
	So
)

func ist_wt(wt Wochentag) {
	switch os := wt; os {
	case So:
		fmt.Println("Sonntag.")
	default:
		fmt.Println("Werktag.")
	}
}

func main() {
	ist_wt(So)
}
