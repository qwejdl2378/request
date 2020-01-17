package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Options struct {
	Url string
	Method string
	Headers map[string]string
	Cookies []*http.Cookie
	Body []byte
}

func Request(options Options)([]byte, *http.Response, error) {
	if options.Method == "" {
		options.Method = "GET"
	}
	var client http.Client
	httpReq, httpErr := http.NewRequest(options.Method, options.Url, bytes.NewBuffer(options.Body))
	if httpErr != nil {
		return nil, nil, httpErr
	}
	headers := http.Header{}
	for k, v := range options.Headers {
		if v != "" {
			headers.Set(k, v)
		}
		continue
	}
	httpReq.Header = headers
	if len(options.Cookies) > 0 {
		for _, cookie := range options.Cookies {
			httpReq.AddCookie(cookie)
		}
	}
	response, requestErr := client.Do(httpReq)
	if requestErr != nil {
		return nil, nil, requestErr
	}
	defer response.Body.Close()
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return nil, response, readErr
	}
	return body, response, nil
}