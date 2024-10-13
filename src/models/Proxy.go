package models

type ProxyObj struct {
	//Id       string `json:"id,omitempty"`
	Ip       string `json:"Ip,omitempty"`
	Port     uint64 `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Proxy    string `json:"proxy,omitempty"`
}
