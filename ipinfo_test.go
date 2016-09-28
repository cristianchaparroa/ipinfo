package ipinfo

import (
	"encoding/json"
	"testing"
)

func TestIPInfoAPI_GetInfo(t *testing.T) {

	jsonCaseOne := `{"ip": "199.115.116.7","hostname": "No Hostname","city": "Manassas","region": "Virginia","country": "US","loc": "38.7701,-77.6321","org": "AS30633 Leaseweb USA, Inc.","postal": "20109"}`
	jsonCaseTwo := `{"ip": "52.21.250.237","hostname": "ec2-52-21-250-237.compute-1.amazonaws.com","city": "Seattle","region": "Washington","country": "US","loc": "47.6344,-122.3422","org": "AS14618 Amazon.com, Inc.","postal": "98109"}`

	cases := []struct {
		IP       string
		expected string
	}{
		{"199.115.116.7", jsonCaseOne},
		{"52.21.250.237", jsonCaseTwo},
	}

	api := NewIPInfoAPI()

	for _, test := range cases {
		result, err := api.GetInfo(test.IP)
		if err != nil {
			t.Error(err)
		}
		jsonBytes := []byte(test.expected)
		var assert *IPInfo
		json.Unmarshal(jsonBytes, &assert)

		if !assert.Equals(result) {
			t.Errorf("Expected %v \n but get \n %v", assert, result)
		}
	}

}
