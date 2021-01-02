package lib

import (
	"io/ioutil"
	"net/http"
)

func GetBodyBinaryFromHttpResp(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
