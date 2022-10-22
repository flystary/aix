package main

import (
	"fmt"
	n "aix/net"
	r "aix/net/route"
)

var (
	route  r.Route
	router  n.Router = &route
)

func init() {
	fileName := "route.yaml"
	path := fmt.Sprintf("C:/aix/%s", fileName)
	route = r.LoadRoute(path)
}

func main() {

	fmt.Println(route.Modes)

	for _, mode := range route.Modes {
		fmt.Println(router.GetCpeFromRoute(mode))
	}


}
