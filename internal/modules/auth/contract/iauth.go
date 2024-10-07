package contract

import (
	"net/http"
)

type IAuth interface {
	Registration(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}
