package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init(){
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
}

func main() {

	// make a new dog
	dog := NewDog("Fred")

	// begin simulating dog
	for {

		// let the dog die when its too old
		dog.Age++
		if dog.Age > 20 {
			break
		}

		// if 100 hungry, the dog has to eat
		if dog.Hunger >= 100 {
			fmt.Println(dog.Name,"is incredibly hungry and stops to eat.")
			dog.Eat()
			continue
		}

		// if 100 tired, the dog has to sleep
		if dog.Tiredness >= 100 {
			fmt.Println(dog.Name,"is incredibly tired and stops to sleep.")
			dog.Sleep()
			continue
		}

		// pick a random thing to do and do it
		choiceNumber := rand.Intn(5) // 5 being the number of things the dog can possibly do
		switch choiceNumber {
		case 1:
			dog.Bark()
		case 2:
			dog.Run()
		case 3:
			dog.Eat()
		case 4:
			dog.Sleep()
		case 5:
			dog.Play()
		}
		time.Sleep(time.Second) // slow simulation down to one action per second
	}

	fmt.Println(dog.Name,"has lived a good life but is now dead.")
}

type Dog struct {
	Name string
	Hunger int
	Tiredness int
	Age int
}

// NewDog creates a new dog and returns a pointer to it
func NewDog(name string) *Dog {
	return &Dog {
		Name: name,
		Hunger: 0,
		Tiredness: 0,
	}
}

// Sleep makes the dog take a rest. resets tiredness to 0
func (d *Dog) Sleep() {
	fmt.Println(d.Name, "sleeps...")
	d.Tiredness = 0
}

// Bark makes the dog bark
func (d *Dog) Bark() {
	fmt.Println(d.Name, "goes BARK BARK BARK")
	d.Tiredness = d.Tiredness + 20
}

// Eat makes the dog eat. returns hunger to 0
func (d *Dog) Eat() {
	fmt.Println(d.Name, "eats some food")
	d.Hunger = 0
}

// Play makes the dog play
func (d *Dog) Play() {
	fmt.Println(d.Name, "plays with toys")
	d.Tiredness = d.Tiredness + 30
}

// Run makes the dog run around for awhile
func (d *Dog) Run() {
	fmt.Println(d.Name, "runs around")
	d.Tiredness = d.Tiredness + 40
}
