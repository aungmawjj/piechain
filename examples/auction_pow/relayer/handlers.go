package main

import (
	"encoding/json"
	"time"

	"github.com/aungmawjj/piechain/examples/auction_pow"
)

func handleOnEndAuction(payload []byte) {
	var args auction_pow.EndAuctionArgs
	err := json.Unmarshal(payload, &args)
	check(err)

	_, err = assetClient.EndAuction(args)
	check(err)
}

func runAuctionListener() {
	for {
		a := listenNewAuction()
		go onNewAuction(a)
	}
}

func onNewAuction(a *auction_pow.Auction) {
	payload, _ := json.Marshal(a)
	err := ccsvc.Publish(auction_pow.TopicOnBindAuction, payload)
	check(err)
}

func listenNewAuction() *auction_pow.Auction {
	lastID, err := assetClient.GetLastAuctionID()
	check(err)

	for {
		time.Sleep(1 * time.Second)
		auctionID, err := assetClient.GetLastAuctionID()
		check(err)
		if auctionID > lastID {
			auction, err := assetClient.GetAuction(auctionID)
			check(err)
			return auction
		}
	}
}
