package crontask

import (
	"log"

	"github.com/manhrev/lodas/backend/lodas/internal/repository"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
	"github.com/robfig/cron/v3"
)

type CronTask struct {
	repository *repository.Repository
	cron       *cron.Cron
}

func New(entClient *ent.Client) *CronTask {
	return &CronTask{
		repository: repository.New(entClient),
		cron:       cron.New(),
	}
}

func (t *CronTask) Run() {
	t.cron.AddFunc("* * * * *", func() {
		log.Println("one min")
	})
	t.cron.Start()
}
