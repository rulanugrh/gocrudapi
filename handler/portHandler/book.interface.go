package porthandler

import "net/http"

type BookInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	HomePage(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
