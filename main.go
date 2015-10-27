package main

import (
	"./context"
	"./fluent"
)

func Init(ctx *context.Context) {
	client := fluent.NewActionLogClient(ctx)
	ctx.Set("fluent", client)
}

func SomeProc(ctx *context.Context) {
	client := ctx.Get("fluent").(*fluent.ActionLogClient)
	client.AppendActionLog(ActionLog{}, ActionLog{}, ActionLog{})
}

func Success(ctx *context.Context) {
	client := ctx.Get("fluent").(*fluent.ActionLogClient)
	client.Send()
}

func Failure(ctx *context.Context) {
}

func main() {
	ctx := &context.Context{}
	Init(ctx)

	SomeProc(ctx)

	if true {
		// Success
		Success(ctx)
		return
	}
	// False
	Failure(ctx)
}
