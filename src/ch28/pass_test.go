package ch28

import (
	"testing"
)

type Person struct {
	Age int
	Name string
}
func TestPass(t *testing.T)  {
	var person = Person{
		Age : 10,
		Name: "xiaoming",
	}
	t.Logf("数据是：%v，内存地址是： %p", person, &person)
	passValue(person, t)
	t.Logf("值传递后的数据是：%v，内存地址是： %p", person, &person)//值传递过去之后，传递的只是副本，副本改了以后并不影响原来的数据
	var person1 = &Person{
		Age : 10,
		Name: "xiaoming",
	}
	t.Logf("指针传递前的数据是：%v，内存地址是： %p, 指针的内存地址是： %p", person1, &person1, person1)//值传递过去之后，传递的只是副本，副本改了以后并不影响原来的数据
	passPoint(person1, t)
	t.Logf("指针传递后的数据是：%v，内存地址是： %p, 指针的内存地址是： %p", person1, &person1, person1)//值传递过去之后，传递的只是副本，副本改了以后并不影响原来的数据
}

//传递值
func passValue(person Person, t *testing.T)  {
	person.Name += "hello"
	person.Age += 10
	t.Logf("值传递过程中的数据是：%v，内存地址是： %p", person, &person)//值传递过去之后，传递的只是副本，副本改了以后并不影响原来的数据
}

//传递指针
func passPoint(person *Person, t *testing.T)  {
	person.Name += "hello"
	person.Age += 10
	t.Logf("指针传递过程中的数据是：%v，内存地址是： %p， 指针的内存地址是：%p", person, &person, person)//值传递过去之后，传递的只是副本，副本改了以后并不影响原来的数据
}

//测试变量赋值
func TestAssin(t *testing.T){
	person := Person{
		Age:10,
		Name:"hello",
	}
	person1 := person
	t.Logf("数据是：%v，内存地址是： %p", person, &person)
	t.Logf("数据是：%v，内存地址是： %p", person1, &person1)
}
