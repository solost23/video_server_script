package video_import

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/solost23/tools/readFile"
	"gorm.io/gorm"
	"video_server_script/models"
)

type Action struct {
	conn *gorm.DB
	ctx  context.Context
}

func NewActionWithCtx(ctx context.Context) *Action {
	r := &Action{}
	r.ctx = ctx
	return r
}

func (a *Action) SetMysqlConn(conn *gorm.DB) {
	a.conn = conn
}

func (a *Action) GetMysqlConn() *gorm.DB {
	return a.conn
}

func (a *Action) Deal(_ context.Context) (err error) {
	// 解析excel文件
	// 写入数据库
	file := readFile.ReadFile{
		FileName: "this_vid.xlsx",
		FileType: "excel",
	}
	fileContent, err := file.ReadFile()
	if err != nil {
		panic(err)
	}
	videoInfoList := make([]models.Video, 0, 10000)
	for index, value := range fileContent {
		if index == 0 {
			continue
		}
		// 生成对象，存入切片
		videoInfoList = append(videoInfoList, a.buildRequest(value))
	}
	// 数据入库
	tx := models.NewVideoWithConn(a.GetMysqlConn()).Connection().Begin()
	for index, videoInfo := range videoInfoList {
		if err = tx.Create(&videoInfo).Error; err != nil {
			tx.Rollback()
			return err
		}
		videoInfoJson, _ := json.Marshal(videoInfo)
		fmt.Printf("第%d条数据插入成功, data:%s \n", index+1, string(videoInfoJson))
	}
	tx.Commit()
	fmt.Println("数据存储完毕")
	return err
}

func (a *Action) buildRequest(data []string) models.Video {
	return models.Video{
		ID:           models.NewUUID(),
		UserID:       "0",
		ClassID:      "0",
		Title:        data[0],
		Introduce:    data[1],
		ImageUrl:     data[2],
		VideoUrl:     data[3],
		ThumbCount:   0,
		CommentCount: 0,
		DeleteStatus: "DELETE_STATUS_NORMAL",
		CreateTime:   models.GetCurrentTime(),
		UpdateTime:   models.GetCurrentTime(),
	}
}
