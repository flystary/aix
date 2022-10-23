package request

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"crypto/md5"
	"encoding/hex"

	n "/aix/net"
)

var (
	// token单例
	onces = &sync.Once{}
	// TOKEN token
	// TOKEN string
)

func newMD5(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}

// GetToken 获取token
func GetToken(s *n.Service) *n.Service {
	var result = make(map[string]interface{})
	requestData := make(url.Values)
	requestData["username"] = []string{"matrix"}
	requestData["password"] = []string{newMD5(newMD5("4A9sOpYL"))}
	requestData["client_id"] = []string{"browser"}
	requestData["client_secret"] = []string{"b7n3i7kzg22y3p035rw3rd9sfzvs4cv0"}
	requestData["grant_type"] = []string{"password"}

	onces.Do(func() {
		res, err := http.PostForm(s.Router.GetTokenFromRoute(), requestData)
		if err != nil {
			fmt.Printf("Login Error: %v", err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("ReadAll IO Error: %v", err)
		}
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Printf("Unmarshal body Error: %v", err)
		}
		s.Token = result["access_token"].(string)
	})
	return s
}

// func GetMode(s *n.Service) *n.Service {
// 	return s
// }