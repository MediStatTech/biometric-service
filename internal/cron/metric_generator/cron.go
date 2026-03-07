package metric_generator

import (
	"context"
	"time"

	"github.com/MediStatTech/biometric-service/internal/app/biometric/usecases/process"
	"github.com/MediStatTech/logger"
)

const jobName = "metric.generator"

type MetricGeneratorCronJob struct {
	log      *logger.Logger
	interval time.Duration
	process  *process.Interactor
}

type Options struct {
	Log      *logger.Logger
	Interval time.Duration
	Process  *process.Interactor
}

func New(opts *Options) *MetricGeneratorCronJob {
	interval := opts.Interval
	if interval == 0 {
		interval = 1 * time.Second
	}

	return &MetricGeneratorCronJob{
		log:      opts.Log,
		interval: interval,
		process:  opts.Process,
	}
}

func (j *MetricGeneratorCronJob) Start(ctx context.Context) {
	ticker := time.NewTicker(j.interval)
	defer ticker.Stop()

	j.log.Info("[MetricGenerator Cron] Started", map[string]any{
		"job_name": jobName,
		"interval": j.interval.String(),
	})

	for {
		select {
		case <-ctx.Done():
			j.log.Info("[MetricGenerator Cron] Stopped", map[string]any{
				"job_name": jobName,
			})
			return
		case <-ticker.C:
			if err := j.process.Execute(ctx); err != nil {
				j.log.Error("[MetricGenerator Cron] Failed to generate metrics", map[string]any{
					"job_name": jobName,
					"error":    err.Error(),
				})
			}
		}
	}
}
