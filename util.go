package whois

import (
	"io/ioutil"
	"net"
	"sync/atomic"
	"time"
)

func AsyncQueryWithTimeout(data string, server []string) <-chan string {
	var (
		ch       = make(chan string)
		count    = int32(len(server))
		isReturn = int32(0)
	)

	go func() {
		for i := range server {
			serverAddr := server[i]

			go func() {
				resp, err := QueryWithTimeout(data, serverAddr, 10*time.Second)
				if err == nil && resp != "" && atomic.CompareAndSwapInt32(&isReturn, 0, 1) {
					ch <- resp
				}

				atomic.AddInt32(&count, -1)

				if atomic.LoadInt32(&count) == 0 && atomic.LoadInt32(&isReturn) == 0 {
					if err != nil {
						ch <- err.Error()
					} else {
						ch <- "error"
					}
				}
			}()
		}
	}()

	return ch
}

func QueryWithTimeout(data, server string, timeout time.Duration) (string, error) {
	connection, err := net.DialTimeout("tcp", net.JoinHostPort(server, "43"), timeout)
	if err != nil {
		return "", err
	}
	defer connection.Close()

	if _, err = connection.Write([]byte(data + "\r\n")); err != nil {
		return "", err
	}

	var buf []byte
	if buf, err = ioutil.ReadAll(connection); err != nil {
		return "", err
	}
	return string(buf), nil
}
