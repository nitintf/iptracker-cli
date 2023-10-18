package ip

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nitintf/iptracker/internal/print"
)

type Ip struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

func ShowIpData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := getIpData(url)

	data := Ip{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response")
	}

	print.Heading("\nDATA FOUND :")

	print.Printf(" IP :%s\n CITY :%s\n REGION :%s\n COUNTRY :%s\n LOCATION :%s\n TIMEZONE:%s\n POSTAL :%s\n", data.IP, data.City, data.Region, data.Country, data.Loc, data.Timezone, data.Postal)

}

func getIpData(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response")
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read the response")
	}

	return responseByte
}
