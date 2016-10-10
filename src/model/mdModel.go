package model

import (
	. "db"
	"github.com/polaris1119/logger"
	"github.com/zouhuigang/md2txt"
	"golang.org/x/net/context"
	"net/url"
	"time"
)

type Anote struct {
	Id           int       `json:"id" xorm:"pk autoincr" `
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Commentnum   int       `json:"commentnum"`
	Uid          int       `json:"uid"`
	Newslist_tpl int       `json:"newslist_tpl"`
	Ctime        time.Time `json:"ctime"  xorm:"created"`
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

//添加或修改笔记 schemaDecoder.Decode把form转换成struct格式。
func MdModelAddUpdate(ctx context.Context, form url.Values, username string) (id int, err error) {
	notes := &Anote{}
	err1 := schemaDecoder.Decode(notes, form)
	if err1 != nil {
		logger.Errorln("reading SaveReading error", err1)
		return 0, err1
	}

	title := md2txt.Parse([]byte(notes.Title), md2txt.BASIC)

	notes.Title = string(title)

	if notes.Id != 0 {
		_, err = MasterDB.Id(notes.Id).Update(notes)
	} else {
		_, err = MasterDB.Insert(notes)
	}

	if err != nil {
		logger.Errorln("notes save:", err)
		return
	}

	return notes.Id, nil

}
