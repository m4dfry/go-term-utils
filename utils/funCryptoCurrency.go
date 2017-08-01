package utils

import (
  "fmt"
)

func SimplifyCurrMap(mp map[string](map[string]float64), reqC string) map[string]float64 {
	ret := map[string]float64{}
  for ck, cv := range mp {
		ret[ck] = cv[reqC]
  }
	return ret
}

func CryptoCurrency(args []string) {
	reqCurrency := ""
	if len(args) > 1 {
		reqCurrency = args[1]
	}

	var ccurr map[string](map[string]float64)
// https://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,LTC,DASH,XRP,XEM,ETH,ETC&tsyms=EUR
  APICall("https://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,LTC,DASH,XRP,XEM,ETH,ETC&tsyms=" + reqCurrency, &ccurr)

  simplified_ccurr := SimplifyCurrMap(ccurr, reqCurrency)
  sorted_ccurr := revSortByValue(simplified_ccurr)

  for _, v := range sorted_ccurr {
        fmt.Printf("1 %-5s -> %10.4f %s\n", v.Key, v.Value, reqCurrency)
  }
}
