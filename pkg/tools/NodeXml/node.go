package NodeXml

import "encoding/xml"

type Project struct {
	XMLName xml.Name `xml:"project"`
	Text    string   `xml:",chardata"`
	Node    []struct {
		Text     string `xml:",chardata"`
		Username string `xml:"username,attr"`
		Hostname string `xml:"hostname,attr"`
		Name     string `xml:"name,attr"`
		Tags     string `xml:"tags,attr"`
		Id2      string `xml:"id2,attr"`
		ID       string `xml:"id,attr"`
		Region   string `xml:"region,attr"`
		OsName   string `xml:"osName,attr"`
		OsFamily string `xml:"osFamily,attr"`
	} `xml:"node"`
}

