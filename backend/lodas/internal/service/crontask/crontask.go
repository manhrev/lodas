package crontask

import (
	"context"
	"log"

	"github.com/manhrev/lodas/backend/lodas/internal/repository"
	"github.com/manhrev/lodas/backend/lodas/internal/service/crawler"
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
	t.cron.AddFunc("0 18 * * *", t.dailyJob)
	t.cron.AddFunc("0 19 * * *", t.dailyJob)
	t.cron.AddFunc("0 20 * * *", t.dailyJob)
	t.cron.Start()
	//time.AfterFunc(time.Duration(1*time.Second), func() {})
}

func (t *CronTask) dailyJob() {
	ctx := context.Background()

	// get daily result
	results := crawler.Crawl()
	err := t.repository.Result.CreateDaily(ctx, results)
	if err != nil {
		log.Printf("Error while creating daily result: %v", err)
	}
	//Check result daily
	err = t.repository.Result.CheckResultDaily(ctx)
	if err != nil {
		log.Printf("Error while checking daily result: %v", err)
	}
}
