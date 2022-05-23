package whois

import (
	"bufio"
	"bytes"
	"embed"
)

var (
	//go:embed .whois.list .non-icann.whois.list
	_whoisList      embed.FS
	whoisServerList = make([]Server, 0, 1100)
)

func init() {
	parseData(".whois.list", ".non-icann.whois.list")
}

func parseData(filenames ...string) {
	for _, filename := range filenames {
		file, err := _whoisList.Open(filename)
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

			whoisServerList = append(whoisServerList, Server{
				Domain: string(arr[0]),
				Server: server,
			})
		}
		_ = file.Close()
	}
}
