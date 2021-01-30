package controllers

import (
	"encoding/json"
	"github.com/pgandla/microservices-go/mvc/services"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request)  {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("user id must be a number"))
		return
	}
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}
	jsonVal, _ := json.Marshal(user)
	resp.Write(jsonVal)
}