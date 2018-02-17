package requester

import (
	"github.com/tanan/moniple/domain"
	"net/http"
	"net/http/httputil"
)

type MontiorRequester struct {
}

func (r MontiorRequester) CheckEndpointStatus(url string) (domain.MonitorResult, error) {
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return domain.MonitorResult{}, err
	}
	dumpResp, _ := httputil.DumpResponse(resp, true)
	return domain.MonitorResult{
		Status: 200,
		TimeTaken: 0,
		Content: string(dumpResp),
	}, nil

}