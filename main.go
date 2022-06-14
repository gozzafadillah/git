package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/xlzd/gotp"
)

type respon struct {
	Messages []struct {
		Status string `json:"Status"`
		To     []struct {
			Email       string `json:"Email"`
			MessageID   string `json:"MessageID"`
			MessageHref string `json:"MessageHref"`
		} `json:"To"`
	} `json:"Messages"`
}

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func main() {
	secretLength := 4
	code := gotp.RandomSecret(secretLength)

	var jsonData = []byte(`{
		"Messages":[
				{
						"From": {
								"Email": "gozza15bdg@gmail.com",
								"Name": "Fadillah"
						},
						"To": [
								{
										"Email": "` + "gozzafadillah@gmail.com" + `",
										"Name": "` + "aziz" + `"
								}
						],
						"Subject": "Cek Ombak",
						"TextPart": "Code Generator",
						"HTMLPart": "<center><h2>OTP Code</h2><br /> <b><u>` + code + `</u></b> </center>"
				}
		]
	}`)

	url := "https://api.mailjet.com/v3.1/send"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	basic := base64.StdEncoding.EncodeToString([]byte("177a3a51988d43f5512cf71bff810623" + ":" + "ba69cb7437c1bad179c8af199ba33dd1"))

	req.Header.Add("Authorization", "Basic "+basic)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var result = respon{}
	json.Unmarshal(body, &result)

	fmt.Println(result)
}
