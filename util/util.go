package util

import (
	"bytes"
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego/orm"
)
//不进行转义
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

//通过openid的到用户id
func GetIdByOpenid(openid string) int{
	id := -1
	var maps []orm.Params
	o := orm.NewOrm()
	num,err := o.Raw("select id from `user` where openid=?",openid).Values(&maps)
	if err == nil && num > 0 {
		strId := maps[0]["id"].(string) // slene
		id,_ = strconv.Atoi(strId)
	}
	return id
}
