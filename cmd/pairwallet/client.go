package main

func main() {
	//client, err := ethclient.Dial("https://rinkeby.infura.io/v3/eda69b1ff98142d7843d1478c4095d1a") //https://mainnet.infura.io/v3/eda69b1ff98142d7843d1478c4095d1a")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("we have a connection")
	//account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	//balance, err := client.BalanceAt(context.Background(), account, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(balance) // 25893180161173005034

	//blockNumber := big.NewInt(5532993)
	//balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(balanceAt) // 25729324269165216042

	//fbalance := new(big.Float)
	//fbalance.SetString(balanceAt.String())
	//ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	//fmt.Println(ethValue) // 25.729324269165216041
	//
	//pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	//fmt.Println(pendingBalance) // 25729324269165216042
}
