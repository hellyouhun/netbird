//go:build android

package android

import "fmt"

type PeerRoutes struct {
	routes []string
}

func (p *PeerRoutes) Get(i int) (string, error) {
	if i >= len(p.routes) || i < 0 {
		return "", fmt.Errorf("%d 超出范围", i)
	}
	return p.routes[i], nil
}

func (p *PeerRoutes) Size() int {
	return len(p.routes)
}
