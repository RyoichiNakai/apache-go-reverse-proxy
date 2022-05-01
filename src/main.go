package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

// constでおみくじのリストを作成できない
// 代わりに関数を使い初期化
func getOmikujiArray() []string {
	return []string{"大吉", "吉", "中吉", "小吉", "末吉", "凶", "大凶"}
}

func generateRandomNum(maxSize int) int {
	// おみくじの種類の大きさから乱数を作成
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxSize)
}

func handleOmikuji(ctx *gin.Context) {
	// おみくじの結果を返す
	omikujiResult := getOmikujiArray()
	randNum := generateRandomNum(len(omikujiResult))
	ctx.JSON(http.StatusOK, gin.H{
		"res": omikujiResult[randNum],
	})
}

func setupServer() *gin.Engine {
	engine := gin.Default()

	engine.GET("/", handleOmikuji)

	err := engine.Run(":3000")
	if err != nil {
		return nil
	}
	return engine
}

func main() {
	if err := setupServer().Run(); err != nil {
		return
	}
}
