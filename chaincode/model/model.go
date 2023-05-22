package model

// Account 账户，虚拟管理员和若干客户账号
type Account struct {
	AccountId string  `json:"accountId"` //账号ID
	UserName  string  `json:"userName"`  //账号名
	Balance   float64 `json:"balance"`   //余额
}

type RealEstate struct {
	RealEstateID string `json:"realEstateId"` //农产品ID
	AccountID    string `json:"accountId"`    //供应商ID
	Proprietor   string `json:"proprietor"`   //所有者(客户)(客户AccountId)
	Yzcxx        string `json:"yzcxx"`        //烟草原料产地
	Yzcwz        string `json:"yzcwz"`        //烟草原料品质（优中良）
	Yzcdh        string `json:"yzcdh"`        //烟草原料生长周期
	Yzcnzebh     string `json:"yzcnzebh"`     //烟草原料采集时间
	Yzctz        string `json:"yzctz"`        //烤烟等级
	Yzcrcsj      string `json:"yzcrcsj"`      //烤烟温度
	Yzcclsj      string `json:"yzcclsj"`      //烘干温湿度
	Yzcymjzcs    string `json:"yzcymjzcs"`    //烹饪状态
	PfsAccountID string `json:"pfsaccountid"` //工业公司ID
	Tzcmc        string `json:"tzcmc"`        //仓库等级和编号
	Tzcdh        string `json:"tzcdh"`        //烟叶年份
	Tzcnzebh     string `json:"tzcnzebh"`     //采购时间
	Tzccpbh      string `json:"tzccpbh"`      //醇化后烟叶等级
	Tzcnrbw      string `json:"tzcnrbw"`      //复烤温湿度
	Tzczl        string `json:"tzczl"`        //醇化详细信息
	Tzctzsj      string `json:"tzctzsj"`      //醇化时间
	Tzcczyxm     string `json:"tzcczyxm"`     //烟叶回潮信息
	Zhisigongyi  string `json:"zhisigongyi"`  //制丝工艺
	Jiagongsj    string `json:"jiagongsj"`    //加工时间
	Chengpinpp   string `json:"chengpinpp"`   //成品品牌
	Chengpingg   string `json:"chengpingg"`   //成品规格
	Ccysccx      string `json:"cysccx"`       //出场运输车车型
	Ccysccph     string `json:"ccysccph"`     //出场运输车车牌号
	Ccyssj       string `json:"ccyssj"`       //运输时间
	Sygsaccid    string `json:"sygsaccid" `   //商业公司ID
	Yuqywlgsmc   string `json:"yuqywlgsmc"`   //成品采购时间
	Yuqycph      string `json:"yuqycph"`      //成品采购规格
	Yuqyysxx     string `json:"yuqyysxx"`     //成品配送车型号
	Syysccp      string `json:"syysccp"`      //成品配送车车牌号
	Syyysj       string `json:"syyysj"`       //成品配送时间
	Xsqymc       string `json:"xsqymc"`       //零售店
	// Yzqytx       string `json:"yzqytx"`
	// Yzqyts       string `json:"yzqyts"`
	// Tzqytx       string `json:"tzqytx"`
	// Tzqyts       string `json:"tzqyts"`
	// Ysqytx       string `json:"ysqytx"`
	// Ysqyts       string `json:"ysqyts"`
	// Xsqytx       string `json:"xsqytx"`
	// Xsqyts       string `json:"xsqyts"`
}

const (
	AccountKey    = "account-key"
	RealEstateKey = "real-estate-key"
)
