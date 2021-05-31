package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"

	quic "github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/h2quic"
)

func h2OverQUIC(network, local, addr, rawurl string,
	tlsCfg *tls.Config, cfg *quic.Config,
	buffer []byte, dst io.Writer) {
	roundTripper := &h2quic.RoundTripper{
		TLSClientConfig: tlsCfg,
		QuicConfig:      cfg,
		Dial:            dialFunc(local, addr),
	}

	defer roundTripper.Close()

	client := &http.Client{
		Transport: roundTripper,
	}

	dial := time.Now()
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

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

		return
	}

	readresp := time.Now()
	fmt.Println("recvresp:", readresp.Sub(dial).String())
	fmt.Println()

	rp, err := DumpResponse(resp)
	fmt.Println(string(rp))
	if err != nil {
		fmt.Println(err)
	}

	if _, err := io.CopyBuffer(dst, resp.Body, buffer); err != nil {
		fmt.Println(err)
	}

	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
}
