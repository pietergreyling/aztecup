// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Install Golang on Ubuntu:
// $ sudo apt update
// $ sudo apt upgrade
// $ sudo apt search golang-go
// $ sudo apt search gccgo-go
// $ sudo apt install golang-go
// $ sudo apt install gccgo-go
//
// Clone this project:
// $ git clone https://github.com/pietergreyling/aztecup.git
// $ cd aztecup
// $ cd src
// $ cd golang
// - - OR - -
// $ cd aztecup/src/golang
//
// Run and build the aztecup program:
// $ go run aztecup.go
// $ go build aztecup.go
// $ ./aztecup
// -- Hello AztecUp!
// -- user:  pieter
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var Aztec_Apis = []string{
	// devnet (dev)
	"https://api.aztec.network/aztec-connect-dev/falafel",
	"https://api.aztec.network/aztec-connect-dev/falafel/status",
	"https://api.aztec.network/aztec-connect-dev/falafel/rollup/0",
	// testnet (test)
	"https://api.aztec.network/aztec-connect-testnet/falafel",
	"https://api.aztec.network/aztec-connect-testnet/falafel/status",
	"https://api.aztec.network/aztec-connect-testnet/falafel/rollup/0",
	// mainnet (prod)
	"https://api.aztec.network/aztec-connect-prod/falafel",
	"https://api.aztec.network/aztec-connect-prod/falafel/status",
	"https://api.aztec.network/aztec-connect-prod/falafel/rollup/0",
}

func main() {

	say_hello()

	// list_aztec_apis()

	ping_aztec_apis()

}

func say_hello() {
	fmt.Println("\n- - - - - - - - - -")
	fmt.Println("-- Hello AztecUp!")
	fmt.Println("-- User:", os.Getenv("USER"))
	fmt.Println("-- Host:", get_host())
	fmt.Println("- - - - - - - - - -")
}

func get_host() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return hostname
}

func list_aztec_apis() {
	for i, url := range Aztec_Apis {
		fmt.Println("=>", i, url)
	}
}

func ping_aztec_apis() {
	for i, url := range Aztec_Apis {

		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\n=>", i, url)
		// fmt.Println("\n", err, resp)
		fmt.Println("\n", resp)

		fmt.Println("\n- - - - - - - - - -")
	}
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
