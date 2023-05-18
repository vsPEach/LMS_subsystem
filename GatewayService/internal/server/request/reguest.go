package request

import (
	"bytes"
	"fmt"
	"net/http"
)

func NewRequest(body []byte, url string) {
	r, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		fmt.Printf("Can't create http request: %s\n", err)
	}
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
	}

	fmt.Printf("%+v", res)
}
