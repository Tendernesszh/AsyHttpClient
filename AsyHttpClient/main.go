package main

import (
	"fmt"
	"github.com/Tendernesszh/AsyHttpClient/work"
	"github.com/Tendernesszh/AsyHttpClient/agency"
)

func main() {
	fmt.Println("The time of Using synchronize is")
	synchronous()
	fmt.Println("The time of using asynchronism way")
	asynchronous()
}
