package services

import (
	"crypto/tls"
	"time"
)

func GetSSLExpiryDays(urlStr string) int {
	host := urlStr[8:]
	if lastSlash := byte(47); host[len(host)-1] == lastSlash {
		host = host[:len(host)-1]
	}

	conn, err := tls.Dial("tcp", host+":443", nil)
	if err != nil {
		return 0
	}
	defer conn.Close()

	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	days := int(time.Until(expiry).Hours() / 24)
	return days
}