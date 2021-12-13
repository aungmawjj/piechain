package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/aungmawjj/piechain/cclib"
	"github.com/aungmawjj/piechain/examples/auction"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	zkNodes     = "localhost:2181"
	ethEndpoint = "localhost:8545"
	platform    = "ethereum"
	signerID    = "1"
	keyfile     = "../../keys/key1"
	keypassword = "password"

	eventService *cclib.EventService
	ethClient    *ethclient.Client
	signer       *cclib.Signer
)

func main() {
	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.StringVar(&ethEndpoint, "eth", ethEndpoint, "eth endpoint")
	flag.StringVar(&signerID, "id", signerID, "signer id")
	flag.StringVar(&keyfile, "key", keyfile, "private key file")
	flag.Parse()

	var err error

	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s", ethEndpoint))
	check(err)

	signer, err = cclib.NewSigner(keyfile, keypassword)
	check(err)

	eventService, err = cclib.NewEventService(
		strings.Split(zkNodes, ","),
		fmt.Sprintf("signer/%s/%s", platform, signerID),
	)
	check(err)

	eventService.Register(auction.AuctionEndingEvent, handleAuctionEnding)

	err = eventService.Start()
	check(err)

	select {}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
