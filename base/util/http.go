package util

import (
	//	"fmt"
	"io/ioutil"
	"net/http"
)

/*
func GetHttp(url string) (body []byte, err error) {
	fmt.Printf("http get:[%s]\n", url)
	//resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/menu/get?access_token=lzbH7J2kclhM1kduLf61Ui209T_aA4aVXEDf_ttfmVKKwbtwl5kf7h28qUqOVzAfPIozpCZ7gYxdUWo1k6qX7Ob7aTH-gsb-B4oCa7jWoneLSjY2O_11D74sNk-lWvPHBRXdABAIYE")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http Get failed")
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll failed")
		return nil, err
	}
	fmt.Println(string(body))
	return body, nil

}
*/
func GetHttp(url string) (body []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
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
