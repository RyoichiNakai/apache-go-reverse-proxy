package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"math/rand"
)

// constでおみくじのリストを作成できない
// 代わりに関数を使い初期化
func getOmikuArray() []string {
	return []string{"大吉", "吉", "中吉", "小吉", "末吉", "凶", "大凶"}
}

func generateRandomNum(maxSize int) int {
	// おみくじの種類の大きさから乱数を作成
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxSize)
}

func handleOmikuji(ctx *gin.Context) {
	// おみくじの結果を返す
	omikujiResult := getOmikuArray() 

	ctx.Writer.WriteString(omikujiResult[generateRandomNum(len(omikujiResult))])
}

func main() {
	engine := gin.Default()
	
	engine.GET("/", handleOmikuji)

	engine.Run(":3000")
}