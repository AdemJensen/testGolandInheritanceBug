package main

import (
	"testGolandInheritanceBug/impl1"
	"testGolandInheritanceBug/impl2"
)

type implementation2 = impl2.Impl2

type impl struct {
	impl1.Impl1
	implementation2
}

func main() {
	var ptr Inf
	ptr = new(impl)
	println(ptr)
}
