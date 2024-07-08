package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/lib/pq" // Driver de conexão com o Postgree
)

// Conectar abre a conexão com o banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.StringConexaoBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
