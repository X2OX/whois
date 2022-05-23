package domain

import (
	"fmt"
	"testing"
)

func TestDomain(t *testing.T) {
	fmt.Println(Parse("asd.sds.com"))
	fmt.Println(Parse("asd.eu.org"))
	fmt.Println(Parse("aa.cn.eu.org"))
	fmt.Println(Parse("asc.ac.book"))
	fmt.Println(Parse("asdas.asdas.我爱你"))
	fmt.Println(Parse("asdas.asdas.xn--6qq986b3xl"))
	fmt.Println(Parse("xn--6qq986b3xl.xn--6qq986b3xl.xn--6qq986b3xl"))
}
