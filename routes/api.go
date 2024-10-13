package routes

import (
	"github.com/gorilla/mux"
	"myproject/http/handlers"
	"myproject/http/middleware"
	"net/http"
)

func BuildApiRouter() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	// user management
	apiUser := api.PathPrefix("/users").Subrouter()

	apiUser.HandleFunc("", handlers.CreateUser).Methods(http.MethodPost)
	apiUser.HandleFunc("", handlers.GetListUser).Methods(http.MethodGet)
	apiUser.HandleFunc("/{id:[0-9]+}", handlers.ShowUser).Methods(http.MethodGet)
	apiUser.HandleFunc("/{id:[0-9]+}", handlers.UpdateUserById).Methods(http.MethodPut)
	apiUser.HandleFunc("/{id:[0-9]+}", handlers.DeleteUserById).Methods(http.MethodDelete)

	// product management
	apiProduct := api.PathPrefix("/products").Subrouter()

	apiProduct.HandleFunc("", handlers.CreateProduct).Methods(http.MethodPost)
	apiProduct.HandleFunc("", handlers.GetListProduct).Methods(http.MethodGet)
	apiProduct.HandleFunc("/{id:[0-9]+}", handlers.ShowProduct).Methods(http.MethodGet)
	apiProduct.HandleFunc("/{id:[0-9]+}", handlers.UpdateProductById).Methods(http.MethodPut)
	apiProduct.HandleFunc("/{id:[0-9]+}", handlers.DeleteProductById).Methods(http.MethodDelete)

	// use middleware
	router.Use(middleware.Auth, mux.CORSMethodMiddleware(router))

	return router
}
