package ipinfo

import (
	"errors"
	"fmt"
	"strings"

	"io/ioutil"
	"net/http"

	"github.com/dustin/gojson"
	"gopkg.in/xmlpath.v2"
)

const (
	URL_IP_INFO = "http://ipinfo.io"
)

// IPInfo is the response to ipinfo.io
type IPInfo struct {
	City     string `json:"city"`
	Country  string `json:"country"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Region   string `json:"region"`
}

func (i *IPInfo) Equals(ipInfo *IPInfo) bool {
	if i.IP != ipInfo.IP && i.City != ipInfo.City && i.Country != ipInfo.Country && i.Hostname != ipInfo.Hostname &&
		i.Loc != ipInfo.Loc && i.Org != ipInfo.Org && i.Region != ipInfo.Region {
		return false
	}
	return true
}

// The API for ipinfo.io
type IPInfoAPI struct {
}

func NewIPInfoAPI() *IPInfoAPI {
	return &IPInfoAPI{}
}

// GetInfo, retrieve information about an IP Adress
// @param IP Address to search info
// @return IPInfo struct with the info about the ip
// @return error
func (i *IPInfoAPI) GetInfo(ipAddress string) (*IPInfo, error) {
	url := fmt.Sprintf("%s/%s", URL_IP_INFO, ipAddress)
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	xpath := "//*[@id='content']/section[1]/section[1]/div/pre[6]"
	htmlInfo := string(bytes)

	reader := strings.NewReader(htmlInfo)

	reader = strings.NewReader(htmlInfo)
	xmlroot, xmlerr := xmlpath.ParseHTML(reader)

	if xmlerr != nil {
		return nil, err

	}

	path := xmlpath.MustCompile(xpath)
	infoString, ok := path.String(xmlroot)
	if !ok {
		return nil, errors.New("Not possible retrieve the information")
	}

	infoBytes := []byte(infoString)
	var result IPInfo
	json.Unmarshal(infoBytes, &result)
	return &result, nil
}

// GetASNInfo
func (i *IPInfoAPI) GetASNInfo(asn string) {

}
