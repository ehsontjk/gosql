package app

import (
	"encoding/json"
	"strconv"
	"log"
	"github.com/ehsontjk/gosql/pkg/customers"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
	customersSvc *customers.Service

}
func NewServer(mux *http.ServeMux, customersSvc *customers.Service) *Server {
	return &Server{mux: mux, customersSvc: customersSvc}
}
func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.mux.ServeHTTP(writer, request)
}
func (s *Server) Init(){
	s.mux.HandleFunc("/customers.getAll", s.handleGetAllBanners)
	s.mux.HandleFunc("/customers.getById", s.handleGetCustomerByID)
	s.mux.HandleFunc("/customers.save", s.handleSaveBanner)
	s.mux.HandleFunc("/customers.removeById", s.handleRemoveByID)
}
func (s *Server) handleGetCustomerByID(writer http.ResponseWriter, request *http.Request) {
	idParam := request.URL.Query().Get("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
	log.Print(err)
	http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	return
	}
	item, err := s.customersSvc.ByID(request.Context(), id)
	if errors.Is(err, customers.ErrNotFound) {
		http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	
	data, err := json.Marshal(item)
	if err != nil {
	log.Print(err)
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
	}
	writer.Header().Set("Content-Type","application/json")
	_, err = writer.Write(data)
	if err != nil {
	log.Print(err)
	}
}

}