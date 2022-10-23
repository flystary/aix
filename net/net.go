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

// Requester 请求
type Requester interface{
	GetToken(service *Service) *Service
	// GetMode(service  *Service) *Service
	// GetBytes(service *Service) *Service
}


// Serializer 序列化
type Serializer interface{}