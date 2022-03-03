package acc

import (
	"github.com/stretchr/testify/require"
	"os"
	"redshift-db-admin/postgresql"
	"testing"
)

func TestRole(t *testing.T) {
	if os.Getenv("ACC") != "1" {
		t.Skip("Set ACC=1 to run e2e tests")
	}

	db := createDb(t)
	defer db.Close()

	role := postgresql.Role{
		Name:     "role-test-user",
		Password: "role-test-password",
	}
	require.NoError(t, role.Create(db), "unexpected error")

	find := &postgresql.Role{Name: "role-test-user"}
	require.NoError(t, find.Read(db), "read user")
}
