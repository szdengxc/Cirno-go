package ciweimao

import (
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
	"github.com/zsakvo/Cirno-go/structure"
	"github.com/zsakvo/Cirno-go/util"
)

func GetDetail(bid string) structure.BookInfo {
	var err error
	var res []byte
	paras := req.Param{
		"book_id": bid,
	}
	res, err = util.Get("/book/get_info_by_id", paras)
	util.PanicErr(err)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.DetailStruct
	err = json.Unmarshal(res, &result)
	util.PanicErr(err)
	return result.Data.BookInfo
}
