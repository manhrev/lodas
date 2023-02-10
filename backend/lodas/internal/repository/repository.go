package repository

import (
	"github.com/manhrev/lodas/backend/lodas/internal/repository/record"
	"github.com/manhrev/lodas/backend/lodas/internal/repository/sheet"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
)

type Repository struct {
	entClient *ent.Client

	Sheet  sheet.Sheet
	Record record.Record
}

func New(entClient *ent.Client) *Repository {
	return &Repository{
		entClient: entClient,

		Sheet:  sheet.New(entClient),
		Record: record.New(entClient),
	}
}
