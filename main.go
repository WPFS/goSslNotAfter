package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/WPFS/goSslNotAfter/mail"
)

func main() {
	utcTime := time.Now().UTC().Unix()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	domains := []string{"https://www.baidu.com", "https://www.aliyun.com"}

	for _, domain := range domains {
		resp, err := client.Get(domain)
		defer resp.Body.Close()
		if err != nil {
			fmt.Errorf(domain, " 请求失败")
			panic(err)
		}
		remainingTime := (resp.TLS.PeerCertificates[0].NotAfter.Unix() - utcTime) / 86400
		if remainingTime > 10 {
			body := fmt.Sprintf("%s 的ssl证书10天后过期,请及时更新!\n", domain)

			mail.Sendmail(body)
		}
	}

}
