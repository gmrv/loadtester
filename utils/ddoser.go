package utils

import "net/http"

type ddoser struct {}

func (d *ddoser) performRequest(url string) (status int) {
	status = 0
	resp, err := http.Get(url)
	if err == nil {
		status = resp.StatusCode
		WriteLog(status, " ", url)
		resp.Body.Close()
	} else {
		status = -1
		WriteLog(err)
	}
	return status
}

func (d *ddoser) PerformMultiRequest(url string, count int) {
	if count == 0 {
		for {
			if !GetHelper().FL_KEEP_GOING {
				break
			}
			GetHelper().C <- d.performRequest(url)
		}
	} else {
		for i := 0; i < count; i++ {
			if !GetHelper().FL_KEEP_GOING {
				break
			}
			GetHelper().C <- d.performRequest(url)
		}
	}
}