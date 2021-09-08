package baidutranslate

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const baiduApiEndpoint = "https://fanyi-api.baidu.com/api/trans/vip/translate"

type BaiduTranslate struct {
	AppId  string
	AppKey string
}

type TransResult struct {
	Source string `json:"src"`
	Dest   string `json:"dst"`
}

type Result struct {
	ErrorCode *string       `json:"error_code"`
	ErrorMsg  *string       `json:"error_msg"`
	From      string        `json:"form"`
	To        string        `json:"to"`
	Items     []TransResult `json:"trans_result"`
}

var inst *BaiduTranslate

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Default() *BaiduTranslate {
	return inst
}

func NewBaiduTranslate(appId, appKey string) *BaiduTranslate {
	t := BaiduTranslate{
		AppId:  appId,
		AppKey: appKey,
	}
	inst = &t
	return &t
}

func (t *BaiduTranslate) makeSign(query, salt string) string {
	hashVal := md5.Sum([]byte(t.AppId + query + salt + t.AppKey))
	return fmt.Sprintf("%x", hashVal)
}

//
//Source document : http://api.fanyi.baidu.com/product/113
//Text translate
func (t *BaiduTranslate) Text(from, to, content string) (string, error) {
	if len(content) <= 0 {
		return "", nil
	}

	if len(from) <= 0 {
		from = "auto"
	}

	if len(to) <= 0 {
		return "", errors.New("invalid `to` params")
	}

	salt := strconv.FormatInt(rand.Int63(), 10)
	sign := t.makeSign(content, salt)

	data := url.Values{
		"appid": []string{t.AppId},
		"q":     []string{content},
		"from":  []string{from},
		"to":    []string{to},
		"salt":  []string{salt},
		"sign":  []string{sign},
	}

	client := &http.Client{}
	r, err := http.NewRequest(http.MethodPost, baiduApiEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tr Result
	err = json.Unmarshal(body, &tr)
	if err != nil {
		return "", err
	}

	if tr.ErrorCode != nil {
		msg := *tr.ErrorCode
		if tr.ErrorMsg != nil {
			msg = *tr.ErrorMsg
		}
		return "", errors.New(msg)
	}

	if len(tr.Items) <= 0 {
		return "", nil
	}
	return tr.Items[0].Dest, nil
}
