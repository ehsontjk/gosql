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
func NewServer(mux *http.ServeMux, customersSvc *banners.Service) *Server {
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
func (s *Server) handleGetBannerByID(writer http.ResponseWriter, request *http.Request) {
	idParam := request.URL.Query().Get("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
	log.Print(err)
	http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	return
	}
	item, err := s.customersSvc.ByID(request.Context(), id)
	if err != nil {
	log.Print(err)
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
func (s *Server) handleSaveBanner(writer http.ResponseWriter, request *http.Request) {
	idParam := request.PostFormValue("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
	log.Print(err)
	http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	return
	}
	var banner = &banners.Banner{
	ID: id,
	Title: request.PostFormValue("title"),
	Content: request.PostFormValue("content"),
	Button:	request.PostFormValue("button"),
	Link:	request.PostFormValue("link"),
	Image:	"",
	}
	item, err := s.bannersSvc.Save(request, banner)
	if err != nil {
	log.Print(err)
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
func (s *Server) handleGetAllBanners(writer http.ResponseWriter, request *http.Request) {
	items, _ := s.bannersSvc.All(request.Context())
	data, err := json.Marshal(items)
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
func (s *Server) handleRemoveByID(writer http.ResponseWriter, request *http.Request) {
	idParam := request.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam,10,64)
	if err != nil {
	log.Print(err)
	http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	return
	}
	item, err := s.bannersSvc.RemoveByID(request.Context(), id)
	if err != nil {
	log.Print(err)
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