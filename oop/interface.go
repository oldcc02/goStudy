package main

import "fmt"

// AnimalIF interface本质是一个指针
type AnimalIF interface {
	aType() string
	aSleep()
}

type cat struct {
	catType string
}

func (cat cat) aSleep() {
	fmt.Println("cat sleep")
}

func (cat cat) aType() string {
	return cat.catType
}

type dog struct {
	dogType string
}

func (dog dog) aSleep() {
	fmt.Println("dog sleep")
}

func (dog dog) aType() string {
	return dog.dogType
}

func showType(animalIF AnimalIF) {
	fmt.Println(animalIF.aType())
}
func main() {
	var cat1 AnimalIF = &cat{"美短"}
	cat1.aSleep() // ==> cat sleep

	var dog1 AnimalIF = &dog{"土狗"}
	dog1.aSleep() // ==> dog sleep

	cat2 := cat{"银渐层"}
	showType(&cat2) // ==> 银渐层

	dog2 := dog{"沙皮狗"}
	showType(&dog2) // ==> 沙皮狗
}
