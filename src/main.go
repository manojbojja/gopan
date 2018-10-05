package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"wishfinocr/config"
)

func main() {

	fName := os.Args[1]
	lName := os.Args[2]
	pannumber := os.Args[3]

	const imageUrl =  "https://enqlfa.bn.files.1drv.com/y4m2T6nIcvcyiURp8Xi7xegcDKQUgp-AuvhSGQwPlL8nXul7OTxIgvGPVVW-ypNYL829mWJNKj7ChdELUVPBm_KSSzUvTXJNgoOJXfEDCP-VW-DtaZt6KsWXHs7PdxBqqoo0Z4PDyB2BDiosVs1yw0XZOhaSgCy0PPqWoJReFY9Q8v4SaHX5grLj_hMUHXVeoQ37z5kS0tFwgjQ-ykZaCYFVA?width=780&height=1040&cropmode=none"
	//const imageUrl = "https://t69f7w.bn.files.1drv.com/y4mjTgU2vBarx1apwbDbDItSsWxOxmwEtANRVXTR-6GnetmqTbDDIRx_mYqQ9GkVuczwcLE4LeNPcrF4wdSMqKAOPIoVqWec4_PKdWZLfaWbTsrzgr1dqHuWq7WaBP7jzcRa9eLOui_MYpqmUB4_lLthb3SxcTzZiqXqPJkcNQ231O5_9uwk9haRYqX4qrV9yHZ-6rmUAWnxNjBuo1qrxtjjA?width=720&height=464&cropmode=none"
	const uri = config.BASE_URL + config.PARAMS
	const imageUrlEnc = "{\"url\":\"" + imageUrl + "\"}"
	reader := strings.NewReader(imageUrlEnc)
	client := &http.Client{
		Timeout: time.Second * 20,
	}
	req , err := http.NewRequest("POST", uri, reader)
	if err != nil{
		panic(err)
	}

	// Add headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", config.SUB_KEY)

	// Send the request and retrieve the response
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Read the response body.
	// Note, data is a byte array
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var f interface{}
	json.Unmarshal(data, &f)


	jsonFormatted, _ := json.MarshalIndent(f, "", "  ")
	count := 0

	jsonString := strings.ToLower(string(jsonFormatted))
	if strings.Contains(jsonString, "income"){
		count = count + 1
	}

	if strings.Contains(jsonString, "tax"){
		count = count + 1
	}

	if strings.Contains(jsonString, fName){
		count = count + 1
	}

	if strings.Contains(jsonString, lName){
		count = count + 1
	}

	if strings.Contains(jsonString, pannumber){
		count = count + 5
	}

	if count > 4 {
		fmt.Println("verified")
	}

	if count > 2 && count < 4{
		fmt.Println("Need mannual verification")
	}

	if count <= 2{
		fmt.Println("Picture not clear or invalid pancard")
	}
}