package CMDB

import (
	"bytes"
	"fmt"
	"github.com/devops-package/tractorFrame/pkg/tools/Guzzle"
	"time"
)

type DNSCreate struct {
	BkInstName string `json:"bk_inst_name"`
	RecordID   string `json:"record_id"`
	Value      string `json:"value"`
	TTL        int64  `json:"ttl" xml:"ttl"`
	Remark     string `json:"remark" xml:"remark"`
	DomainName string `json:"domain_name" xml:"domain_name"`
	RR         string `json:"rr" xml:"rr"`
	Priority   int64  `json:"priority" xml:"priority"`
	RecordId   string `json:"record_id" xml:"record_id"`
	Status     string `json:"status" xml:"status"`
	Locked     bool   `json:"locked" xml:"locked"`
	Weight     int    `json:"weight" xml:"weight"`
	Line       string `json:"line" xml:"line"`
	Type       string `json:"type" xml:"type"`
	SSLPort    int    `json:"ssl_port" xml:"ssl_port"`
}

type DNSResponse struct {
	Result      bool        `json:"result"`
	BkErrorCode int         `json:"bk_error_code"`
	BkErrorMsg  string      `json:"bk_error_msg"`
	Permission  interface{} `json:"permission"`
	Data        struct {
		Count int `json:"count"`
		Info  []struct {
			BkInstID          int         `json:"bk_inst_id"`
			BkInstName        string      `json:"bk_inst_name"`
			BkObjID           string      `json:"bk_obj_id"`
			BkSupplierAccount string      `json:"bk_supplier_account"`
			CreateTime        time.Time   `json:"create_time"`
			DomainName        string      `json:"domain_name"`
			IsDeleted         string      `json:"is_deleted"`
			IsExpired         string      `json:"is_expired"`
			IsMonitored       interface{} `json:"is_monitored"`
			LastTime          time.Time   `json:"last_time"`
			Line              string      `json:"line"`
			Locked            bool        `json:"locked"`
			Priority          interface{} `json:"priority"`
			RecordID          string      `json:"record_id"`
			Remark            string      `json:"remark"`
			Rr                string      `json:"rr"`
			SslExpireDate     string      `json:"ssl_expire_date"`
			SslExpireStatus   string      `json:"ssl_expire_status"`
			SslIsMonitor      string      `json:"ssl_is_monitor"`
			SslPort           string      `json:"ssl_port"`
			SslPortMonitor    string      `json:"ssl_port_monitor"`
			SslType           string      `json:"ssl_type"`
			Status            string      `json:"status"`
			TTL               interface{} `json:"ttl"`
			Type              string      `json:"type"`
			Value             string      `json:"value"`
			Vendor            string      `json:"vendor"`
			Weight            interface{} `json:"weight"`
		} `json:"info"`
	} `json:"data"`
}

func (cd *CMDBGuzzle) SearchDNS(field, key string) (*DNSResponse, error) {
	uri := "/api/v3/find/instassociation/object/dns"
	req := cd.DoNewRequest("POST", uri)
	req.SetParam("Header", "Bk_user", "migrate")
	req.SetParam("Header", "Http_blueking_supplier_id", "0")
	req.SetParam("Header", "Content-Type", "application/json")
	payload := fmt.Sprintf(
		`{"condition": {"dns": [{"field": "%s", "operator": "$eq", "value": "%s"} ] }, "page": {"start": 0, "sort": "%s"} }`,
		field, key, field)
	jsonStr := []byte(payload)
	req.SetBody(bytes.NewBuffer(jsonStr))
	out := Guzzle.RequireOK(cd.NewDoRequest(req))
	f := &DNSResponse{}
	if err := out.Json(f); err != nil {
		return nil, err
	}
	return f, nil
}

func (cd *CMDBGuzzle) CreateDNS() {
	uri := "/api/v3/create/instance/object/dns"
	req := cd.DoNewRequest("POST", uri)
	req.SetParam("Header", "Bk_user", "migrate")
	req.SetParam("Header", "Http_blueking_supplier_id", "0")
	req.SetParam("Header", "Content-Type", "application/json")

}
