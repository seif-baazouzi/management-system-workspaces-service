package db

import (
	"fmt"
	"os"
)

func Migrations() {
	conn := GetPool()
	defer ClosePool(conn)

	_, err := conn.Exec(
		`CREATE TABLE IF NOT EXISTS workspaces (
			workspaceID UUID PRIMARY KEY,
			workspace VARCHAR NOT NULL, 
			userID UUID NOT NULL,
			parentWorkspace UUID,
			createdAt TIMESTAMP DEFAULT NOW() 
		)`,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}
