# AztecUp

AztecUp is a program that reports on the status, or health, of the overall Aztec network platform service endpoints.

The current proof of concept is written in the [Go programming language](https://go.dev/) on Linux (Ubuntu).

A standalone Linux executable of `aztecup` is available here:

[https://github.com/pietergreyling/aztecup/blob/main/src/golang/aztecup](https://github.com/pietergreyling/aztecup/blob/main/src/golang/aztecup)


### How to install Go on Linux (Ubuntu)

```shell
$ sudo apt update
$ sudo apt upgrade
$ sudo apt search golang-go
$ sudo apt install golang-go 
```

### How to install Go on macOS with Homebrew

[https://formulae.brew.sh/formula/go](https://formulae.brew.sh/formula/go)

Install command: 

`brew install go`

### Recommended: 
#### Install [the Delve Go debugger](https://github.com/go-delve/delve) 

```shell
$ go install github.com/go-delve/delve/cmd/dlv@latest
```

#### Check your Go environment 

```shell
$ go env
```

### The Go/Golang documentation

[Build simple, secure, scalable systems with Go](https://go.dev/)

### Clone this project
```shell
$ git clone https://github.com/pietergreyling/aztecup.git
$ cd aztecup
$ cd src
$ cd golang
// - - OR - -
$ cd aztecup/src/golang
```

### Run and build the aztecup program

```shell
$ go run aztecup.go
$ go build aztecup.go 
$ ./aztecup
-- Hello AztecUp!
-- User: [yourusername]
-- Host: [yourcomputername]
```

Expected current behaviour:

```shell
./aztecup                                                      

- - - - - - - - - -
-- Hello AztecUp!
-- User: [yourusername]
-- Host: [yourcomputername]
- - - - - - - - - -

=> 0 https://api.aztec.network/aztec-connect-dev/falafel

 &{200 OK 200 HTTP/2.0 2 0 map[Content-Length:[40] Content-Type:[application/json; charset=utf-8] Date:[Tue, 20 Dec 2022 21:13:31 GMT] Vary:[Accept-Encoding, Origin]] {0xc00018a300} 40 [] false false map[] 0xc000114000 0xc00012a160}

- - - - - - - - - -

=> 1 https://api.aztec.network/aztec-connect-dev/falafel/status

 &{200 OK 200 HTTP/2.0 2 0 map[Content-Type:[application/json; charset=utf-8] Date:[Tue, 20 Dec 2022 21:13:32 GMT] Vary:[Accept-Encoding, Origin]] 0xc000438120 -1 [] false true map[] 0xc0001ac100 0xc00012a160}

- - - - - - - - - -

=> 2 https://api.aztec.network/aztec-connect-dev/falafel/rollup/0

 &{200 OK 200 HTTP/2.0 2 0 map[Content-Type:[application/json; charset=utf-8] Date:[Tue, 20 Dec 2022 21:13:32 GMT] Vary:[Accept-Encoding, Origin]] 0xc0004383c0 -1 [] false true map[] 0xc000114200 0xc00012a160}

- - - - - - - - - -

=> 3 https://api.aztec.network/aztec-connect-testnet/falafel

 &{503 Service Unavailable 503 HTTP/2.0 2 0 map[Content-Length:[162] Content-Type:[text/html] Date:[Tue, 20 Dec 2022 21:13:32 GMT] Server:[awselb/2.0]] {0xc00018a600} 162 [] false false map[] 0xc0001ac200 0xc00012a160}

- - - - - - - - - -

=> 4 https://api.aztec.network/aztec-connect-testnet/falafel/status

 &{503 Service Unavailable 503 HTTP/2.0 2 0 map[Content-Length:[162] Content-Type:[text/html] Date:[Tue, 20 Dec 2022 21:13:32 GMT] Server:[awselb/2.0]] {0xc00018a780} 162 [] false false map[] 0xc0001ac300 0xc00012a160}

- - - - - - - - - -

=> 5 https://api.aztec.network/aztec-connect-testnet/falafel/rollup/0

 &{503 Service Unavailable 503 HTTP/2.0 2 0 map[Content-Length:[162] Content-Type:[text/html] Date:[Tue, 20 Dec 2022 21:13:32 GMT] Server:[awselb/2.0]] {0xc00018a900} 162 [] false false map[] 0xc0001ac400 0xc00012a160}
```

## Development Next Steps:

- Continue to build with Go.
- Destructure JSON responses (application/json) and report specific properties into a map from the Aztec apis.
- Design output/api response formats. 
- Drill down into Aztec APIs.
- Build a localhost REST service/api.
- Host a public service/api.
- Build a web front-end.

### Drilling down: An example is lifting information from an API like this

[https://api.aztec.network/aztec-connect-prod/falafel/status](https://api.aztec.network/aztec-connect-prod/falafel/status)

```json
{
    "version":"2.1.4",
    "numTxsPerRollup":896,
    "numUnsettledTxs":595,
    "numTxsInNextRollup":295,
    "pendingTxCount":300,
    "pendingSecondClassTxCount":0,
    "totalTxs":656812,
    "totalBlocks":5867,
    "nextPublishNumber":116093,
    "proverless":false,
    "rollupSize":1024,
    "blockchainStatus": 
... }
```

## Aztec Platform Network Services Endpoints

This documentation section serves as a collection of Aztec platform network endpoints.

A selection of these will be accessed by AztecUp to determine, monitor, and report on the status or health of the overall Aztec network platform.  

### Service Naming Scheme

Across all of the Aztec connect services, one can replace 

**`aztec-connect-dev`** 

with

**`aztec-connect-testnet`** 

and 

**`aztec-connect-prod`** 

to access the different networks.

### Aztec Endpoints: Sites and Dashboards

https://explorer.aztec.network/
https://aztec-connect-testnet-explorer.aztec.network/
https://metrics.aztec.network/?orgId=1
https://metrics.aztec.network/d/1HehplOVk/aztec-connect-kpis?search=open&orgId=1&refresh=10s
https://dune.com/gm365/Aztec
https://dune.com/gm365/aztec-v2
https://aztec-connect-testnet.zk.money/signin
https://dune.com/emesrever/aztec-connect-bridge-leaderboard
https://aztec-connect-testnet-eth-host.aztec.network:8545/
https://aztec-connect-dev-eth-host.aztec.network:8545/
https://aztec-connect-dev-eth-host.aztec.network:8545/k8auhfxwp931irx704r6
https://aztec-connect-testnet-faucet.aztec.network/
https://aztec-connect-testnet-eth-host.aztec.network:8545/6rw8n2loc5bfb5x1l17w
https://aztec-connect-dev.zk.money/

### Aztec Endpoints: APIs

https://api.aztec.network/aztec-connect-dev/falafel
https://api.aztec.network/aztec-connect-dev/falafel/status
https://api.aztec.network/aztec-connect-dev/falafel/rollup/0
https://api.aztec.network/aztec-connect-testnet/falafel
https://api.aztec.network/aztec-connect-testnet/falafel/status
https://api.aztec.network/aztec-connect-testnet/falafel/rollup/0
https://api.aztec.network/aztec-connect-testnet/falafel/rollups?take=10&skip=100
https://api.aztec.network/aztec-connect-prod/falafel
https://api.aztec.network/aztec-connect-prod/falafel/status
https://api.aztec.network/aztec-connect-prod/falafel/rollup/0
https://api.aztec.network/aztec-connect-prod/falafel/rollups?take=10&skip=100
https://api.aztec.network/aztec-connect-prod/falafel/rollups?take=50&skip=100

#### References
https://docs.aztec.network/developers/sequencer-api

###  Aztec Endpoints: Testing

https://aztec-connect-testnet-explorer.aztec.network/
https://api.aztec.network/aztec-connect-testnet/falafel/status
https://aztec-connect-testnet-eth-host.aztec.network:8545/
https://api.aztec.network/aztec-connect-testnet/falafel
https://api.aztec.network/aztec-connect-prod/falafel
https://aztec-connect-dev-eth-host.aztec.network:8545/
https://aztec-connect-dev-eth-host.aztec.network:8545/k8auhfxwp931irx704r6
https://aztec-connect-testnet-eth-host.aztec.network:8545/6rw8n2loc5bfb5x1l17w

### Aztec Endpoints: Smart Contracts

https://etherscan.io/address/0xff1f2b4adb9df6fc8eafecdcbf96a2b351680455#code
https://etherscan.io/address/0xff1f2b4adb9df6fc8eafecdcbf96a2b351680455
https://etherscan.io/address/0xff1f2b4adb9df6fc8eafecdcbf96a2b351680455#analytics

### Aztec Endpoints: CircleCI

https://status.circleci.com/

#### Pipelines
https://app.circleci.com/launchpad/github/AztecProtocol/getting-started
https://app.circleci.com/pipelines/github/AztecProtocol/aztec2-internal
https://app.circleci.com/pipelines/github/AztecProtocol/aztec2-internal?branch=v2.1-testnet
https://app.circleci.com/pipelines/github/AztecProtocol/aztec-connect-bridges
https://app.circleci.com/pipelines/github/AztecProtocol/barretenberg

