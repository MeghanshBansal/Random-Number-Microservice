package Handler

import (
	"RandomNumberMicroservice/Services"
	"RandomNumberMicroservice/Structures"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type APIService struct {
	serv   Service.Service
	Logger Service.Logging
}

func NewAPIService(serv Service.Service, log Service.Logging) *APIService {
	return &APIService{
		serv:   serv,
		Logger: log,
	}
}

func (a *APIService) HandleGetRandomNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		a.HandleGetNRandomNumbers(w, r)
		return
	}
	ctx := context.Background()
	response := "Number is " + strconv.Itoa(a.serv.GetRandomNumber(ctx))
	a.responseJson(ctx, w, http.StatusOK, response)

}

func (a *APIService) HandleGetNRandomNumbers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var body Types.RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		a.Logger.Fatal(ctx, err, "Unable to parse request")
	}
	nums := a.serv.GetNRandomNumbers(ctx, body.N)
	response := "Numbers are " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), ","), "[]")
	a.responseJson(ctx, w, http.StatusOK, response)

}

func (a *APIService) responseJson(ctx context.Context, w http.ResponseWriter, s int, v any) {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		a.Logger.Fatal(ctx, err, "Can not write json")
	}
}
