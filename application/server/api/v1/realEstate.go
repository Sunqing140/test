package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RealEstateRequestBody struct {
	AccountId    string `json:"accountId"` //操作人ID
	Proprietor   string `json:"proprietor"`
	Yzcxx        string `json:"yzcxx"`
	Yzcwz        string `json:"yzcwz"`
	Yzcdh        string `json:"yzcdh"`
	Yzcnzebh     string `json:"yzcnzebh"`
	Yzctz        string `json:"yzctz"`
	Yzcrcsj      string `json:"yzcrcsj"`
	Yzcclsj      string `json:"yzcclsj"`
	Yzcymjzcs    string `json:"yzcymjzcs"`
	PfsAccountID string `json:"pfsaccountid"`
	Tzcmc        string `json:"tzcmc"`
	Tzcdh        string `json:"tzcdh"`
	Tzccpbh      string `json:"tzccpbh"`
	Tzcnrbw      string `json:"tzcnrbw"`
	Tzczl        string `json:"tzczl"`
	Tzctzsj      string `json:"tzctzsj"`
	Tzcczyxm     string `json:"tzcczyxm"`
	Zhisigongyi  string `json:"zhisigongyi"`
	Jiagongsj    string `json:"jiagongsj"`
	Chengpinpp   string `json:"chengpinpp"`   
	Chengpingg   string `json:"chengpingg"`  
	Ccysccx      string `json:"cysccx"`     
	Ccysccph     string `json:"ccysccph"`    
	Ccyssj       string `json:"ccyssj"`      
	Sygsaccid    string `json:"sygsaccid" `
	Yuqywlgsmc   string `json:"yuqywlgsmc"`
	Yuqycph      string `json:"yuqycph"`
	Yuqyysxx     string `json:"yuqyysxx"`
	Syysccp      string `json:"syysccp"`
	Syyysj       string `json:"syyysj"`
	Xsqymc       string `json:"xsqymc"`
	RealEstateId string `json:"realEstateId"`
	Tzcnzebh     string `json:"tzcnzebh"`
}

type RealEstateQueryRequestBody struct {
	Proprietor string `json:"proprietor"` //所有者(客户)(客户AccountId)
}

func CreateRealEstate(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealEstateRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	bodyBytes = append(bodyBytes, []byte(body.AccountId))
	bodyBytes = append(bodyBytes, []byte(body.Proprietor))
	bodyBytes = append(bodyBytes, []byte(body.Yzcxx))
	bodyBytes = append(bodyBytes, []byte(body.Yzcwz))
	bodyBytes = append(bodyBytes, []byte(body.Yzcdh))
	bodyBytes = append(bodyBytes, []byte(body.Yzcnzebh))
	bodyBytes = append(bodyBytes, []byte(body.Yzctz))
	bodyBytes = append(bodyBytes, []byte(body.Yzcrcsj))
	bodyBytes = append(bodyBytes, []byte(body.Yzcclsj))
	bodyBytes = append(bodyBytes, []byte(body.Yzcymjzcs))
	bodyBytes = append(bodyBytes, []byte(body.PfsAccountID))
	bodyBytes = append(bodyBytes, []byte(body.Tzcmc))
	bodyBytes = append(bodyBytes, []byte(body.Tzcdh))
	bodyBytes = append(bodyBytes, []byte(body.Tzccpbh))
	bodyBytes = append(bodyBytes, []byte(body.Tzcnrbw))
	bodyBytes = append(bodyBytes, []byte(body.Tzczl))
	bodyBytes = append(bodyBytes, []byte(body.Tzctzsj))
	bodyBytes = append(bodyBytes, []byte(body.Tzcczyxm))
	bodyBytes = append(bodyBytes, []byte(body.Zhisigongyi))
	bodyBytes = append(bodyBytes, []byte(body.Jiagongsj))
	bodyBytes = append(bodyBytes, []byte(body.Chengpinpp))
	bodyBytes = append(bodyBytes, []byte(body.Chengpingg))
	bodyBytes = append(bodyBytes, []byte(body.Ccysccx))
	bodyBytes = append(bodyBytes, []byte(body.Ccysccph))
	bodyBytes = append(bodyBytes, []byte(body.Ccyssj))
	bodyBytes = append(bodyBytes, []byte(body.Sygsaccid))
	bodyBytes = append(bodyBytes, []byte(body.Yuqywlgsmc))
	bodyBytes = append(bodyBytes, []byte(body.Yuqycph))
	bodyBytes = append(bodyBytes, []byte(body.Yuqyysxx))
	bodyBytes = append(bodyBytes, []byte(body.Syysccp))
	bodyBytes = append(bodyBytes, []byte(body.Syyysj))
	bodyBytes = append(bodyBytes, []byte(body.Xsqymc))
	bodyBytes = append(bodyBytes, []byte(body.RealEstateId))
	bodyBytes = append(bodyBytes, []byte(body.Tzcnzebh))

	//调用智能合约
	resp, err := bc.ChannelExecute("createRealEstate", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryRealEstateList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealEstateQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Proprietor != "" {
		bodyBytes = append(bodyBytes, []byte(body.Proprietor))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryRealEstateList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}
