package my_package_test

import (
	"testing"
	//包的导入路径是 src 下一级，到包所在的目录，下面就可以访问包里面的方法和数据了，小邪不可被访问
	"ch18/serial"
)

func TestPackage(t *testing.T)  {
	t.Log(serial.Fibnacci(10))
}