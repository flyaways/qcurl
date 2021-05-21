package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/lucas-clemente/quic-go"
)

func h1OverQUIC(network, local, addr, rawurl string,
	tlsCfg *tls.Config, cfg *quic.Config,
	buffer []byte, dst io.Writer) {
	d := time.Now()

	var session quic.Session

	stream, err := dial(local, tlsCfg, cfg, session)(network, addr)
	if err != nil {
		fmt.Println(err)

		return
	}

	defer func() {
		if session != nil {
			session.Close()
		}
	}()
	defer stream.Close()

	handshake := time.Now()

	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		fmt.Println(err)

		return
	}

	data, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println()
	fmt.Println(string(data))

	if _, err := stream.Write(data); err != nil {
		fmt.Println(err)

		return
	}

	reader := bufio.NewReader(stream)
	resp, err := http.ReadResponse(reader, req)
	if err != nil {
		fmt.Println(err)

		return
	}

	readresp := time.Now()
	fmt.Println("handshake:", handshake.Sub(d).String())
	fmt.Println("recvresp:", readresp.Sub(handshake).String())
	fmt.Println()

	rp, err := DumpResponse(resp)
	fmt.Println(string(rp))
	if err != nil {
		fmt.Println(err)
	}

	if resp != nil && resp.Body != nil {
		if _, err := io.CopyBuffer(dst, resp.Body, buffer); err != nil {
			fmt.Println(err)

			return
		}
	}

	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}

	return
}
