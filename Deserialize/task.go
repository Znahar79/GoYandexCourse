package main

import (
	"encoding/json"
	"fmt"
)

const rawResp = `
{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
}
`

type (
	Response struct {
		Header struct {
			Code    int    `json: "code"`
			Message string `json: "message"`
		} `json: "header"`

		Data []DataItem `json: "data,omitempty"`
	}

	DataItem struct {
		Type       string `json: "type"`
		Id         int    `json: "id"`
		Attributes struct {
			Email      string `json: "email"`
			ArticleIds []int  `json:"article_ids"`
		} `json: "attributes"`
	}
)

func main() {
	fmt.Println(ReadResponse(rawResp))
}

func ReadResponse(rawResp string) (Response, error) {
	var response Response
	err := json.Unmarshal([]byte(rawResp), &response)
	return response, err
}
