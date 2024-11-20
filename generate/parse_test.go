package generate

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"testing"
)

func createPlugin() (*plugin.Plugin, error) {
	var p plugin.Plugin

	content := `info(
    title: "type title here"
    desc: "type desc here"
    author: "type author here"
    email: "type email here"
    version: "type version here"
)

type RegisterReq {
  x map[string]string
y []string
}

service user-api {
    @doc(
        summary: 注册
    )
    @handler register
    post /api/user/register (RegisterReq)
}`

	var info struct {
		ApiFilePath string
		Style       string
		Dir         string
	}
	// err := json.Unmarshal([]byte(content), &info)
	// if err != nil {
	//		return nil, err
	//	}

	p.ApiFilePath = info.ApiFilePath
	p.Style = info.Style
	p.Dir = info.Dir
	api, err := parser.ParseContent(content)
	if err != nil {
		return nil, err
	}

	p.Api = api
	return &p, nil
}

func TestApplyGenerate(t *testing.T) {
	p, err := createPlugin()
	if err != nil {
		t.Fatal(err)
	}

	sw, err := applyGenerate(p, "host", "basePath", "http")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", sw.Definitions["RegisterReq"])
}
