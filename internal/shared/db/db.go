package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nicolas-170/Industria-Xpert/cmd/logger"
	"github.com/nicolas-170/Industria-Xpert/config"
)

func NewDb(dbConfig config.ConfigDb) *sql.DB {
	// Nos conectamos a la base de datos
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbConfig.DbUser, dbConfig.DbPassword, dbConfig.DbHost, dbConfig.DbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Fatal(err)
	}
	// Probamos la conexión con Ping
	if err := db.Ping(); err != nil {
		// Cerramos si hubo error
		db.Close()
		logger.Fatal(err)
	}
	logger.Info("¡Se realizo la conexion a la base de datos de manera exitosa!")
	return db
}
