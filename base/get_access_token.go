package main

import (
	"./util"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

type Config struct {
	App *App
}

type App struct {
	AppID     string
	AppSecret string
}

type AccessToken struct {
	Access_token string
	Expires_in   int
}

func config() (c *Config, err error) {
	var (
		file *os.File
		blob []byte
	)

	c = new(Config)
	if file, err = os.Open("config.toml"); err != nil {
		return nil, err
	}
	defer file.Close()
	if blob, err = ioutil.ReadAll(file); err != nil {
		return nil, err
	}
	if err = toml.Unmarshal(blob, c); err != nil {
		return nil, err
	}
	return c, nil
}

func main() {
	// read config
	var c *Config
	var err error
	if c, err = config(); err != nil {
		fmt.Println("read Config err,", err)
		return
	}

	// get content
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + c.App.AppID + "&secret=" + c.App.AppSecret
	body, err := util.GetHttp(url)
	if err != nil {
		fmt.Println("http url request failed", err)
		return
	}
	var at AccessToken
	if err = json.Unmarshal(body, &at); err != nil {
		fmt.Println(err)
		return
	}

	// write to file
	fout, err := os.Create("access_token.txt")
	defer fout.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fout.WriteString(at.Access_token)

	return
}
