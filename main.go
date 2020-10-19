package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func sendSticker(token string, channel string) string {
	client := &http.Client{}
	var link = "https://discord.com/api/v8/channels/" + channel + "/messages"
	var jsonstr = []byte(`{"content": "", "nonce": "", "tts": false, "sticker_ids": ["748293342357356564"]}`)
	req, err := http.NewRequest("POST", link, bytes.NewBuffer(jsonstr))
	req.Header.Add("Authorization", token)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/0.0.306 Chrome/78.0.3904.130 Electron/7.1.11 Safari/537.36")
	req.Header.Add("via", "1.1 google")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	return string(bodyString)
}
func main() {
	fmt.Print("Token: ")
	var token string
	fmt.Scanln(&token)
	var channel string
	fmt.Print("ChannelID: ")
	fmt.Scanln(&channel)
	for {
		resp := sendSticker(token, channel)
		var dresp map[string]interface{}
		err := json.Unmarshal([]byte(resp), &dresp)
		if err != nil {
			log.Fatal(err)
		}
		if dresp["id"] != nil {
			fmt.Println("Sent a sticker!")
		} else {
			fmt.Println("Rate limit reached...")
			time.Sleep(10 * time.Second)
			fmt.Println("Starting again...")
		}
	}

}
