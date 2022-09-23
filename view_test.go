package gojenkins

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initJenkins(ctx context.Context) *Jenkins {
	j := CreateJenkins(nil, "http://47.96.36.207:8088", "admin", "onceas")
	_, _ = j.Init(ctx)
	return j
}

func TestView_GetViewXML(t *testing.T) {
	ctx := context.Background()
	j := initJenkins(ctx)
	view, err := j.GetView(ctx, "testview")
	assert.Nil(t, err)

	str, err := view.GetConfigXML(ctx)
	assert.Nil(t, err)
	fmt.Println(str)
}
