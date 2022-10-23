package cmd


import (
	n "aix/net"
)

type Cmd struct {
	Sns		[]string //SN列表
	Mode    string   //版本
	// Enterprise
	Entn    string   //企业号
}


func (c *Cmd) GetToken(service n.Service) 