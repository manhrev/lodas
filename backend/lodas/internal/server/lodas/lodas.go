package lodas

import (
	"github.com/manhrev/lodas/backend/lodas/internal/repository"
	lodas "github.com/manhrev/lodas/backend/lodas/pkg/api"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
)

func NewServer(entClient *ent.Client) lodas.LodasServer {
	return &lodasServer{
		repository: repository.New(entClient),
	}
}

type lodasServer struct {
	repository *repository.Repository
	// Other service client connection, db adapter go here
	lodas.UnimplementedLodasServer
}
