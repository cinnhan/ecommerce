package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetId(r *http.Request) int {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	return id
}
