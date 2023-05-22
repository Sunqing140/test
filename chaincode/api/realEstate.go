package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// CreateRealEstate 新建农产品(养殖户)
func CreateRealEstate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 34 {
		return shim.Error("参数个数不满足")
	}

	accountId := args[0] //accountId用于验证是否为养殖户
	proprietor := args[1]
	yzcxx := args[2]
	yzcwz := args[3]
	yzcdh := args[4]
	yzcnzebh := args[5]
	yzctz := args[6]
	yzcrcsj := args[7]
	yzcclsj := args[8]
	yzcymjzcs := args[9]
	pfsaccountid :=args[10]
	tzcmc := args[11]
	tzcdh := args[12]
	tzccpbh := args[13]
	tzcnrbw := args[14]
	tzczl := args[15]
	tzctzsj := args[16]
	tzcczyxm := args[17]
	zhisigongyi := args[18]
	jiagongsj := args[19]
	chengpinpp := args[20]
	chengpingg := args[21]
	cysccx := args[22]
	ccysccph := args[23]
	ccyssj := args[24]
	sygsaccid := args[25]
	yuqywlgsmc := args[26]
	yuqycph := args[27]
	yuqyysxx := args[28]
	syysccp := args[29]
	syyysj := args[30]
	xsqymc := args[31]
	realEstateId := args[32]
	tzcnzebh := args[33]

	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{accountId})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("操作人权限验证失败%s", err))
	}
	var account model.Account
	if err = json.Unmarshal(resultsAccount[0], &account); err != nil {
		return shim.Error(fmt.Sprintf("查询操作人信息-反序列化出错: %s", err))
	}
	// if account.UserName != "养殖户模拟账户" {
	// 	return shim.Error(fmt.Sprintf("操作人权限不足%s", err))
	// }
	//判断客户是否存在
	resultsProprietor, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{proprietor})
	if err != nil || len(resultsProprietor) != 1 {
		return shim.Error(fmt.Sprintf("客户proprietor信息验证失败%s", err))
	}
	realEstate := &model.RealEstate{
		// RealEstateID: stub.GetTxID()[:16],
		RealEstateID: realEstateId,
		AccountID:    accountId,
		Proprietor:   proprietor,
		Yzcxx:        yzcxx,
		Yzcwz:        yzcwz,
		Yzcdh:        yzcdh,
		Yzcnzebh:     yzcnzebh,
		Yzctz:        yzctz,
		Yzcrcsj:      yzcrcsj,
		Yzcclsj:      yzcclsj,
		Yzcymjzcs:    yzcymjzcs,
		PfsAccountID: pfsaccountid,
		Tzcmc:        tzcmc,
		Tzcdh:        tzcdh,
		Tzccpbh:      tzccpbh,
		Tzcnrbw:      tzcnrbw,
		Tzczl:        tzczl,
		Tzctzsj:      tzctzsj,
		Tzcczyxm:     tzcczyxm,
		Zhisigongyi:  zhisigongyi,
		Jiagongsj:    jiagongsj,
		Chengpinpp:   chengpinpp,   
		Chengpingg:   chengpingg,   
		Ccysccx:      cysccx,       
		Ccysccph:     ccysccph,    
		Ccyssj:       ccyssj,     
		Sygsaccid:    sygsaccid,
		Yuqywlgsmc:   yuqywlgsmc,
		Yuqycph:      yuqycph,
		Yuqyysxx:     yuqyysxx,
		Syysccp:      syysccp,
		Syyysj:       syyysj,
		Xsqymc:       xsqymc,
		Tzcnzebh:     tzcnzebh,
		// Yzqytx:       yzqytx,
		// Yzqyts:       time.Unix(int64(yzqyts.GetSeconds()), int64(yzqyts.GetNanos())).Local().Format("2006-01-02 15:04:05"),
		// Tzqytx:       tzqytx,
		// Tzqyts:       tzqyts,
		// Ysqytx:       ysqytx,
		// Ysqyts:       ysqyts,
		// Xsqytx:       xsqytx,
		// Xsqyts:       xsqyts,
	}

	// 写入账本
	if err := utils.WriteLedger(realEstate, stub, model.RealEstateKey, []string{realEstate.Proprietor, realEstate.RealEstateID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	realEstateByte, err := json.Marshal(realEstate)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(realEstateByte)
}

// QueryRealEstateList 查询农产品(可查询所有，也可根据所有人查询名下农产品)
func QueryRealEstateList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var realEstateList []model.RealEstate
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.RealEstateKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var realEstate model.RealEstate
			err := json.Unmarshal(v, &realEstate)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryRealEstateList-反序列化出错: %s", err))
			}
			realEstateList = append(realEstateList, realEstate)
		}
	}
	realEstateListByte, err := json.Marshal(realEstateList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryRealEstateList-序列化出错: %s", err))
	}
	return shim.Success(realEstateListByte)
}
