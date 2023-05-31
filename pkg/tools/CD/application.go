package CD

import "github.com/devops-package/tractorFrame/pkg/tools/Guzzle"

type AppServerStruct struct {
	CmsData
	AppServerDataRow
}

type AppServerDataRow struct {
	Data []struct {
		CreatedAt    string        `json:"created_at"`
		UpdatedAt    string        `json:"updated_at"`
		Owners       []interface{} `json:"owners"`
		Applications []struct {
			ID     int `json:"id"`
			Owners []struct {
				Mobile   string `json:"mobile"`
				Email    string `json:"email"`
				FullName string `json:"full_name"`
				PersonID int    `json:"personId"`
			} `json:"owners"`
			AppName string `json:"app_name"`
		} `json:"applications"`
		CreateTime          string   `json:"create_time"`
		InstanceName        string   `json:"instance_name"`
		InstanceID          string   `json:"instance_id"`
		InstanceType        string   `json:"instance_type"`
		InstanceNetworkType string   `json:"instance_network_type"`
		ZoneID              string   `json:"zone_id"`
		RegionID            string   `json:"region_id"`
		Category            string   `json:"category"`
		Hostname            string   `json:"hostname"`
		InnerIPAddresses    []string `json:"inner_ip_addresses"`
		PublicIPAddresses   []string `json:"public_ip_addresses"`
		CPU                 int      `json:"cpu"`
		Memory              int      `json:"memory"`
		OsName              string   `json:"os_name"`
		OsType              string   `json:"os_type"`
		Disks               []struct {
			Size   int    `json:"size"`
			Type   string `json:"type"`
			Device string `json:"device"`
		} `json:"disks"`
		Env    string      `json:"env"`
		Subenv interface{} `json:"subenv"`
	} `json:"data"`
}

func (cd *CDGuzzle) ServerSearch(key string, alarm bool) (*AppServerStruct, error) {
	uri := "/api/v1/service_instances/"
	req := cd.DoNewRequest("GET", uri)
	if key != "" {
		req.SetParam("Query", "search", key)
	}
	if alarm {
		req.SetParam("Query", "alarm", "true")
	}

	out := Guzzle.RequireOK(cd.NewDoRequest(req))
	f := &AppServerStruct{}
	if err := out.Json(f); err != nil {
		return nil, err
	}
	return f, nil
}
