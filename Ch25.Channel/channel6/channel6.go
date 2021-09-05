package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Number int
	Body   string
	Tire   string
	Color  string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Println("Start Factory")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	var carNum int = 1
	for {
		select {
		case <-tick:
			//Make body
			fmt.Println("Make Body", carNum)
			car := &Car{}
			car.Number = carNum
			car.Body = "Sports car"
			tireCh <- car

		case <-after:
			fmt.Println("MakeBody Close")
			close(tireCh)
			wg.Done()
			return

		}

		carNum++
	}
}

func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		//Make a body
		fmt.Printf("Install Tire %d\n", car.Number)
		time.Sleep(time.Second)
		car.Tire = "Winter Tire"
		paintCh <- car
	}

	fmt.Println("Tire Done")
	wg.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		//Make a body
		time.Sleep(time.Second)
		fmt.Printf("Paint Car %d\n", car.Number)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car : %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}

	fmt.Println("Paint Done")
	wg.Done()
}
