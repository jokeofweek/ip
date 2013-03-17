package ip

import (
	"encoding/json"
	"fmt"
)

type IpRecord struct {
	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`
	City string
	Ip string
	Lat *string
	Lng *string
}

func getIpRecordFromUrl(url string) (*IpRecord, error) {
	content, err := getContent(url)
	if err != nil {
		return nil, err
	}

	var record IpRecord
	err = json.Unmarshal(content, &record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func GetIpRecord(ip string) (*IpRecord, error) {
	return getIpRecordFromUrl(fmt.Sprintf("http://api.hostip.info/get_json.php?position=true&ip=%s", ip))
}

func GetOwnIpRecord() (*IpRecord, error) {
	return getIpRecordFromUrl("http://api.hostip.info/get_json.php?position=true")
}