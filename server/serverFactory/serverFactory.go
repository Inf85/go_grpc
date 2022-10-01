package serverFactory

type serverFbInterface interface {
	Build()
}

var serverFactoryBuilderMap = map[string]serverFbInterface{
	"http": &HttpFactory{},
	"grpc": &GrpcFactory{},
}

func GetServerFactoryBuilder(key string) serverFbInterface {
	return serverFactoryBuilderMap[key]
}
