package main
import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ehsontjk/gosql/pkg/customers"
	"github.com/ehsontjk/gosql/cmd/server/app"
	"database/sql"
	"net"
	"net/http"
	"log"
	"os"
)
func main () {
	host := "0.0.0.0"
	port := "9999"
	dsn := "postgres://app:pass@localhost:5432/db"
	if err := execute(host, port, dsn); err != nil {
		os.Exit(1)
	}
}
func execute(host string, port string, dsn string) (err error) {
		db, err := sql.Open("pgx", dsn)
		if err != nil {
			return err
		}
	
	defer func() {
		if cerr := db.Close(); cerr != nil {
			if err == nil {
				err = cerr
				return	
			}
			log.Println(err)
		}
	}()
	mux := http.NewServeMux()
	customersSvc := customers.NewService(db)
	server := app.NewServer(mux, customersSvc)
	server.Init()
	srv := &http.Server{
		Addr: net.JoinHostPort(host, port),
		Handler: server,
	}
	return srv.ListenAndServe()
}