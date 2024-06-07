package cookiemonster

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func SaveAsSQL(cookies []Cookie, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	sqlQuery := ` CREATE TABLE cookies (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"host" TEXT,
		"name" TEXT,
		"value" TEXT,
		"path" TEXT,
		"is_secure" BOOLEAN,
		"is_httponly" BOOLEAN,
		"creation_utc" INTEGER,
		"expiry_utc" INTEGER
	);`

	_, err = db.Exec(sqlQuery)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	insertQuery := `INSERT INTO cookies 
	(host, name, value, path, is_secure, is_httponly, creation_utc, expiry_utc) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	statement, err := db.Prepare(insertQuery)
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement: %w", err)
	}
	defer statement.Close()

	for _, cookie := range cookies {
		_, err = statement.Exec(
			cookie.Host, cookie.Name, cookie.Value, cookie.Path,
			cookie.IsSecure, cookie.IsHttpOnly, cookie.CreationUtc, cookie.ExpiryUtc,
		)
		if err != nil {
			return fmt.Errorf("failed to insert cookie: %w", err)
		}
	}

	log.Println("[+] Cookies saved to", path)
	return nil
}
