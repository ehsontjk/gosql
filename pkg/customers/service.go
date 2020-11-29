package customers

import (
	"strconv"
	"os"
	"net/http"
	"errors"
	"context"
	"encoding/json"
	"database/sql"
	"path/filepath"
	"io"
)
var ErrNotFound = errors.New("item not found")
var ErrInternal = errors.New("internal error")
type Service struct {
	db 	*sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}



type Customer struct {
	ID int64  'json:"id"'
	Name string  'json:"name"'
	Phone string  'json:"phone"'
	Active bool  'json:"active"'
	Created time.Time  'json:"created"'
	
}


func (s *Service) ByID(ctx context.Context, id int64) (*Customer, error) {
	item := &Customer{}
	err := s.db.QueryRowContext(ctx, SELECT id, name, phone, active, created FROM customers WHERE id = $1, id).Scan(&item.ID, &item.Name, &item.Phone, &item.Active, &item.Created)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		log.Println(err)
		return nil, ErrInternal
	}
	return item, nil
}


