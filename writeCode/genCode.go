package writeCode

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/buildingBlocksCoder/tool/file"
	"github.com/jianyuezhexue/buildingBlocksCoder/tool/http"
	"github.com/jianyuezhexue/buildingBlocksCoder/tool/resp"
)

// 生成代码
func WriteCode(c *gin.Context) {

	// 1. 解析请求参数
	req := &GenerateCodeReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 2. 逻辑实现
	writeResp, err := writeCodeLogic(req)
	if err != nil {
		resp.BizError(c, "生成代码请求失败："+err.Error())
		return
	}

	// 5. 返回结果
	resp.Success(c, writeResp)
}

// writeCodeLogic
func writeCodeLogic(req *GenerateCodeReq) (any, error) {

	logInfo := []string{}

	// 1. 接口拉取数据
	http := http.NewHttp()
	httpUrl := "http://localhost:2400" + "/v1/entity/genCode"
	resBytes, err := http.Post(httpUrl, req)
	if err != nil {
		return nil, errors.New("拉取代码请求失败：" + err.Error())
	}

	// 3. 解析响应结果
	httpResp := &CommonResp[GenerateCodeResp]{}
	err = json.Unmarshal(resBytes, httpResp)
	if err != nil {
		return nil, errors.New("解析代码请求失败：" + err.Error())
	}

	// 接口错误判断
	if httpResp.Code != 0 {
		return nil, errors.New("接口错误：" + httpResp.Msg)
	}

	// 4. 代码生成执行
	if len(httpResp.Data.Codes) == 0 {
		return nil, errors.New("没有生成代码")
	}

	// 获取当前文件路径
	dirPath, _ := os.Getwd()
	for _, itemFileCode := range httpResp.Data.Codes {

		// 完善路径地址
		if len(itemFileCode.FilePath) > 0 && itemFileCode.FilePath[0] == '/' {
			// 如果path以/开头，直接使用
			itemFileCode.FilePath = dirPath + itemFileCode.FilePath
		} else {
			// 如果path不以/开头，添加/
			itemFileCode.FilePath = dirPath + "/" + itemFileCode.FilePath
		}

		// 判断路径是否存在
		if !file.IsExist(itemFileCode.FilePath) {
			// 创建目录
			err := file.Mkdir(itemFileCode.FilePath)
			if err != nil {
				return nil, errors.New("创建目录失败：" + err.Error())
			}
		}

		// 判断是否是前端文件
		// itemFileCode.FileName 中读取 . 后面的文件名
		fileType := ""
		// 通过 . 号找到后面的文件类型名
		dotIndex := strings.LastIndex(itemFileCode.FileName, ".")
		if dotIndex != -1 && dotIndex < len(itemFileCode.FileName)-1 {
			fileType = itemFileCode.FileName[dotIndex+1:]
		}

		// 判断是否是前端文件 | [ts,js,vue]
		isFrountendFile := false
		if fileType == "ts" || fileType == "js" || fileType == "vue" {
			isFrountendFile = true
		}

		// 后端跳过前端文件，前端跳过后端文件
		if req.Type == 0 && isFrountendFile {
			continue
		}
		if req.Type == 1 && !isFrountendFile {
			continue
		}

		// 写入文件
		fullFIlePath := itemFileCode.FilePath + itemFileCode.FileName
		err = file.WriteFile(fullFIlePath, itemFileCode.Content)
		if err != nil {
			return nil, errors.New("写入文件失败：" + err.Error())
		}
		logInfo = append(logInfo, "生成文件："+fullFIlePath)
	}

	// 5. 返回结果
	return logInfo, nil
}
