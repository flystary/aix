package net


// Neter 网络
type Neter interface {
	Router
	Requester
	Serializer
}

// Router 路由
type Router interface {
	GetCpeFromRoute(mode string) string
	GetPopFromRoute(mode string) string
	GetDvcFromRoute(mode string) string
	GetOperationFromRoute() string
	GetTokenFromRoute() string
}

// Requester 登录
type Requester interface{
	GetToken(URL string) string
	GetBytes()
}

// Serializer 序列化
type Serializer interface{}