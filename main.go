package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
}

func main() {

	// make a new dog
	dog := NewDog("Ace")
	fmt.Println("Bruce Wayne goes to the pet store and comes home with a dog named", dog.Name)

	// begin simulating dog
	for {
		time.Sleep(time.Second) // slow simulation down to one action per second

		// let the dog die when its too old
		dog.Birthday()
		if dog.Age > 22 {
			break
		}

		// if 100 hungry, the dog has to eat
		if dog.Hunger >= 100 {
			fmt.Println(dog.Name, "is incredibly hungry and stops to eat.")
			dog.Eat()
			continue
		}

		// if 100 tired, the dog has to sleep
		if dog.Tiredness >= 100 {
			fmt.Println(dog.Name, "is incredibly tired and stops to sleep.")
			dog.Sleep()
			continue
		}

		// pick a random thing to do and do it
		choiceNumber := rand.Intn(7) // 7 being the number of things the dog can possibly do
		switch choiceNumber + 1 {
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
		case 6:
			dog.ChaseBird()
		case 7:
			dog.StopsCrime()
		}
	}

	fmt.Println(dog.Name, ", what a life, he came, he saw, he conquered, now he is dead and Bruce is sad.")
}

// Dog represents a four legged friend
type Dog struct {
	Name      string
	Hunger    int
	Tiredness int
	Age       int
}

// NewDog creates a new dog and returns a pointer to it
func NewDog(name string) *Dog {
	return &Dog{
		Name:      name,
		Hunger:    0,
		Tiredness: 0,
	}
}

// Birthday makes the dog get older
func (d *Dog) Birthday() {
	d.Age++
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

//ChaseBird makes the dog chase a bird
func (d *Dog) ChaseBird() {
	fmt.Println(d.Name, "chases a bird")
	d.Tiredness = d.Tiredness + 50

}

//StopsCrime makes Ace stop a crime
func (d *Dog) StopsCrime() {
	fmt.Println(d.Name, "stops a crime")
	d.Hunger = d.Hunger + 50
}
