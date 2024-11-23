package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"risk/utils"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

var esClient *elasticsearch.Client

// 向 Elasticsearch 写入数据
func WriteToElasticsearch(c *gin.Context) {
	var document map[string]interface{}
	if err := c.ShouldBindJSON(&document); err != nil {
		utils.ErrorResponse(c, 400, "参数错误: "+err.Error())
		return
	}

	// 将 document 转换为 JSON 字符串
	jsonData, err := json.Marshal(document)
	if err != nil {
		utils.ErrorResponse(c, 500, "JSON 转换失败: "+err.Error())
		return
	}

	res, err := esClient.Index(
		"risk_strategy",
		bytes.NewReader(jsonData),
	)

	if err != nil {
		utils.ErrorResponse(c, 500, "写入 Elasticsearch 失败: "+err.Error())
		return
	}

	utils.SuccessResponse(c, res)
}

// 从 Elasticsearch 读取数据
func ReadFromElasticsearch(c *gin.Context) {
	// 创建搜索请求
	var query map[string]interface{}
	if err := c.ShouldBindJSON(&query); err != nil {
		utils.ErrorResponse(c, 400, "参数错误: "+err.Error())
		return
	}

	// 将查询转换为 JSON
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		utils.ErrorResponse(c, 500, "查询转换失败: "+err.Error())
		return
	}

	// 执行搜索
	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex("risk_strategy"),
		esClient.Search.WithBody(&buf),
	)

	if err != nil {
		utils.ErrorResponse(c, 500, "搜索失败: "+err.Error())
		return
	}
	defer res.Body.Close()

	// 解析响应
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		utils.ErrorResponse(c, 500, "解析响应失败: "+err.Error())
		return
	}

	utils.SuccessResponse(c, result)
}

// Elasticsearch 连接配置
func ElasticsearchConfig() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://127.0.0.1:9200",
			// Username: "elastic",
			// Password: "123456",
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	esClient = client
}
