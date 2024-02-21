package dialer

import (
	"crypto/tls"
	"net/url"

	"github.com/go-gost/core/common/net/dialer"
	"github.com/go-gost/core/logger"
	"github.com/xtls/reality"
)

type Options struct {
	Auth                *url.Userinfo
	TLSConfig           *tls.Config
	REALITYClientConfig *RealityClientConfig
	Logger              logger.Logger
	ProxyProtocol       int
}

type RealityClientConfig struct {
	*reality.Config
	//TODO: Fingerprint string
	Show       bool
	ServerName string
	PublicKey  []byte
	ShortId    []byte
}

type Option func(opts *Options)

func AuthOption(auth *url.Userinfo) Option {
	return func(opts *Options) {
		opts.Auth = auth
	}
}

func TLSConfigOption(tlsConfig *tls.Config) Option {
	return func(opts *Options) {
		opts.TLSConfig = tlsConfig
	}
}

func REALITYClientConfigOption(config *RealityClientConfig) Option {
	return func(opts *Options) {
		opts.REALITYClientConfig = config
	}
}

func LoggerOption(logger logger.Logger) Option {
	return func(opts *Options) {
		opts.Logger = logger
	}
}

func ProxyProtocolOption(ppv int) Option {
	return func(opts *Options) {
		opts.ProxyProtocol = ppv
	}
}

type DialOptions struct {
	Host      string
	NetDialer *dialer.NetDialer
}

type DialOption func(opts *DialOptions)

func HostDialOption(host string) DialOption {
	return func(opts *DialOptions) {
		opts.Host = host
	}
}

func NetDialerDialOption(netd *dialer.NetDialer) DialOption {
	return func(opts *DialOptions) {
		opts.NetDialer = netd
	}
}

type HandshakeOptions struct {
	Addr string
}

type HandshakeOption func(opts *HandshakeOptions)

func AddrHandshakeOption(addr string) HandshakeOption {
	return func(opts *HandshakeOptions) {
		opts.Addr = addr
	}
}
