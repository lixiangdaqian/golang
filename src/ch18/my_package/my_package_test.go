package my_package

import (
	"testing"
	//包的导入路径是 src 下一级，到包所在的目录，下面就可以访问包里面的方法和数据了，小邪不可被访问
	"ch18/serial"
)

/*func init(){
	fmt.Print("init1")
}*/
func TestPackage(t *testing.T) {
	t.Log(serial.Fibnacci(10))
}

//TODO 注意： init方法，每一个依赖的package的init方法都会执行到，执行顺序取决于依赖的package对应的init方法，每个package可以有多个init函数
//TODO 注意：如果需要把自己的项目作为别人可以go get获取到的数据，注意别把src传上去，而应该把src下一级的内容传上去
