package writeCode

// 生成代码入参
type GenerateCodeReq struct {
	Id      uint64 `json:"id" binding:"required"`      // 业务模型ID
	Type    int8   `json:"type" `                      // 0-后端 1-前端
	SysCode string `json:"sysCode" binding:"required"` // 系统编码
	Domain  string `json:"domain" binding:"required"`  // 领域名称
}

// 公共返回
type CommonResp[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type GenerateCodeResp struct {
	Codes       []*GenerateCodeRespItemData
	SourceDatas *GenerateCodeSourceData
}

// 生成代码源数据 | 新
type GenerateCodeSourceData struct {
	VueScrtipTags         string                 // 语法有影响 <script setup>
	VueScrtipTsTags       string                 // 语法有影响 <script lang="ts" setup>
	Domain                string                 // 领域域名(小驼峰)
	DomainUpperName       string                 // 领域域名(大驼峰)
	DomainZhName          string                 // 域名中文名
	TableName             string                 // 数据表名(蛇形)
	ModelName             string                 // 模型名称(小驼峰)
	ModelUpperName        string                 // 模型名称(大驼峰)
	ModelZhName           string                 // 模型名称(中文)
	BuildingBlock         *BuildingBlock         // 积木系统模板源数据
	BuildingBlockFrontend *BuildingBlockFrontend // 积木系统前端模板源数据
}

// 积木系统模板源数据
type BuildingBlock struct {
	EntityContent  string `json:"entityContent"`  // 业务实体内容
	RequestContent string `json:"requestContent"` // 请求内容
	SearchContent  string `json:"searchContent"`  // 搜索内容
}

// 积木系统前端模板源数据
type BuildingBlockFrontend struct {
	QueryFormContent   string `json:"queryFormContent"` // 搜索表单内容
	QueryDataContent   string `json:"queryDataContent"` // 搜索数据内容
	ColumnsDataContent string `json:"listDataContent"`  // 列表数据内容
	FormContent        string `json:"formContent"`      // 表单内容
	FormDataContent    string `json:"formDataContent"`
}

// 生成代码响应单条数据
type GenerateCodeRespItemData struct {
	FileName string `json:"fileName" common:"文件名称"`
	FilePath string `json:"filePath" common:"文件路径"`
	Content  string `json:"content" common:"文件内容"`
}
