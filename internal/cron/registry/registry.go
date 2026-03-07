package registry

import (
	"context"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/process"
	"github.com/MediStatTech/biometric-service/internal/cron/metric_generator"
	"github.com/MediStatTech/logger"
)

type CronRegistry struct {
	log            *logger.Logger
	metricGenJob   *metric_generator.MetricGeneratorCronJob
	cancel         context.CancelFunc
}

type Options struct {
	Log     *logger.Logger
	Process *process.Interactor
}

func New(opts *Options) *CronRegistry {
	metricGenJob := metric_generator.New(&metric_generator.Options{
		Log:      opts.Log,
		Interval: 1 * time.Second,
		Process:  opts.Process,
	})

	return &CronRegistry{
		log:          opts.Log,
		metricGenJob: metricGenJob,
	}
}

func (r *CronRegistry) Start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	r.cancel = cancel

	go r.metricGenJob.Start(ctx)

	r.log.Info("[Cron Registry] Scheduler started", map[string]any{})
}

func (r *CronRegistry) Shutdown() {
	if r.cancel != nil {
		r.cancel()
	}
	r.log.Info("[Cron Registry] Scheduler stopped", map[string]any{})
}
