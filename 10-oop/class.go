package main

import "fmt"

type Animal struct {
	name string
	legs int
}

func (animal Animal) Eat() {
	fmt.Println(animal.name, "can eat")
}

func (animal Animal) leg() {
	fmt.Println("animal have leg")
}

type dog struct {
	Animal
	food string
}

func (dog dog) Eat2() {
	fmt.Println(dog.name, "eat", dog.food)
}

func (dog dog) leg() {
	dog.Animal.leg()
	fmt.Println(dog.name, "有", dog.legs)
}
func main() {
	dog1 := Animal{
		name: "泰迪",
		legs: 4,
	}
	dog1.Eat() // ==> 泰迪 can eat

	dog2 := dog{
		Animal: Animal{
			name: "柴犬",
			legs: 4,
		},
		food: "屎",
	}
	dog2.Eat()  //父类的方法 ==> 柴犬 can eat
	dog2.Eat2() //子类的方法 ==> 柴犬 eat 屎

	var dog3 dog
	dog3.name = "哈士奇"
	dog3.legs = 4
	dog3.food = "沙发"

	dog3.Eat2() // 子类的方法 ==> 哈士奇 eat 沙发
	dog3.leg()  // 继承并重写父类的方法
}
