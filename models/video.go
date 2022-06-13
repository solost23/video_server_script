package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

type Video struct {
	conn         *gorm.DB `gorm:"-"`
	ID           string   `gorm:"id:primary_key"`
	UserID       string   `gorm:"user_id"`
	ClassID      string   `gorm:"class_id"`
	Title        string   `gorm:"title" json:"title" form:"title"`
	Introduce    string   `gorm:"introduce" json:"introduce" form:"introduce"`
	ImageUrl     string   `gorm:"image_url" json:"image_url" form:"image_url"`
	VideoUrl     string   `gorm:"video_url"`
	ThumbCount   int64    `gorm:"thumb_count;default:0"`
	CommentCount int64    `gorm:"comment_count;default:0"`
	DeleteStatus string   `gorm:"delete_status;type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`
	CreateTime   int64    `gorm:"create_time"`
	UpdateTime   int64    `gorm:"update_time"`
}

func NewVideoWithConn(conn *gorm.DB) *Video {
	v := &Video{}
	v.conn = conn
	return v
}

func (v *Video) TableName() string {
	return "video"
}

func (v *Video) Connection() *gorm.DB {
	return v.conn.Table(v.TableName())
}

// 生成唯一UUID
func NewUUID() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Println(err.Error())
		return NewUUID()
	}
	return uuid.String()
}

// 返回当前时间
func GetCurrentTime() int64 {
	return time.Now().Unix()
}
