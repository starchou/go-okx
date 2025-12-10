package rest

import (
	"github.com/starchou/go-okx/examples"
	rc "github.com/starchou/go-okx/rest"
)

// 敏感信息申请的模拟盘KEY，不确定何时会删除
var TestClient = rc.New("", examples.Auth, nil)
