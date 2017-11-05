package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)

func FindLinks(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("read body err: %s", err)
	}
	return b, nil
}

func main() {
	url := "https://api.weibo.cn"
	body, err := FindLinks(url)
	if err != nil {
		fmt.Fprintf(os.Stdout, "request err: %s\n", err)
	}
	fmt.Printf("response: %s", body)
	// fmt.Println(body)
}

