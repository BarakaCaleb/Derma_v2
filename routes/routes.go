package routes

import (
    "github.com/gorilla/mux"
    "net/http"
    "dermadelight/controllers"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/signup", controllers.SignUp).Methods("POST")
    router.HandleFunc("/signin", controllers.SignIn).Methods("POST")

    api := router.PathPrefix("/api").Subrouter()
    api.Use(controllers.Authenticate)

    api.HandleFunc("/products", controllers.GetProducts).Methods("GET")
    api.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
    api.HandleFunc("/products/{id}", controllers.GetProduct).Methods("GET")
    api.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
    api.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")

    return router
}
