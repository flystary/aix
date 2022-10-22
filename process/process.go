package process



// Processer 数据处理
type Processer interface {
	Searcher
	Checker
	Displayer
	Remoter
	Execer
}

// Searcher  查询
type Searcher  interface{}

// Checker   检查
type Checker  interface{}

// Displayer 展示
type Displayer interface{}

// Remoter   连接
type Remoter   interface{}

// Execer    远程执行
type Execer    interface{}