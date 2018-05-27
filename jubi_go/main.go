// jubi_go project main.go
package main

import (
	_ "encoding/json"
	"fmt"
	"jubi_md"
	"log"
	_ "log"
	"strconv"
	"time"
)

func GoFunc(Coin string, sNumb, Nmb, Cp float32, ch chan<- int) {
	//var buycount, sellcount float32
	//slep := make(chan int, 2)
	//var TP float32 = 0
	var maxb float32 = 0
	var maxs float32 = 0
	var err error
	var sellprice float64 = 0
	for {
		buycount, sellcount, TP := jubi_md.Orders(Coin)
		fmt.Println("bct=", buycount, "\t", "sct=", sellcount, Coin, "=", TP)
		info := jubi_md.Dicker(Coin)
		if maxb < buycount {
			maxb = buycount

		}
		if maxs < sellcount {
			maxs = sellcount

		}
		fmt.Println("买", maxb)
		fmt.Println("卖", maxs)
		if buycount > Nmb {
			fmt.Println("大单买入")
			amount, _, _ := jubi_md.Blance(Coin)
			jubi_md.Add_reqiest(info.Sell, amount, Coin, "sell")
			Nmb += 50000
		}
		if sellcount > sNumb {
			fmt.Println("大单卖出")
			_, _, total := jubi_md.Blance(Coin)
			if f, err := strconv.ParseFloat(info.Sell, 32); err == nil {
				amount := total / 3 / float32(f)
				jubi_md.Add_reqiest(info.Sell, amount, Coin, "buy")
				sNumb += 50000
			}

		}

		if sellprice, err = strconv.ParseFloat(info.Sell, 32); err != nil {
			log.Fatal(err)
		}
		Qp := float32(sellprice)
		if Qp < Cp {
			//清仓
			//取消所有买单和买单数量
			mb, ms := jubi_md.Trade_list(Coin, "open")
			omb := jubi_md.OrderStatusQuery(mb)
			for cb, value := range omb {
				for i, _ := range value {
					jubi_md.Trade_cancel(cb, value[i])
				}
			}
			Oms := jubi_md.OrderStatusQuery(ms)
			for cs, value := range Oms {
				for i, _ := range value {
					jubi_md.Trade_cancel(cs, value[i])
				}
			}
			b_amount, _, _ := jubi_md.Blance(Coin)
			jubi_md.Add_reqiest(info.Sell, b_amount, Coin, "sell")
		}
		time.Sleep(time.Second * 3)
	}
}

func main() {
	ch := make(chan int, 1)
	//go GoFunc("ans", 100000, 10000, 185, ch)
	//go GoFunc("wdc", 1000000, 1000000, 0.5, ch)
	//go GoFunc("rss", 1000000, 1000000, 3.5, ch)
	go GoFunc("btc", 1000, 1000, 24000, ch)
	//go GoFunc("xsgs", 100000, 1000000, 60, ch)
	//	<-ch
	//<-ch
	<-ch
	//<-ch
	//<-ch
	//jubi_md.Orders("btc")
	//	for {
	//		var dicker chan jubi_data.Coin
	//		dicker = make(chan jubi_data.Coin, 1)
	//
	//		jubi_md.Dicker("btc")
	//		blance()
	//
	//		if coin_rio, ok := <-dicker; ok {
	//			_ = coin_rio
	//			//jubi_md.Add_reqiest(coin_rio.High, 1000, "btc", "buy", slep)
	//			mb, ms := jubi_md.Trade_list("rio", "all", slep)
	//			omb := jubi_md.OrderStatusQuery(mb, slep)
	//			for coin, value := range omb {
	//				for i, _ := range value {
	//					fmt.Println(coin, "=", value[i])
	//				}
	//			}
	//			Oms := jubi_md.OrderStatusQuery(ms, slep)
	//			for coin, value := range Oms {
	//				for i, _ := range value {
	//					fmt.Println(coin, "=", value[i])
	//				}
	//			}

	//		} else {
	//			//fmt.Println("失败")
	//		}

	//		<-slep
	//		<-slep
	//		<-slep
	//		ncount++
	//		fmt.Println("循环了第", ncount)
	//	}

}
