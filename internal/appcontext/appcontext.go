package appcontext

import (
	"context"

	"git.selly.red/Cashbag-B2B/server-aff/internal/logger"
)

type AppContext struct {
	RequestID string
	TraceID   string
	Logger    *logger.Logger
	Context   context.Context
}

type Fields = logger.Fields

// New ...
func New(ctx context.Context) *AppContext {
	var (
		requestID = generateID()
		traceID   = generateID()
	)

	return &AppContext{
		RequestID: requestID,
		TraceID:   traceID,
		Logger:    logger.NewLogger(logger.Fields{"requestId": requestID, "traceId": traceID}),
		Context:   ctx,
	}
}

// AddLogData ...
func (appCtx *AppContext) AddLogData(fields Fields) {
	appCtx.Logger.AddData(fields)
}
