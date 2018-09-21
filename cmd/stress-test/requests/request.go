package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendTx(authTx string) ([]byte, error) {
	res, err := http.Get(NodeUrl + fmt.Sprintf(BroadcastTx, authTx))

	if err != nil {
		return []byte(""), err
	}

	content, err1 := processResponse(res)
	return content, err1

}

func QueryTx(tx string) ([]byte, error) {
	res, err := http.Get(NodeUrl + fmt.Sprintf(QueryUri, tx))

	if err != nil {
		return []byte(""), err
	}

	return processResponse(res)
}

func QueryNonceTx(tx string) (int64, error) {
	res, err := QueryTx(tx)

	if err != nil {
		return 0, err
	}

	return decodeNonceResponse(res), nil
}

func QueryBalanceTx(tx string) (map[string]int64, error) {
	ret := map[string]int64{}

	res, err := QueryTx(tx)

	if err != nil {
		return ret, err
	}

	for _, coin := range decodeBalanceResponse(res) {
		ret[coin.Denom] = coin.Amount
	}

	return ret, nil
}

func processResponse(res *http.Response) ([]byte, error) {
	if res.StatusCode != 200 {
		return []byte(""), fmt.Errorf("Non-successful response")
	} else {
		content, err := ioutil.ReadAll(res.Body)

		defer res.Body.Close()

		if err != nil {
			return []byte(""), err
		}

		return content, nil
	}

}
