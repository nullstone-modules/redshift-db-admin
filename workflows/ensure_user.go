package workflows

import (
	"database/sql"
	"log"
	"redshift-db-admin/postgresql"
)

func EnsureUser(db *sql.DB, newUser postgresql.Role) error {
	log.Printf("ensuring user %q\n", newUser.Name)

	return newUser.Ensure(db)
}
