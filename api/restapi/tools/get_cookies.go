package tools

import (
	"fmt"
	"net/http"
)

func GetCookies(req *http.Request, findCookie string) (map[string]string, error) {

	var cookies map[string]string
	for _, cookie := range req.Cookies() {
		cookies[cookie.Name] = cookie.Value
	}

	if cookies == nil {
		return cookies, fmt.Errorf("No cookies found")
	}
	if cookies[findCookie] == "" {
		return cookies, fmt.Errorf("%s cookie has no value", findCookie)
	}

	return cookies, nil
}
