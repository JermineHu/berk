package config

var (
	resourceAddr string
	resourcePath string
)

func SetResourceAddr(addr string) {
	resourceAddr = addr
}
func SetResourcePath(path string) {
	resourcePath = path
}

func GetResourceAddr() string {
	return resourceAddr
}
func GetResourcePath() string {
	return resourcePath
}
