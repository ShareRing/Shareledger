package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sharering/shareledger/pkg/swap/abi/sharetoken"
	"strings"
)

func decodeTxParams(abi abi.ABI, v map[string]interface{}, data []byte) (map[string]interface{}, error) {
	m, err := abi.MethodById(data[:4])
	if err != nil {
		return map[string]interface{}{}, err
	}
	if err := m.Inputs.UnpackIntoMap(v, data[4:]); err != nil {
		return map[string]interface{}{}, err
	}
	return v, nil
}

func main() {
	conn, err := ethclient.Dial("https://ropsten.infura.io/v3/bf1a5b4c59cb45ea8ebf48497d3295ae")
	if err != nil {
		panic(err)
	}

	tx, _, err := conn.TransactionByHash(context.Background(), common.HexToHash("0x1ed9b34ec35245cba3014aa76939a82206df4f85c632d77613a23fffb46b57a7"))
	erc20Abi, err := abi.JSON(strings.NewReader(string(sharetoken.SharetokenMetaData.ABI)))
	v := make(map[string]interface{})
	a, e := decodeTxParams(erc20Abi, v, tx.Data())
	fmt.Println("aaa", a, e)
}

//
//func main() {
//	conn, err := ethclient.Dial("https://ropsten.infura.io/v3/bf1a5b4c59cb45ea8ebf48497d3295ae")
//	if err != nil {
//		panic(err)
//	}
//	ctx := context.Background()
//	tx, pending, er := conn.TransactionByHash(ctx, common.HexToHash("0x1ed9b34ec35245cba3014aa76939a82206df4f85c632d77613a23fffb46b57a7"))
//	fmt.Println(tx, pending, er)
//	fmt.Println("TX", tx.Value().String(), tx.To().Hex(), string(tx.Data()))
//
//	erc20Abi, err := abi.JSON(strings.NewReader(string(sharetoken.SharetokenMetaData.ABI)))
//	if err != nil {
//		panic(err)
//	}
//	a, e := decodeTxParams2(erc20Abi, "transfer", tx.Data())
//	c := a[0].(common.Address)
//	d := a[1].(*big.Int)
//	fmt.Println("c", c.Hex(), "d", d.String())
//
//	fmt.Println("a", a, "e", e)
//
//}
//
//func decodeTxParams(abi abi.ABI, method string, data []byte) ([]interface{}, error) {
//	txParams, err := abi.Methods[method].Inputs.UnpackValues(data)
//	if err != nil {
//		return []interface{}{}, err
//	}
//	return txParams, nil
//}
//func decodeTxParams2(abi abi.ABI, v map[string]interface{}, data []byte) (map[string]interface{}, error) {
//	m, err := abi.MethodById(data[:4])
//	if err != nil {
//		return map[string]interface{}{}, err
//	}
//	if err := m.Inputs.UnpackIntoMap(v, data[4:]); err != nil {
//		return map[string]interface{}{}, err
//	}
//	return v, nil
//}
