package domain

import (
	"fmt"
	"strings"
)

type WhoisServer struct {
	Domain string
	Server []string
}

func (ws WhoisServer) String() string {
	return fmt.Sprintf("%s %s", ws.Domain, strings.Join(ws.Server, ""))
}
