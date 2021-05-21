package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"os"

	"github.com/lucas-clemente/quic-go"
)

func run(network, local, addr, rawurl, filename string,
	tlsCfg *tls.Config, cfg *quic.Config,
	buffer []byte, dst *os.File,
	u *url.URL, rtmpType bool,
	quit chan os.Signal) {
	defer func() {
		close(quit)
	}()

	if addr == "" {
		ips, err := net.LookupIP(u.Host)
		if err != nil || len(ips) == 0 {
			fmt.Println(err)

			return
		}

		printDNS(u.Host, ips)

		switch u.Scheme {
		case "http":
			addr = ips[0].String() + ":80"
		case "https":
			addr = ips[0].String() + ":443"
		case "rtmp":
			addr = ips[0].String() + ":1935"
		default:
		}
	}

	switch u.Scheme {
	case "http":
		h1OverQUIC(network, local, addr, rawurl, tlsCfg, cfg, buffer, dst)
	case "https":
		h2OverQUIC(network, local, addr, rawurl, tlsCfg, cfg, buffer, dst)
	case "rtmp":
		rtmpOverQUIC(network, local, addr, rawurl, tlsCfg, cfg, filename, rtmpType)

	default:
		fmt.Println("unsupport")
	}
}
