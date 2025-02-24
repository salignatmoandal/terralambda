package lambda

type Deployer interface {
	Deploy(functionName string, zipPath string) error
	Cleanup() error
}

type Invoker interface {
	Invoke(functionName string, payload []byte) ([]byte, error)
}
