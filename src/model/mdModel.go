package model

import (
	. "db"
	"github.com/polaris1119/logger"
	"golang.org/x/net/context"
	"time"
)

type Anote struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Commentnum   int       `json:"commentnum"`
	Uid          int       `json:"uid"`
	Newslist_tpl int       `json:"newslist_tpl"`
	Ctime        time.Time `json:"ctime"`
}

// FindById 获取单条博文
func MdModelFindById(ctx context.Context, id string) (*Anote, error) {
	anote := &Anote{}
	_, err := MasterDB.Where("id = ?", id).Get(anote)
	if err != nil {
		logger.Errorln("article logic FindById Error:", err)
	}
	return anote, err
}

//获取列表
func MdModelList(limit int) ([]*Anote, error) {
	anote := make([]*Anote, 0)
	err := MasterDB.Where("1=?", 1).
		OrderBy("id DESC, ctime DESC").Limit(limit).Find(&anote)

	return anote, err
}
