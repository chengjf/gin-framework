package lib

import (
	"context"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

// LogrusLogger 实现 gorm.Logger.Interface
type LogrusLogger struct {
	*Logger
}

func (l *LogrusLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l // 可以根据需要实现不同日志级别
}

func (l *LogrusLogger) Info(ctx context.Context, msg string, data ...any) {
	l.Logger.WithContext(ctx).Infof(msg, data...)
}

func (l *LogrusLogger) Warn(ctx context.Context, msg string, data ...any) {
	l.Logger.WithContext(ctx).Warnf(msg, data...)
}

func (l *LogrusLogger) Error(ctx context.Context, msg string, data ...any) {
	l.Logger.WithContext(ctx).Errorf(msg, data...)
}

func (l *LogrusLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	var reqID string
	if ginCtx, ok := ctx.(*gin.Context); ok {
		reqID = requestid.Get(ginCtx)
	}
	fields := logrus.Fields{
		"duration": time.Since(begin),
		"rows":     rows,
		"req_id":   reqID,
	}
	if err != nil {
		fields["error"] = err
		l.Logger.WithContext(ctx).WithFields(fields).Error(sql)
	} else {
		l.Logger.WithContext(ctx).WithFields(fields).Debug(sql)
	}
}
