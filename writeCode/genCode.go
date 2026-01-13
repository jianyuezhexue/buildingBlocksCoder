package writeCode

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
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
	writeResp, err := WriteCodeLogic(req)
	if err != nil {
		resp.BizError(c, "生成代码请求失败："+err.Error())
		return
	}

	// 5. 返回结果
	resp.Success(c, writeResp)
}

// writeCodeLogic
func WriteCodeLogic(req *GenerateCodeReq) (any, error) {

	logInfo := []string{}

	// 1. 接口拉取数据
	http := http.NewHttp()
	httpUrl := "http://8.155.47.77:2400" + "/v1/entity/genCode"
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

		// 完善路径地址
		if len(itemFileCode.FilePath) > 0 && itemFileCode.FilePath[0] == '/' {
			// 如果path以/开头，直接使用绝对路径
			itemFileCode.FilePath = filepath.Clean(itemFileCode.FilePath)
		} else {
			// 相对路径，拼接后使用filepath.Clean规范化，自动处理../和./
			itemFileCode.FilePath = filepath.Clean(filepath.Join(dirPath, itemFileCode.FilePath))
		}

		// 判断路径是否存在
		if !file.IsExist(itemFileCode.FilePath) {
			// 创建目录
			err := file.Mkdir(itemFileCode.FilePath)
			if err != nil {
				return nil, errors.New("创建目录失败：" + err.Error())
			}
		}

		// 拼接完成路径
		fullFIlePath := itemFileCode.FilePath + itemFileCode.FileName

		// 代码生成.存在跳过
		if itemFileCode.WriteType == 2 {

			// 存在跳过
			if file.IsExist(fullFIlePath) {
				continue
			}

			// 不存在，将WriteType改成1,方便执行写入
			itemFileCode.WriteType = 0
		}

		// 代码生成.进行替换
		if itemFileCode.WriteType == 1 {

			// 0.校验文件是否存在，不存在跳过
			if !file.IsExist(fullFIlePath) {
				continue
			}

			// 1. 读取文件内容
			content, err := file.ReadFile(fullFIlePath)
			if err != nil {
				return nil, errors.New("读取文件失败：" + err.Error())
			}

			// 2. 替换文件内容
			// 拼接文件内容 + \n + 替换标识
			itemFileCode.Content = itemFileCode.Content + "\n" + itemFileCode.ReplacementFlag
			content = strings.Replace(content, itemFileCode.ReplacementFlag, itemFileCode.Content, 1)

			err = file.WriteFile(fullFIlePath, content)
			if err != nil {
				return nil, errors.New("写入文件失败：" + err.Error())
			}
			logInfo = append(logInfo, "替换文件："+fullFIlePath)
		}

		// 覆盖文件，直接写入
		if itemFileCode.WriteType == 0 {
			err = file.WriteFile(fullFIlePath, itemFileCode.Content)
			if err != nil {
				return nil, errors.New("写入文件失败：" + err.Error())
			}
			logInfo = append(logInfo, "生成文件："+fullFIlePath)
		}
	}

	// 5. 返回结果
	return logInfo, nil
}
