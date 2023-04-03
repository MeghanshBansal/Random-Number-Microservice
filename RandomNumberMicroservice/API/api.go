package API

import (
	Handler "RandomNumberMicroservice/API/handler"
	"context"
	"net/http"
)

func Run(ctx context.Context, api Handler.APIService) {
	http.HandleFunc("/", api.HandleGetRandomNumber)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		api.Logger.Fatal(ctx, err, "Unable to start the server")
	}
}
