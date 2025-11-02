package writeCode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试写代码函数
func TestWriteCode(t *testing.T) {
	req := &GenerateCodeReq{
		Id:      48,
		Type:    0,
		SysCode: "buildingBlocks",
		Domain:  "master",
	}
	resp, err := writeCodeLogic(req)
	assert.Nil(t, err)
	t.Log(resp)
}
