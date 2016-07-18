package main

import (
	"./util"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// read access_token
	buf, err := util.ReadFile("access_token.txt")
	if err != nil {
		fmt.Println("ReadFile error,", err)
		return
	}

	// send url
	url := "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=" + string(buf)
	body, err := util.GetHttp(string(url))
	if err != nil {
		fmt.Println("http url request failed", err)
		return
	}
	// indent json
	var out bytes.Buffer
	json.Indent(&out, body, "", "\t")
	out.WriteTo(os.Stdout)
	return
}
