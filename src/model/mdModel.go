package model

import (
	. "db"
	"github.com/polaris1119/logger"
	"golang.org/x/net/context"
)

type Anote struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Ctime   string `json:"ctime"`
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
