package extention_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speck() {
	fmt.Print("...")
}

func (p *Pet) SpeckTo(somebody string) {
	p.Speck()
	fmt.Println(" ", somebody)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speck() {
	d.p.Speck()
	fmt.Print("wang")
}

func (d *Dog) SpeckTo(somebody string) {
	d.p.SpeckTo(somebody)
}

//匿名复用，重写无效，不能重载
type DogNew struct {
	Pet
}

func TestDog(t *testing.T) {
	dog := new(DogNew)
	dog.SpeckTo("haha")
}
