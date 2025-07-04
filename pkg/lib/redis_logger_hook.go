package lib

import (
	"context"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisHook struct {
	logger        *Logger
	slowThreshold time.Duration
}

func NewRedisHook(logger *Logger, slowThreshold time.Duration) *RedisHook {
	return &RedisHook{
		logger:        logger,
		slowThreshold: slowThreshold,
	}
}

func (h *RedisHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		h.logger.WithFields(logrus.Fields{
			"network": network,
			"addr":    addr,
		}).Debug("Dialing Redis server")
		return next(ctx, network, addr)
	}
}

func (h *RedisHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		start := time.Now()

		// 执行前日志
		h.logger.WithContext(ctx).WithFields(logrus.Fields{
			"cmd":  cmd.Name(),
			"args": cmd.Args(),
		}).Debug("Executing Redis command")

		err := next(ctx, cmd)

		// 执行后日志
		duration := time.Since(start)
		fields := logrus.Fields{
			"cmd":      cmd.Name(),
			"args":     cmd.Args(),
			"duration": duration,
		}

		if err != nil && err != redis.Nil {
			fields["error"] = err.Error()
			h.logger.WithContext(ctx).WithFields(fields).Error("Redis command failed")
		} else if duration > h.slowThreshold {
			h.logger.WithContext(ctx).WithFields(fields).Warn("Slow Redis command")
		} else {
			h.logger.WithContext(ctx).WithFields(fields).Debug("Redis command completed")
		}

		return err
	}
}

func (h *RedisHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		h.logger.WithContext(ctx).WithField("count", len(cmds)).Debug("Executing Redis pipeline")
		start := time.Now()

		err := next(ctx, cmds)

		duration := time.Since(start)
		h.logger.WithContext(ctx).WithFields(logrus.Fields{
			"count":    len(cmds),
			"duration": duration,
			"error":    err != nil,
		}).Debug("Redis pipeline completed")

		return err
	}
}
