package server

import (
	"context"
	"video_server_script/models"
	"video_server_script/scheduler/video_import"
)

func Run() {
	ctx := context.Background()
	action := video_import.NewActionWithCtx(ctx)
	action.SetMysqlConn(models.NewMysqlConnect())
	if err := action.Deal(ctx); err != nil {
		panic(err)
	}
}
