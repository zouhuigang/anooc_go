package model

import (
	"errors"
	"github.com/gorilla/schema"
)

var schemaDecoder = schema.NewDecoder()

func init() {
	schemaDecoder.SetAliasTag("json")
	schemaDecoder.IgnoreUnknownKeys(true)
}

var NotModifyAuthorityErr = errors.New("没有修改权限")
