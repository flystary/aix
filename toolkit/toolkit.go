package toolkit

// Toolkit 功能包
type Toolkit interface {
	Colorer
	Tabler
}

// Colorer 颜色
type Colorer interface {
	Red(iput string)   string
	Blue(iput string)  string
	Cyan(iput string)  string
	White(iput string) string
}

// Tabler 表格
type Tabler  interface {
	Basic()
	Markdown()
	NoWhiteSpace()
}
