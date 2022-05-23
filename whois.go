package whois

import (
	"fmt"
	"strings"
)

type Server struct {
	Domain string
	Server []string
}

func (ws Server) String() string {
	return fmt.Sprintf("%s %s", ws.Domain, strings.Join(ws.Server, ""))
}

func (ws Server) Query(domain string) <-chan string {
	return AsyncQueryWithTimeout(domain, ws.Server)
}
