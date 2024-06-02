package cookiemonster

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Cookie struct {
	Host        string
	Name        string
	Value       string
	Path        string
	IsSecure    bool
	IsHttpOnly  bool
	CreationUtc int64
	ExpiryUtc   int64
}

func ParseCookies(cookiesFile string) ([]Cookie, error) {
	db, err := sql.Open("sqlite", cookiesFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open cookies file: %w", err)
	}
	defer db.Close()

	query := "SELECT host_key, name, encrypted_value, path, is_secure, is_httponly, creation_utc, expires_utc FROM cookies"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query cookies: %w", err)
	}
	defer rows.Close()

	var cookies []Cookie
	for rows.Next() {
		var cookie Cookie
		if err := rows.Scan(&cookie.Host, &cookie.Name, &cookie.Value, &cookie.Path, &cookie.IsSecure, &cookie.IsHttpOnly, &cookie.CreationUtc, &cookie.ExpiryUtc); err != nil {
			return nil, fmt.Errorf("failed to scan cookie: %w", err)
		}

		cookies = append(cookies, cookie)

	}

	return cookies, nil
}
