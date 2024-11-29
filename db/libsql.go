package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tursodatabase/go-libsql"
)

func GetConn() {
  dbName := "maindb.db"
  dbURL := "libsql://maindb-jpradass.turso.io"
  authToken := os.Getenv("WHATTOEAT_DB_AUTH") 

  dir, err := os.MkdirTemp("", "libsql-*")
  if err != nil {
    fmt.Println("Error creating temporary directory:", err)
    os.Exit(1)
  }

  dbPath := filepath.Join(dir, dbName)
  connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, dbURL, libsql.WithAuthToken(authToken))
  if err != nil {
    fmt.Println("Error creating temporary directory:", err)
    os.Exit(1)
  }

  defer connector.Close()

  db := sql.OpenDB(connector)
  defer db.Close()

  rows, err := db.Query("SELECT * FROM recipes");
  if err != nil {
    fmt.Println("Error creating temporary directory:", err)
    os.Exit(1)
  }

  for rows.Next() {
    var id, name, desc string 
    if err := rows.Scan(&id, &name, &desc); err != nil {
      fmt.Println("Error creating temporary directory:", err)
      os.Exit(1)
    }

    fmt.Printf("id: %s", id)
  }
}
