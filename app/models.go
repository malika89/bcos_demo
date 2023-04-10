package app

// Ticket 交易信息
type Ticket struct {
	PlayAddress   string `json:"player_id"`      // 玩家地址
	PoolAddress   string `json:"pool_address"`   // 奖金池地址
	OrderHash     string `json:"order_hash"`     //	hashId
	Amount        int    `json:"amount"`         //	数量
	Timestamp     uint64 `json:"timestamp"`      // 时间戳
	LotteryNumber string `json:"lottery_number"` // 号码
}

type (
	QueryDetail struct {
		PlayAddress string `json:"player_id"`    // 玩家地址
		PoolAddress string `json:"pool_address"` // 奖金池地址
		OrderHash   string `json:"order_hash"`   //	hashId
		Data        Data   `json:"data"`         //	交易明细
	}

	Data struct {
		Amount        int    `json:"amount"`         //数量
		Timestamp     uint64 `json:"timestamp"`      // 时间戳
		LotteryNumber string `json:"lottery_number"` //号码
	}
)

type ContractQueryData struct {
	From  string      `json:"from"`
	To    string      `json:"to"`
	Value string      `json:"value"`
	Data  interface{} `json:"data"`
}
