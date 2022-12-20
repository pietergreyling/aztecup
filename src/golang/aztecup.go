// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// $ go run aztecup.go
// $ go build aztecup.go 
// $ ./aztecup
// -- Hello AztecUp!
// -- user:  pieter
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

package main
 
import ( 
	"fmt" 
	"os" 
	"net/http"
	"log"
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

	fmt.Println("\n- - - - - - - - - -")

}

func say_hello() {
	fmt.Println("- - - - - - - - - -")
    fmt.Println("-- Hello AztecUp!")
    fmt.Println("-- User:", os.Getenv("USER"))
    fmt.Println("- - - - - - - - - -")
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

		fmt.Println("\n- - - - - - - - - -")
		fmt.Println("\n=>", i, url)
	    // fmt.Println("\n", err, resp)
	    fmt.Println("\n", resp)
	}
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
