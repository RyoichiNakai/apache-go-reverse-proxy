package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// constでArrayを作成できないので、
func TestGetOmikujiArray(t *testing.T) {
	expected := []string{"大吉", "吉", "中吉", "小吉", "末吉", "凶", "大凶"}
	actual := getOmikujiArray()
	for i := range actual {
		if expected[i] != actual[i] {
			t.Errorf("%d: expected %s, actual %s", i, expected[i], actual[i])
		}
	}
}

// generateRandomNumの単体テスト
func TestGenerateRandomNum(t *testing.T) {
	arraySize := len(getOmikujiArray())
	actual := generateRandomNum(arraySize)

	if 0 > actual || actual > arraySize {
		t.Errorf("generateRandomNum must be within %d", arraySize)
	}
}

// ginで立てたサーバのテスト
func TestHandleOmikuji(t *testing.T) {
	router := setupServer()
	w := httptest.NewRecorder()

	// テスト用のサーバーを立てる
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
