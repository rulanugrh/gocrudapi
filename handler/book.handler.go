package handler

import (
	"encoding/json"
	"fmt"
	"gocrudapi/entity"
	porthandler "gocrudapi/handler/portHandler"
	portserv "gocrudapi/service/portServ"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type bookstruct struct {
	service portserv.BookInterface
}

func NewBookHandler(service portserv.BookInterface) porthandler.BookInterface {
	return &bookstruct{service: service}
}

func (handler *bookstruct) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To Home Page")
}

func (handler *bookstruct) Create(w http.ResponseWriter, r *http.Request) {
	data, errRead := ioutil.ReadAll(r.Body)

	book := entity.Book{}
	json.Unmarshal(data, &book)

	result, err := handler.service.Create(book)
	if err != nil {
		response := entity.ResponseAPI{
			Code:    http.StatusBadRequest,
			Message: "Cant create book",
			Data:    nil,
		}

		res, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	}
	if errRead != nil {
		log.Println("Cant read request Body")
	}

	response := entity.ResponseAPI{
		Code:    http.StatusOK,
		Message: "Can create book",
		Data:    result,
	}

	res, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (handler *bookstruct) GetAll(w http.ResponseWriter, r *http.Request) {

	result, err := handler.service.GetAll()
	if err != nil {
		response := entity.ResponseAPI{
			Code:    http.StatusBadRequest,
			Message: "Cant find book",
			Data:    nil,
		}

		res, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	}

	response := entity.ResponseAPI{
		Code:    http.StatusOK,
		Message: "Can find book",
		Data:    result,
	}

	res, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (handler *bookstruct) Update(w http.ResponseWriter, r *http.Request) {
	book := entity.Book{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &book)

	IdBook := mux.Vars(r)
	getID := IdBook["id"]
	id, _ := strconv.Atoi(getID)

	result, err := handler.service.Update(uint(id), book)
	if err != nil {
		response := entity.ResponseAPI{
			Code:    http.StatusBadRequest,
			Message: "Cant update book",
			Data:    nil,
		}

		res, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	}

	response := entity.ResponseAPI{
		Code:    http.StatusOK,
		Message: "Can update book",
		Data:    result,
	}

	res, _ := json.Marshal(response)

	fmt.Println(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (handler *bookstruct) Delete(w http.ResponseWriter, r *http.Request) {
	IdBook := mux.Vars(r)
	getID := IdBook["id"]
	id, _ := strconv.Atoi(getID)

	err := handler.service.Delete(uint(id))
	if err != nil {
		response := entity.ResponseAPI{
			Code:    http.StatusBadRequest,
			Message: "Cant delete book",
			Data:    nil,
		}

		res, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	}

	response := entity.ResponseAPI{
		Code:    http.StatusOK,
		Message: "Can delete book",
		Data:    "data successfull deleted",
	}

	res, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func (handler *bookstruct) Get(w http.ResponseWriter, r *http.Request) {
	IdBook := mux.Vars(r)
	getID := IdBook["id"]
	id, _ := strconv.Atoi(getID)

	result, err := handler.service.Get(uint(id))
	if err != nil {
		response := entity.ResponseAPI{
			Code:    http.StatusBadRequest,
			Message: "Cant find book",
			Data:    nil,
		}

		res, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	}

	response := entity.ResponseAPI{
		Code:    http.StatusOK,
		Message: "Can find book",
		Data:    result,
	}

	res, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
