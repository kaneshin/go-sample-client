package fluent

import (
	"./../context"
)

type ActionLog string

type Client struct {
	ctx *context.Context
}

type ActionLogClient struct {
	*Client
	actionLogs []ActionLog
}

func NewActionLogClient(ctx *context.Context) *ActionLogClient {
	return &ActionLogClient{&Client{ctx}}
}

func (c *ActionLogClient) AppendActionLog(actionLog ...ActionLog) *ActionLogClient {
	c.actionLogs = append(c.actionLogs, actionLog...)
	return c
}

func (c *ActionLogClient) Send() error {
	// Send All c.actionLogs
	c.actionLogs = nil
	return nil
}
