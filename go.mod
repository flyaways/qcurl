module github.com/flyaways/qcurl

go 1.15

require (
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/lucas-clemente/quic-go v0.10.2
	github.com/onsi/ginkgo v1.14.0 // indirect
	github.com/q191201771/lal v0.22.0
	github.com/q191201771/naza v0.18.5
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace (
	github.com/lucas-clemente/quic-go => github.com/flyaways/quic-go v0.10.9
	github.com/q191201771/lal => github.com/flyaways/lal v0.22.1
)
