package main

import (
	"io/ioutil"
	"os"
	"net/http"
	"github.com/BurntSushi/toml"
	"encoding/json"
	"fmt"
)

type Config struct {
	App *App	
}

type App struct {
	AppID string
	AppSecret string	
}

type AccessToken struct {
	Access_token string
	Expires_in int
}

func config() (c *Config, err error){
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

func GetHttp(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body) 
	if err != nil {
		return nil, err
	}
	return body, nil

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
	body, err := GetHttp(url)
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

	// read from file
	/*
	fin, err := os.Open("access_token.txt")
	defer fin.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := make([]byte, 1024)
	for{
		n, _ := fin.Read(buf)
		if 0 == n { break }
	}
	fmt.Println(string(buf))
	*/

	return
}
