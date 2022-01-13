//go:build !amd64 && !arm64 && !ppc64
// +build !amd64,!arm64,!ppc64

package fasthttp

const (
	maxHexIntChars = 7
)
