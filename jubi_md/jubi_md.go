// jubi_md project jubi_md.go
package jubi_md

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	_ "io"
	"jubi_data"
	"log"
	"net/url"
	_ "reflect"
	"strconv"
	_ "strings"
	"time"

	"github.com/levigross/grequests"
)

func signature(s, privateKey string) string {
	h := md5.New()
	h.Write([]byte(privateKey))
	ss := fmt.Sprintf("%x", h.Sum(nil))
	mac := hmac.New(sha256.New, []byte(ss))
	mac.Write([]byte(s))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func Blance(coin string) (b, l, t float32) {
	v := url.Values{}
	v.Add("key", "v4ab2-d6ky6-dzwz6-i35m1-27fic-njwu7-qqa11")
	v.Add("nonce", millisecond())
	v.Add("version", strconv.Itoa(2))
	v.Add("signature", signature(v.Encode(), "uvf(r-trqjq-CZvLv-yQK.d-&EV,s-^H3gH-5J^kF"))

	m := make(map[string]string)
	for i, _ := range v {
		m[i] = v.Get(i)

	}
	resp, err := grequests.Post("https://www.jubi.com/api/v1/balance", &grequests.RequestOptions{Data: m, RequestTimeout: time.Second})
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	var blance jubi_data.Blance
	if err := json.Unmarshal(resp.Bytes(), &blance); err == nil {
		b, l, t = blance.Get(coin)
		fmt.Println(coin, "账户数量=", b, "锁定数量=", "可用总额=", t)
		//fmt.Println("Cny_balance=", blance.Cny_balance, "Cny_lock=", blance.Cny_lock, "Rss_balance=", blance.Rss_balance, "Rss_lock=", blance.Rss_lock, "Ifc_balance=", blance.Ifc_balance, "Ifc_lock=", blance.Ifc_lock, "Wdc_balance=", blance.Wdc_balance, "Wdc_lock=", blance.Wdc_lock, "Btc_balance=", blance.Btc_balance, "Bcc_lock=", blance.Btc_lock)
	} else {
		log.Printf("Blance error decoding sakura response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("sakura response: %q", resp.Bytes())
	}
	return
}

/*
市场挂牌查询：
high - 最高价
low - 最低价
buy - 买一价
sell - 卖一价
last - 最近一次成交价
vol - 成交量
volume - 成交额
*/
func Dicker(coin string) (co jubi_data.Coin) {
	resp, err := grequests.Get("https://www.jubi.com/api/v1/ticker", &grequests.RequestOptions{Params: map[string]string{"coin": coin}, RequestTimeout: time.Second * 5})
	fmt.Println("DickerStatusCode=", resp.StatusCode)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	if err := json.Unmarshal(resp.Bytes(), &co); err == nil {
		return co
	} else {
		log.Printf("Dicker error decoding sakura response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("sakura response: %q", resp.Bytes())
	}
	return
}

//https://www.jubi.com/api/v1/Trade_add?key=p25ir-c99rq-62fsj-hjhkf-h952a-zbsu6-hi8n2&nonce=1501778185&price=1.0&type=buy&coin=rio&amount=1000.0&signature=aca087b461d205937457b77b5d0c4c7aca86c4c7468044ccebdea8259e2874b0
func Add_reqiest(price string, count float32, coin, typ string) {
	t_add := url.Values{}
	t_add.Add("key", "p25ir-c99rq-62fsj-hjhkf-h952a-zbsu6-hi8n2")
	t_add.Add("nonce", millisecond())
	t_add.Add("price", price)
	t_add.Add("amount", strconv.Itoa(int(count)))
	t_add.Add("coin", coin)
	t_add.Add("type", typ)
	t_add.Add("signature", signature(t_add.Encode(), ")1qhX-2x]Be-t)Y/I-et;sx-7qX@d-^tSa1-X713q"))
	mc := make(map[string]string, 10)
	for i, _ := range t_add {
		mc[i] = t_add.Get(i)
	}
	resp, err := grequests.Post("https://www.jubi.com/api/v1/Trade_add", &grequests.RequestOptions{Data: mc, RequestTimeout: time.Second * 5})
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	} else {
		faileInfo(resp.Bytes())
	}
}

/*
查询挂单Id
*/
func Trade_list(coin, typ string) (mbuy, msell map[string][]string) {
	tr_list := url.Values{}
	tr_list.Add("nonce", millisecond())
	tr_list.Add("key", "d3dcb-hcke3-1d9ts-7d9k7-6d3nk-f8327-hfawt")
	tr_list.Add("version", "2")
	tr_list.Add("coin", coin)
	tr_list.Add("type", typ)
	tr_list.Add("signature", signature(tr_list.Encode(), "Nwbe[-igbE2-^Sydk-NsOU$-gafpf-y]OTS-zn4KC"))
	trm := make(map[string]string)
	for i, _ := range tr_list {
		trm[i] = tr_list.Get(i)
	}
	resp, err := grequests.Post("https://www.jubi.com/api/v1/Trade_list", &grequests.RequestOptions{Data: trm, RequestTimeout: time.Second * 5})
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	var Vlist []jubi_data.Tr_list
	n := len(resp.String())
	if n > 30 {
		fmt.Println(resp.String())
		if err = json.Unmarshal(resp.Bytes(), &Vlist); err == nil {
			//解析数据
			mbuy = make(map[string][]string, 100)
			msell = make(map[string][]string, 100)
			for _, v := range Vlist {
				if v.Type == "buy" {
					mbuy[coin] = append(mbuy[coin], v.Id)
				}
				if v.Type == "sell" {
					msell[coin] = append(msell[coin], v.Id)
				}

			}
			fmt.Println("查询成功")
		} else {
			log.Printf("Trade_list error decoding sakura response: %v", err)
			if e, ok := err.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("sakura response: %q", resp.Bytes())
		}
	} else if n > 10 {
		fmt.Println("没有挂单信息")
		faileInfo(resp.Bytes())
	}
	return
}

/*
查询订单状态
Trade_view
*/
func OrderStatusQuery(param map[string][]string) (Pending map[string][]string) {
	Om := make(map[string]string, 10)
	Oder := url.Values{}
	Oder.Add("key", "b9at4-kudxg-sas8x-bbg6q-6zpgs-di4d5-iqdti")
	Oder.Add("version", "2")
	Oder.Add("nonce", "000")
	Oder.Add("id", "00")
	Oder.Add("coin", "00")
	Pending = make(map[string][]string, len(param))
	for coin, value := range param {
		Oder.Set("coin", coin)
		for i, _ := range value {
			Oder.Set("id", value[i])
			Oder.Set("nonce", millisecond())
			Oder.Add("signature", signature(Oder.Encode(), "Cyu/r-/ONN@-H^c15-}c4KP-,(d&q-pB1Sm-GbS{$"))
			for j, _ := range Oder {
				Om[j] = Oder.Get(j)
			}
			//time.Sleep(time.Second)
			//fmt.Println(Om)
			resp, err := grequests.Post("https://www.jubi.com/api/v1/Trade_view", &grequests.RequestOptions{Data: Om, RequestTimeout: time.Second * 5})
			Oder.Del("signature")
			if err != nil {
				log.Printf("OrderStatusQuery error decoding sakura response: %v", err)
				if e, ok := err.(*json.SyntaxError); ok {
					log.Printf("syntax error at byte offset %d", e.Offset)
				}
				log.Printf("sakura response: %q", resp.Bytes())
			}
			if len(resp.String()) > 30 {
				var Ostatus jubi_data.Tr_status
				if err := json.Unmarshal(resp.Bytes(), &Ostatus); err != nil {
					log.Fatal(err)
				}
				if Ostatus.Status == "open" {
					Pending[coin] = append(Pending[coin], strconv.Itoa(Ostatus.Id))
				}

			} else if len(resp.String()) > 10 {

				faileInfo(resp.Bytes())
			}

		}
	}
	return
}

/*
取消订单
*/
func Trade_cancel(coin, id string) {
	Vcen := url.Values{}
	Vcen.Add("nonce", millisecond())
	Vcen.Add("key", "urqah-59uqp-tpaye-r58dt-ks2dy-ran1c-tu3i5")
	Vcen.Add("version", "2")
	Vcen.Add("coin", coin)
	Vcen.Add("id", id)
	Vcen.Add("signature", signature(Vcen.Encode(), "KTyyx-8i8xe-GbqZG-9uncY-i%3}^-vTYBh-%1%qz"))
	Cm := make(map[string]string, 10)
	for i, _ := range Vcen {
		Cm[i] = Vcen.Get(i)
	}
	resp, err := grequests.Post("https://www.jubi.com/api/v1/Trade_cancel", &grequests.RequestOptions{Data: Cm, RequestTimeout: time.Second * 5})
	if err != nil {
		log.Fatal(err)
	}
	faileInfo(resp.Bytes())
}

/*
市产交易
*/
func Orders(coin string) (Dadanbuy, Dadansell, price float32) {
	resp, err := grequests.Get("https://www.jubi.com/api/v1/orders/", &grequests.RequestOptions{Params: map[string]string{"coin": coin}, RequestTimeout: time.Second * 5})
	fmt.Println("OrdersStatusCode=", resp.StatusCode)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Close()
	var value []jubi_data.Orders
	if len(resp.Bytes()) > 30 {
		if err := json.Unmarshal(resp.Bytes(), &value); err != nil {
			log.Printf("Orders error decoding sakura response: %v", err)
			if e, ok := err.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("sakura response: %q", resp.Bytes())
		} else {
			timemui := time.Now().Unix()
			timemui = timemui - 900
			Dadanbuy, Dadansell = 0, 0
			for i := range value {
				price = value[i].Price
				if value[i].Date > strconv.FormatInt(timemui, 10) {
					if value[i].Type == "buy" {
						Dadanbuy += value[i].Amount
					} else {
						Dadansell += value[i].Amount
					}
				}
			}
		}
	} else {
		faileInfo(resp.Bytes())
	}
	return
}

/*
错误提示函数
*/
func faileInfo(param []byte) {
	var errfail jubi_data.Verror
	if len(param) < 10 {
		return
	}
	if err := json.Unmarshal(param, &errfail); err != nil {
		log.Printf("faileInfo error decoding sakura response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("sakura response: %q", param)
	}
	if errfail.Result {
		fmt.Println("成功。", "Id=", errfail.Id)
	} else {
		errInfo(errfail.Code)
	}
}

/*
错误信息提示
*/
func errInfo(numb string) {
	var code int
	var err error
	if code, err = strconv.Atoi(numb); err != nil {
		log.Fatal(err)
	}
	switch code {
	default:
		fmt.Println("未知错误")
	case 100:
		fmt.Println("必选参数不能为空")
	case 101:
		fmt.Println("非法参数")
	case 102:
		fmt.Println("请求的虚拟币不存在")
	case 103:
		fmt.Println("密钥不存在")
	case 104:
		fmt.Println("签名不匹配")
	case 105:
		fmt.Println("权限不足")
	case 106:
		fmt.Println("请求过期(nonce错误)")
	case 200:
		fmt.Println("余额不足")
	case 201:
		fmt.Println("买卖的数量小于最小买卖额度")
	case 202:
		fmt.Println("下单价格必须在 0 - 1000000 之间")
	case 204:
		fmt.Println("成交金额不能少于 10 元")
	case 203:
		fmt.Println("订单不存在")
	case 205:
		fmt.Println("gooc限制挂单价格")
	case 401:
		fmt.Println("系统错误")
	case 402:
		fmt.Println("请求过于频繁")
	case 403:
		fmt.Println("非开放API")
	case 404:
		fmt.Println("P限制不能请求该资源")
	case 405:
		fmt.Println("币种交易暂时关闭")
	}

}

func millisecond() string {
	tempt := strconv.FormatInt(time.Now().UnixNano()/100000, 10)
	return tempt
}
