// jubi_data project jubi_data.go
package jubidata

/*
市场挂牌情况
*/
type Coin struct {
	High   string  `json:'high'`
	Low    string  `json:'low'`
	Buy    string  `json:'buy'`
	Sell   string  `json:'sell'`
	Last   string  `json:'last'`
	Vol    float32 `json:'Vol'`
	Volume float32 `json:'Volume'`
}
type uM struct {
	Ci map[string]string
}

/*
错误信息提示解析
*/
type Verror struct {
	Result bool   `json:'result'`
	Id     string `json:'id'`
	Code   string `json:'code'`
}

/*
订单数量查询结构体
*/
type TrList struct {
	ID                 string  `json:'ID'`
	Datetime           string  `json:'datetime'`
	Type               string  `json:'type'`
	Price              float32 `json:'price'`
	Amount_original    float32 `json:'amount_original'`
	Amount_outstanding float32 `json:'Amount_outstanding'`
}

/*
 订单状态结构体
*/
type Tr_status struct {
	Id                 int     `json:'id'`
	Datetime           string  `json:'datetime'`
	Type               string  `json:'type'`
	Price              float32 `json:'price'`
	Amount_original    float32 `json:'amount_original'`
	Amount_outstanding float32 `json:'Amount_outstanding'`
	Status             string  `json:'status'`
}

/*
市场交易情况
*/
type Orders struct {
	Date   string  `json:'date'`
	Price  float32 `json:'price'`
	Amount float32 `json:'amount'`
	Tid    string  `json:'tid'`
	Type   string  `json:'buy'`
}

/*
账户信息
*/
type Blance struct {
	Uid          int     `json:'uid'`
	Nameauth     int     `json:'nameauth'`
	Moflag       int     `json:'moflag'`
	Asset        float32 `json:'asset'`
	Btc_balance  float32 `json:'btc_balance'`
	Btc_lock     float32 `json:'btc_lock'`
	Drk_balance  float32 `json:'drk_balance'`
	Drk_lock     float32 `json:'drk_lock'`
	Blk_balance  float32 `json:'blk_balance'`
	Blk_lock     float32 `json:'blk_lock'`
	Vrc_balance  float32 `json:'vrc_balance'`
	Vrc_lock     float32 `json:'vrc_lock'`
	Tfc_balance  float32 `json:'tfc_balance'`
	Tfc_lock     float32 `json:'tfc_lock'`
	Jbc_balance  float32 `json:'jbc_balance'`
	Jbc_lock     float32 `json:'jbc_lock'`
	Ltc_balance  float32 `json:'ltc_balance'`
	Ltc_lock     float32 `json:'ltc_lock'`
	Doge_balance float32 `json:'doge_balance'`
	Doge_lock    float32 `json:'doge_lock'`
	Xpm_balance  float32 `json:'xpm_balance'`
	Xpm_lock     float32 `json:'xpm_lock'`
	Ppc_balance  float32 `json:'ppc_balance'`
	Ppc_lock     float32 `json:'ppc_lock'`
	Wdc_balance  float32 `json:'wdc_balance'`
	Wdc_lock     float32 `json:'wdc_lock'`
	Vtc_balance  float32 `json:'vtc_balance'`
	Vtc_lock     float32 `json:'vtc_lock'`
	Max_balance  float32 `json:'max_balance'`
	Max_lock     float32 `json:'max_lock'`
	Ifc_balance  float32 `ifc_balance`
	Ifc_lock     float32 `json:'ifc_lock'`
	Zcc_balance  float32 `json:'zcc_balance'`
	Zcc_lock     float32 `json:'zcc_lock'`
	Zet_balance  float32 `zet_balance`
	Zet_lock     float32 `json:'zet_lock'`
	Eac_balance  float32 `json:'eac_balance'`
	Eac_lock     float32 `json:'eac_lock'`
	Plc_lock     float32 `json:'plc_lock'`
	Lkc_lock     float32 `json:'lkc_lock'`
	Hlb_lock     float32 `josn:'hlb_lock'`
	Rio_lock     float32 `json:'rio_lock'`
	Mryc_lock    float32 `json:'mryc_lock'`
	Dnc_balance  float32 `json:'dnc_balance'`
	Nxt_lock     float32 `json:'nxt_lock'`
	Nhgh_balance float32 `json:'nhgh_balance'`
	Bts_balance  float32 `josn:'bts_balance'`
	Tic_lock     float32 `json:'tic_lock'`
	Cny_balance  float32 `json:'cny_balance'`
	Mtc_balance  float32 `json:'mtc_balance'`
	Met_balance  float32 `json:'met_balance'`
	Game_lock    float32 `json:'game_lock'`
	Ktc_lock     float32 `json:'ktc_lock'`
	Mcc_balance  float32 `json:'mcc_balance'`
	Dnc_lock     float32 `json:'dnc_lock'`
	Lsk_balance  float32 `json:'lsk_balance'`
	Nhgh_lock    float32 `json:'nhgh_lock'`
	Bts_lock     float32 `json:'bts_lock'`
	Ico_balance  float32 `json:'ico_balance'`
	Cny_lock     float32 `json:'cny_lock'`
	Mtc_lock     float32 `json:'mtc_lock'`
	Met_lock     float32 `json:'met_lock'`
	Rss_balance  float32 `json:'rss_balance'`
	Pgc_balance  float32 `json:'pgc_balance'`
	Mcc_lock     float32 `json:'mcc_lock'`
	Gooc_balance float32 `json:'gooc_balance'`
	Lsk_lock     float32 `json:'lsk_lock'`
	Xsgs_balance float32 `json:'xsgs_balance'`
	Ugt_lock     float32 `json:'ugt_lock'`
	Ico_lock     float32 `json:'ico_lock'`
	Qec_balance  float32 `json:'qec_balance'`
	Ytc_balance  float32 `json:'ytc_balance'`
	Rss_lock     float32 `json:'rss_lock'`
	Pgc_lock     float32 `json:'pgc_lock'`
	Eth_balance  float32 `json:'eth_balance'`
	Gooc_lock    float32 `json:'gooc_lock'`
	Xas_balance  float32 `json:'xas_balance'`
	Xsgs_lock    float32 `json:'xsgs_lock'`
	Eos_balance  float32 `json:'eos_balance'`
	Bcc_balance  float32 `json:'bcc_balance'`
	Qec_lock     float32 `json:'qec_lock'`
	ytc_lock     float32 `json:'ytc_lock'`
	Rio_balance  float32 `json:'rio_balance'`
	Mryc_balance float32 `json:'mryc_balance'`
	Eth_lock     float32 `json:'eth_lock'`
	Xrp_balance  float32 `json:'xrp_balance'`
	Xas_lock     float32 `json:'xas_lock'`
	Ans_balance  float32 `json:'ans_balance'`
	Eos_lock     float32 `json:'eos_lock'`
	Bcc_lock     float32 `json:'bcc_lock'`
	Fz_balance   float32 `json:'fz_balance'`
	Fz_lock      float32 `json:'fz_lock'`
	Skt_balance  float32 `json:'skt_balance'`
	Skt_lock     float32 `json:'skt_lock'`
	Plc_balance  float32 `json:'plc_balance'`
	Lkc_balance  float32 `json:'lkc_balance'`
	Hlb_balance  float32 `json:'hlb_balance'`
	Etc_balance  float32 `json:'etc_balance'`
	Etc_lock     float32 `json:'etc_lock'`
	Xrp_lock     float32 `json:'xrp_lock'`
	Nxt_balance  float32 `json:'nxt_balance'`
	Peb_balance  float32 `json:'peb_balance'`
	Peb_lock     float32 `json:'peb_lock'`
	Ans_lock     float32 `json:'ans_lock'`
	Tic_balance  float32 `json:'tic_balance'`
	Btk_balance  float32 `json:'tic_balance'`
	Btk_lock     float32 `json:'btk_lock'`
}

//获取具体的帐户数量
func (this Blance) Get(coin string) (blance, lock, total float32) {
	total = this.Cny_balance
	switch coin {
	case "wdc":
		blance = this.Wdc_balance
		lock = this.Wdc_lock
	case "rss":
		blance = this.Rss_balance
		lock = this.Rss_lock
	case "rio":
		blance = this.Rio_balance
		lock = this.Rio_lock
	case "ifc":
		blance = this.Ifc_balance
		lock = this.Ifc_lock
	case "nxt":
		blance = this.Nxt_balance
		lock = this.Nxt_lock
	case "ans":
		blance = this.Ans_balance
		lock = this.Ans_lock
	case "btc":
		blance = this.Btc_balance
		lock = this.Btc_lock
	case "xsgs":
		blance = this.Xas_balance
		lock = this.Xas_lock
	default:
		blance = 0
		lock = 0

	}
	return
}
