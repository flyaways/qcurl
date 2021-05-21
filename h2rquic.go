package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
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
		Dial:            dialFunc(local),
	}

	defer roundTripper.Close()

	client := &http.Client{
		Transport: roundTripper,
	}

	dial := time.Now()
	resp, err := client.Get(rawurl)
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
