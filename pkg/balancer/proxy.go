package balancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/balchua/consistent-demo/pkg/logging"
)

type proxyRequest struct {
	previousNode string
	url          string
}

func proxy(request proxyRequest) {

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodPost, request.url, nil)
	if err != nil {
		logging.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("previousNode", request.previousNode)
	response, err := client.Do(req)
	if err != nil {
		logging.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	b, err := httputil.DumpResponse(response, true)
	if err != nil {
		log.Fatalln(err)
	}
	logging.Infof("response status is %s", response.Status)
	logging.Infof("response body is %s", b)
}
