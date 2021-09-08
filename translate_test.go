package baidutranslate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	appId := "YOUR KEY" // value from : https://fanyi-api.baidu.com/manage/developer
	appKey := ""        // YOUR SECRET

	if len(appKey) <= 0 {
		return
	}

	bt := NewBaiduTranslate(appId, appKey)
	{
		v, err := bt.Text("", "pt", "翻译成葡萄牙语")
		assert.Nil(t, err)
		assert.Equal(t, v, "Traduzido para Português")
	}
	{
		time.Sleep(1 * time.Second)
		v, err := bt.Text("en", "zh", "Hello World!")
		assert.Nil(t, err)
		assert.Equal(t, v, "你好，世界！")
	}
}
