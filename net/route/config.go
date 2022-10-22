package route

// Route 路由
type Route struct {
	InitURL   string   `yaml:"initurl"`
	Tokenurl  string   `yaml:"tokenurl"`
	Operation string   `yaml:"operation"`
	Modes     []string `yaml:"modes"`

	Valor     valor     `yaml:"valor"`
	Nexus     nexus     `yaml:"nexus"`
	Watsons   watsons   `yaml:"watsons"`
	Tassadar  tassadar  `yaml:"tassadar"`
	WatsonsHa watsonsha `yaml:"watsonsha"`
}

// Domain 信息类别
type Domain struct {
	Cpe string `yaml:"cpe"`
	Pop string `yaml:"pop"`
	Dvc string `yaml:"dvc"`
}

// Valor sdwan7
type valor Domain

// Nexus sdwan6
type nexus Domain

// Watsons 屈臣氏
type watsons Domain

// Tassadar sase2
type tassadar Domain

// WatsonsHa 屈臣氏仓库
type watsonsha Domain

