package main

import (
	"RandomNumberMicroservice/API"
	Handler "RandomNumberMicroservice/API/handler"
	s "RandomNumberMicroservice/Services"
	"context"
)

func main() {
	ctx := context.TODO()
	serv := s.NewRandomService(ctx)
	serv = s.NewLoggingService(serv)
	api := Handler.NewAPIService(serv, s.Logging{})
	API.Run(ctx, *api)
}
