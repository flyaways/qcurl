# qcurl

## Usage

```sh
Usage of ./qcurl:
  -addr string
        example: 1.2.3.4:80
  -bind string
        bind local ip
  -buffer int
        buffer size in byte (default 102400)
  -file string
        specify i/o flv file (default "d.flv")
  -network string
        network (default "udp4")
  -quic-version string
        support 39,43,44 (default "43")
  -sni string
        domain
  -t    pull=true,publish=false (default true)
```

```sh
./qcurl https://domain/app/name.flv

./qcurl http://domain/app/name.flv

./qcurl rtmp://domain/app/name
```