package whois

import (
	"testing"
)

func TestWhoisServer_Query(t *testing.T) {
	ws := &Server{
		Domain: "gov.cn",
		Server: []string{"whois.cnnic.cn"},
	}
	s := ws.Query("scio.gov.cn")
	ss := <-s
	t.Log(ss)
}
