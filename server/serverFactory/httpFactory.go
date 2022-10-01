package serverFactory

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	proto "grpc/gen"
	helpers "grpc/helpers/log"
	"grpc/internal/database"
	"grpc/service"
	"log"
	"net/http"
)

type HttpFactory struct {
}

func (sf *HttpFactory) Build() {
	mux := runtime.NewServeMux()

	database.ConnectDatabase()

	err := proto.RegisterUserServiceHandlerServer(context.Background(), mux, service.NewMicroservice())
	helpers.GetLogger().Println("Start Server")
	if err != nil {
		helpers.GetLogger().Errorf("cannot register this service: %s", err.Error())
	}
	log.Fatalln(http.ListenAndServe(":8081", mux))
}
