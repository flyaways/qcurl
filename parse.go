package main

import (
	"crypto/tls"
	"strings"

	"github.com/lucas-clemente/quic-go"
)

func parseCfg(version, serverName string) (*tls.Config, *quic.Config) {
	var gquicvm = map[string]quic.VersionNumber{
		"39": quic.VersionGQUIC39,
		"43": quic.VersionGQUIC43,
		"44": quic.VersionGQUIC44,
	}

	versions := []quic.VersionNumber{}
	if version != "" {
		vs := strings.Split(version, ",")
		for _, v := range vs {
			if vv, ok := gquicvm[v]; ok {
				versions = append(versions, vv)
			}
		}
	}

	cfg := &quic.Config{Versions: versions}

	tlscfg := &tls.Config{
		ServerName:             serverName,
		InsecureSkipVerify:     true,
		SessionTicketsDisabled: true,
		// NextProtos:             []string{"39", "43", "44"},
	}

	return tlscfg, cfg
}
