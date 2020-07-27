# golang
golang的一些服务


//TODO 注意： init方法，每一个依赖的package的init方法都会执行到，执行顺序取决于依赖的package对应的init方法，每个package可以有多个init函数
//TODO 注意：如果需要把自己的项目作为别人可以go get获取到的数据，注意别把src传上去，而应该把src下一级的内容传上去
