package bsdex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

AUCTION_ENDPOINT = "/auction"
SYMBOL_SEPARATOR = "-"

type AuctionResp struct {
	BuyPrice   string `json:"buy_price"`
	BuyVolume  string `json:"buy_volume"`
	Market     string `json:"error"`
	SellPrice  string `json:"sell_price"`
	SellVolume string `json:"sell_volume"`
}

func Auction(baseAsset string, quoteAsset string) (*AuctionResp, error) {
  b, err := a.requestGET(baseAsset + SYMBOL_SEPARATOR + quoteAsset + AUCTION_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}
  
	err = json.Unmarshal(b, &resp)
	return resp, err
}

