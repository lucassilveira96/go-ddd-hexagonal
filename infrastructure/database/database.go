package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// InitDB inicializa e retorna uma conexão com o banco de dados.
func InitDB() *sql.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE") // Exemplo: "require", "disable"

	if host == "" || port == "" || user == "" || dbname == "" {
		log.Fatal("Variáveis de ambiente para a conexão com o banco de dados não estão todas definidas.")
	}

	// Formatar a string de conexão
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao realizar ping no banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
	return db
}
