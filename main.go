package main

import (
	//"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	cryptmsg, err := getPrice("btc")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("crypt %s buy %s ,sell %s, high: %s, low: %s, vol: %s\n", "eth", cryptmsg.Data.Buy, cryptmsg.Data.Sell, cryptmsg.Data.High, cryptmsg.Data.Low, cryptmsg.Data.Vol)

	secretFile, err := os.Open("secret2.json")

	checkError("Fail to read secret %s", err)
	defer secretFile.Close()

	byteValue, _ := ioutil.ReadAll(secretFile)

	var secretKeeper SecretKeeper
	json.Unmarshal(byteValue, &secretKeeper)

	assets, err := checkAssets(secretKeeper)
	if err != nil {
		log.Fatalln(err)
	}

	for _, asset := range assets {
		fmt.Println(asset)
	}

	trades, err := getTradeHistory(secretKeeper, "btc")

	if err != nil {
		log.Fatalf(err.Error())
	}

	file, err := os.OpenFile("btc.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}

	sort.Sort(Trades(trades))
	count := 0.0
	total := 0.0
	for _, trade := range trades {
		price, _ := strconv.Atoi(trade.Price)
		amount, _ := strconv.ParseFloat(trade.Amount, 64)
		if trade.Side == "buy" {
			count += amount
			total += float64(price) * amount
		} else {
			count -= amount
			total -= float64(price) * amount
		}
		total = normalizeFloat(total)
		count = normalizeFloat(count)
	}
	price, _ := strconv.Atoi(trades[len(trades)-1].Price)
	value := float64(price) * count
	fmt.Printf("COST: %v ", total)
	fmt.Printf("VALUE: %v\n", value)
	fmt.Printf("Rate of Return %v%%", normalizeFloat((value-total)/total)*100)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func normalizeFloat(num float64) float64 {
	return math.Round(num*10000) / 10000
}
