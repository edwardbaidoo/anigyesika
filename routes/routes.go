package routes

import (
	"github.com/gorilla/mux"
)

//import controllers

var Routes = func(router *mux.Router) {
	router.HandleFunc("/user")
	router.HandleFunc("/user/register")
	router.HandleFunc("/user/login")
	router.HandleFunc("/transaction/create")

}
