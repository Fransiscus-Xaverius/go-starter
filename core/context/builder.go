package context

import (
	"context"

	"github.com/cde/go-example/core/modules/auth/dto"
	log "github.com/sirupsen/logrus"
)

type (
	Builder struct {
		ctx context.Context
	}
	Logger    struct{}
	RequestId struct{}
	Session   struct{}
)

func NewContextBuilder(ctx context.Context) *Builder {
	return &Builder{ctx: ctx}
}

func (c *Builder) SetRequestId(id string) *Builder {
	c.ctx = context.WithValue(c.ctx, RequestId{}, id)
	return c
}

func (c *Builder) GetRequestId() string {
	if id, ok := c.ctx.Value(RequestId{}).(string); ok {
		return id
	}

	return ""
}

func (c *Builder) SetLogger(entry *log.Entry) *Builder {
	c.ctx = context.WithValue(c.ctx, Logger{}, entry)
	return c
}

func (c *Builder) GetLogger() *log.Entry {
	if logEntry, ok := c.ctx.Value(Logger{}).(*log.Entry); ok {
		return logEntry
	}

	return log.NewEntry(NewLogger())
}

func (c *Builder) SetSession(authResponse *dto.AuthorizeResponse) *Builder {
	c.ctx = context.WithValue(c.ctx, Session{}, authResponse)
	return c
}

func (c *Builder) GetSession() *dto.AuthorizeResponse {
	if session, ok := c.ctx.Value(Session{}).(*dto.AuthorizeResponse); ok {
		return session
	}

	return nil
}

func (c *Builder) Context() context.Context {
	return c.ctx
}
