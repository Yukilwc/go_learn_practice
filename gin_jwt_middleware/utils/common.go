package utils

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

func PostJson2Map(ctx *gin.Context) map[string]any {
	data, _ := io.ReadAll(ctx.Request.Body)
	var jsonMap map[string]any
	json.Unmarshal(data, &jsonMap)
	return jsonMap
}
