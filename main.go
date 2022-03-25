package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Trade struct {
	ID     int     `json:"id"`
	Market int     `json:"market"`
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
	IsBuy  bool    `json:"is_buy"`
}

type Markets struct {
	ID          int     `json:"market"`
	Trades      int     `json:"total_trades"`
	Buys        int     `json:"total_buys"`
	BuysPercent float64 `json:"percentage_buy"`
	Volume      float64 `json:"total_volume"`
	VolumeMean  float64 `json:"mean_volume"`
	Price       float64 `json:"total_price"`
	PriceMean   float64 `json:"mean_price"`
	VWAP        float64 `json:"volume_weighted_average_price"`
}

type M map[int]Markets

func main() {
	//Start a timer!
	start := time.Now()

	var trade Trade
	markets := make(M)

	//Initialize readers
	reader := bufio.NewReader(os.Stdin)

	//We know that stdin will start with a string -- BEGIN
	passReader := parseStrings(*reader, "BEGIN\n")

	//Decode JSON
	decoder := json.NewDecoder(&passReader)
	for {
		err := decoder.Decode(&trade)
		if err != nil {
			//TODO -- Once we run out of JSON the decoder barfs when reading END
			//Find a way to stop reading json, consume remaining statistical messages and exit
			break
		}
		//Compute trade
		computeTrade(trade, markets)
	}

	//Print results to stdout
	printMarketData(markets)

	//Report execution time
	duration := time.Since(start)
	fmt.Println(duration)
}

func computeTrade(trade Trade, markets M) {
	isBuy := 0

	if trade.IsBuy {
		isBuy = 1
	}

	if market, ok := markets[trade.Market]; ok {
		market.Trades += 1
		market.Buys += isBuy
		market.BuysPercent = float64(market.Buys) / float64(market.Trades)
		market.Volume += trade.Volume
		market.VolumeMean = float64(market.Volume) / float64(market.Trades)
		market.Price += trade.Price
		market.PriceMean = float64(market.Price) / float64(market.Trades)
		market.VWAP = (float64(market.Price) * float64(market.Volume)) / float64(market.Trades)

		markets[trade.Market] = market

	} else {
		market.ID = trade.Market
		market.Trades = 1
		market.Buys = isBuy
		market.BuysPercent = float64(isBuy)
		market.Volume = trade.Volume
		market.VolumeMean = trade.Volume
		market.Price = trade.Price
		market.PriceMean = trade.Price
		market.VWAP = (float64(market.Price) * float64(market.Volume)) / float64(market.Trades)

		markets[trade.Market] = market
	}
}

func printMarketData(markets M) {
	for _, market := range markets {
		b, err := json.Marshal(market)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(b)
	}

}

//Rudimentary way to try and catch BEGIN and END
func parseStrings(reader bufio.Reader, match string) bufio.Reader {
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Reached end of file")
			break
		}
		if err != nil {
			fmt.Println("caught error")
			fmt.Println(err)
		}
		if line == match {
			fmt.Println("Matched line:")
			fmt.Println(line)
			break
		}
		fmt.Println(line)
	}
	return reader
}
