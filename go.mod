module github.com/go-gost/core

go 1.21

toolchain go1.22.0

replace github.com/go-gost/x => ../x

require (
	github.com/go-gost/x v0.0.0-20240131151842-25dcf536c6f5
	github.com/xtls/reality v0.0.0-20231112171332-de1173cf2b19
	golang.org/x/sys v0.16.0
)

require (
	github.com/pires/go-proxyproto v0.7.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
)
