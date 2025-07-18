package databases

import (
	"log"

	"bufferbox_backend_go/configs"
	"bufferbox_backend_go/pkg/utils"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

func NewMSSQLDBConnection(cfg *configs.Configs) (*sqlx.DB, error) {
	mssqlUrl, err := utils.ConnectionUrlBuilder("mssql", cfg)
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("mssql", mssqlUrl)
	if err != nil {
		defer db.Close()
		log.Printf("error, can't connect to database, %s", err.Error())
		return nil, err
	}


	log.Println("MS SQL database has been connected")
	return db, nil
}