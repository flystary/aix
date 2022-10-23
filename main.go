package main

import (
	"fmt"
	n "aix/net"
	r "aix/net/route"
)

var (
	route  r.Route
	router  n.Router = &route
	service n.Service
	re      n.Requester
)

func init() {
	fileName := "route.yaml"
	path := fmt.Sprintf("C:/aix/%s", fileName)
	route = r.LoadRoute(path)
}

func main() {
	fmt.Println(router.GetTokenFromRoute())
	// service.Router =
	// fmt.Println(re.GetToken(&service))


}
