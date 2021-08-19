package Proxy

import (
	"fmt"
	"os"
)

type Image interface {
	Display()
}

type RealImage struct {
	file string
}

func newRealImage(file string) *RealImage {
	r := &RealImage{file}
	r.LoadFromDisk()
	return r
}

func (r *RealImage) LoadFromDisk() {
	if _, err := os.Stat(r.file); err != nil && !os.IsNotExist(err) {
		panic("unsupported real image load err")
	}
}

func (r RealImage) Display() {
	fmt.Println("real image show")
}

type ProxyImage struct {
	image *RealImage
	file  string
}

func newProxyImage(file string) *ProxyImage {
	return &ProxyImage{file: file}
}

func (p *ProxyImage) Display() {
	if p.image == nil {
		p.image = newRealImage(p.file)
	}
	p.image.Display()
}
