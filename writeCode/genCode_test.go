package writeCode

import (
	"path/filepath"
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
	resp, err := WriteCodeLogic(req)
	assert.Nil(t, err)
	t.Log(resp)
}

// TestPathJoin 测试路径拼接逻辑
func TestPathJoin(t *testing.T) {
	tests := []struct {
		name     string
		dirPath  string
		filePath string
		want     string
	}{
		{
			name:     "相对路径回退两级",
			dirPath:  "/Users/jie/projects/codegen",
			filePath: "../../projects/",
			want:     "/Users/jie/projects",
		},
		{
			name:     "相对路径回退一级",
			dirPath:  "/Users/jie/projects/codegen",
			filePath: "../test/",
			want:     "/Users/jie/projects/test",
		},
		{
			name:     "当前目录相对路径",
			dirPath:  "/Users/jie/projects/codegen",
			filePath: "./src/",
			want:     "/Users/jie/projects/codegen/src",
		},
		{
			name:     "普通相对路径",
			dirPath:  "/Users/jie/projects/codegen",
			filePath: "src/controllers/",
			want:     "/Users/jie/projects/codegen/src/controllers",
		},
		{
			name:     "绝对路径",
			dirPath:  "/Users/jie/projects/codegen",
			filePath: "/opt/app/",
			want:     "/opt/app",
		},
		{
			name:     "复杂的相对路径",
			dirPath:  "/Users/jie/projects/codegen/backend",
			filePath: "../../frontend/src/",
			want:     "/Users/jie/projects/frontend/src",
		},
		{
			name:     "空相对路径",
			dirPath:  "/Users/jie/projects/codegen",
			filePath: "",
			want:     "/Users/jie/projects/codegen",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result string
			if len(tt.filePath) > 0 && tt.filePath[0] == '/' {
				// 如果path以/开头，直接使用绝对路径
				result = filepath.Clean(tt.filePath)
			} else {
				// 相对路径，拼接后使用filepath.Clean规范化，自动处理../和./
				result = filepath.Clean(filepath.Join(tt.dirPath, tt.filePath))
			}

			if result != tt.want {
				t.Errorf("路径拼接错误:\n  dirPath:  %s\n  filePath: %s\n  got:      %s\n  want:     %s",
					tt.dirPath, tt.filePath, result, tt.want)
			}
		})
	}
}
