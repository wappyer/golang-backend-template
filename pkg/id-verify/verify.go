package id_verify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

/*
温馨提示：
1.解析结果时，先判断code，当code!=0时，表示调用已失败，无需再继续
code描述
0:成功
400:参数错误
20010:身份证号为空或非法
20310:姓名为空或非法
404:请求资源不存在
500:系统内部错误，请联系服务商
501:第三方服务异常
604:接口停用
1001:其他，以实际返回为准

2.再判断下面result中的res（1 一致；2 不一致；3 无记录）
result.res 出现'无记录'时，有以下几种原因：
(1)现役军人、武警官兵、特殊部门人员及特殊级别官员；
(2)退役不到2年的军人和士兵（根据军衔、兵种不同，时间会有所不同，一般为2年）；
(3)户口迁出，且没有在新的迁入地迁入；
(4)户口迁入新迁入地，当地公安系统未将迁移信息上报到公安部（上报时间地域不同而有所差异）；
(5)更改姓名，当地公安系统未将更改信息上报到公安部（上报时间因地域不同而有所差异）；
(6)移民；
(7)未更换二代身份证；
(8)死亡。
(9)身份证号确实不存在
*/

const (
	host    = "https://eid.shumaidata.com"
	path    = "/eid/check"
	appCode = "733ade8f51504b9a875d673268c22f49"
)

type VerifyReq struct {
	Idcard string `json:"idcard"`
	Name   string `json:"name"`
}

type VerifyResp struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Result  VerifyResult `json:"result"`
}

type VerifyResult struct {
	Name        string `json:"name"`        //姓名
	Idcard      string `json:"idcard"`      //身份证号
	Res         string `json:"res"`         //核验结果状态码，1 一致；2 不一致；3 无记录
	Description string `json:"description"` //核验结果状态描述
	Sex         string `json:"sex"`
	Birthday    string `json:"birthday"`
	Address     string `json:"address"`
}

func IdVerify(param *VerifyReq) (VerifyResp, error) {
	ret := VerifyResp{}

	values := url.Values{}
	values.Add("idcard", param.Idcard)
	values.Add("name", param.Name)
	urlS, _ := url.Parse(fmt.Sprintf("%s%s", host, path))
	urlS.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", urlS.String(), nil)
	if err != nil {
		return ret, err
	}
	req.Header.Set("Authorization", "APPCODE "+appCode)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ret, err
	}
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(bodyByte, &ret)
	if err != nil {
		fmt.Printf("[IdVerify]身份验证返回错误，err: %v ;body:%v \n", err, string(bodyByte))
		return ret, err
	}

	return ret, nil
}
