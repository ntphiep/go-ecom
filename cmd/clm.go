package main

import (
	"database/sql"
	"fmt"

	"github.com/ntphiep/go-ecom/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func (s *APIServer) testClm() {
	userStore := user.NewStore(s.db)
	fmt.Printf("Type of userStore: %T\n", userStore)
}

func main() {
	server := APIServer{}
	server.testClm()
}
