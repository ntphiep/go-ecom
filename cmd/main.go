package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/ntphiep/go-ecom/cmd/api"
	"github.com/ntphiep/go-ecom/config"
	"github.com/ntphiep/go-ecom/db"

	"log"
)

func main() {
  db, err := db.NewMySQLStorage(mysql.Config{
    User: config.Envs.DBUser,
    Passwd: config.Envs.DBPassword,
    Addr: config.Envs.DBAddress,
    DBName: config.Envs.DBName,
    Net: "tcp",
    AllowNativePasswords: true,
		ParseTime:            true,
  })
  if err != nil {
    log.Fatal(err)
  }


  server := api.NewApiServer(":8080", nil)
  if err := server.Run(); err != nil {
    log.Fatal(err)
  }

}
