package ciweimao

import (
	"fmt"

	"../config"
	"../structure"
	"../util"
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
)

func Login() {
	util.InitTmpReq()
	var err error
	var res []byte
	var name string
	var passwd string
	fmt.Printf("account: ")
	fmt.Scanln(&name)
	fmt.Printf("password: ")
	fmt.Scanln(&passwd)
	paras := req.Param{
		"login_name": name,
		"passwd":     passwd,
	}
	res, err = util.Get("/signup/login", paras)
	util.PanicErr(err)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.LoginStruct
	err = json.Unmarshal(res, &result)
	if err != nil {
		panic(err)
	} else {
		config.Write(name, passwd, result.Data.LoginToken, result.Data.ReaderInfo.Account)
	}
}
