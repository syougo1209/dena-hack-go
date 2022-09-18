package service

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/syougo1209/dena-hack-go/model"
)

func GroupingService(groupingRequest *model.GroupingRequest) (*model.UsersGroup, error) {
	/**
		PythonサーバーへHTTPリクエスト(JSON)を送信し、
		HTTPレスポンス(JSON)を受けとる
	**/

	// 送信するJSONの用意
	jsonString, err := json.Marshal(groupingRequest)
	if err != nil {
		return nil, err
	}

	// HTTPリクエストの作成
	endpoint := "http://18.183.252.13:8000/matchig/"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスを構造体に入れる
	usersGroup := new(model.UsersGroup)
	if err := json.NewDecoder(resp.Body).Decode(&usersGroup); err != nil {
		return nil, err
	}
	return usersGroup, nil
}
