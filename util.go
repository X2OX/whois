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
			atomic.AddInt32(&count, 1)
			serverAddr := server[i]

			go func() {
				if resp, _ := QueryWithTimeout(data, serverAddr, 10*time.Second); resp != "" {
					if atomic.CompareAndSwapInt32(&isReturn, 0, 1) {
						ch <- resp
					}
				}

				if atomic.AddInt32(&count, -1); atomic.LoadInt32(&count) == 0 &&
					atomic.LoadInt32(&isReturn) == 0 {
					ch <- "error"
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
