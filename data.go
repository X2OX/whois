package domain

import (
	"bufio"
	"bytes"
	"embed"
	"io/ioutil"
	"net"
	"time"
)

var (
	//go:embed .whois.list .non-icann.whois.list
	_nonIcannWhoisList embed.FS

	whoisServerList = make([]WhoisServer, 0, 1100)
)

func (ws WhoisServer) Query(domain string) <-chan string {
	return AsyncQueryWithTimeout(domain, ws.Server)
}

func queryWhoisServerTimeout(domain, server string, timeout time.Duration) (string, error) {
	connection, err := net.DialTimeout("tcp", net.JoinHostPort(server, "43"), timeout)
	if err != nil {
		return "", err
	}
	defer connection.Close()

	if _, err = connection.Write([]byte(domain + "\r\n")); err != nil {
		return "", err
	}

	var buf []byte
	if buf, err = ioutil.ReadAll(connection); err != nil {
		return "", err
	}
	return string(buf), nil
}

func init() {
	file, err := _nonIcannWhoisList.Open(".whois.list")
	if err != nil {
		panic(err)
	}
	br := bufio.NewReader(file)

	for {
		var line []byte
		if line, _, err = br.ReadLine(); err != nil {
			break
		}
		arr := bytes.Split(line, []byte(" "))
		if len(arr) < 2 {
			continue
		}
		server := make([]string, len(arr)-1)
		for i := 1; i < len(arr); i++ {
			if len(arr[i]) > 3 {
				server = append(server, string(arr[i]))
			}
		}

		whoisServerList = append(whoisServerList, WhoisServer{
			Domain: string(arr[0]),
			Server: server,
		})
	}
}
