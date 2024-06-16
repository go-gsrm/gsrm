package gsrm

import "github.com/go-gsrm/gsrm/migrate"

func AutoMigrate[T any](db DB, structType T) {
	db.Query(migrate.GenerateTableSQL[T]())
}
