package main

import (
	"fmt"
	"time"
)

func doWork(working, customer chan int) {
	fmt.Println("Barber working...")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Barber end")

	customer <- 1
	working <- 0
}

func barbershop(wait_room chan int, customer chan chan int) {
	busy := false
	working := make(chan int)
	for {
		select {
		case x := <-customer:
			if busy {
				wait_room <- 0
			} else {
				busy = true
				wait_room <- 1
				go doWork(working, x)
			}
		case <-working:
			busy = false
			wait_room <- 2
		}
	}
}

func waiting_room(customer chan chan int) {
	var free_seat bool = true
	var customer_seat chan int

	wait_room := make(chan int)

	barbershopChan := make(chan chan int)

	go barbershop(wait_room, barbershopChan)

	for {
		select {
		case x := <-customer:
			if free_seat {
				barbershopChan <- x
				busy := <-wait_room
				if busy == 0 {
					fmt.Println("Seat is free, but barber is busy")
					free_seat = false
					customer_seat = x
				} else {
					fmt.Println("Seat is free and customer goes in")
				}

			} else {
				fmt.Println("Seat isnt free")
				// fmt.Println("Seat isnt free")
				x <- 0
			}
		case x := <-wait_room:
			if x == 2 {
				if !free_seat {
					barbershopChan <- customer_seat
					<-wait_room
					free_seat = true
				}
			}
		}
	}
}

func newCustomer(waiting_room chan chan int, number int, quit chan int) {
	customer := make(chan int)
	waiting_room <- customer
	for {
		select {
		case i := <-customer:
			if i == 0 {
				time.Sleep(1500 * time.Millisecond)
				fmt.Printf("Customer %d tries again\n", number)
				waiting_room <- customer
			} else {
				fmt.Printf("Customer %d quit\n", number)
				quit <- number
				return
			}
		}
	}
}

func main() {
	customer := make(chan chan int)
	quit := make(chan int)
	go waiting_room(customer)
	n := 5
	for i := 0; i < n; i++ {
		fmt.Printf("Customer %d go in\n", i)
		go newCustomer(customer, i, quit)
		time.Sleep(2000 * time.Millisecond)
	}

	for i := 0; i < n; {
		<-quit
		i++
	}
}
