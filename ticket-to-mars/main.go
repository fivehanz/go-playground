/*
* Exercise from the book "Get Programming with Go"
* Unit 1: Lesson 5
 */

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const distanceToMars int = 62100000      // km
const secondsInOneDay int = 24 * 60 * 60 // 24hrs = 60mins * 24 = 60s * 60 * 24

var SpaceLines = [...]string{"Space Adventures", "SpaceX", "Virgin Galactic"}
var speedIntervals = [2]int{16, 30} // speed range is in km/s
var priceIntervals = [2]int{36, 50} // price range is in million $
var typesOfTicket = map[string]int{
	"One-way":    1,
	"Round-trip": 2,
}

func createTicket(spaceline string, speed int, price int, ticketType string) string {
	price = typesOfTicket[ticketType] * price // calculate price according to the trip
	days := calculateDays(speed)
	return fmt.Sprintf("%-20v %-4v %-10v $%4v", spaceline, days, ticketType, price)
}

func calculateDays(speed int) int {
	return distanceToMars / speed / secondsInOneDay
}

func getRandomTicketType() string {
	randNum := rand.Intn(2)
	if (randNum % 2) == 0 {
		return "One-way"
	} else {
		return "Round-trip"
	}
}

func getTicketPrice(speed int) int {
	return speed + priceIntervals[1] - speedIntervals[1]
}

func getRandomSpeed() int {
	return rand.Intn(1+speedIntervals[1]-speedIntervals[0]) + speedIntervals[0]
}

func getRandomSpaceLine() string {
	randNum := rand.Intn(3)
	return SpaceLines[randNum]
}

func generateTickets(num int) {
	for i := 0; i < num; i++ {
		spaceline := getRandomSpaceLine()
		speed := getRandomSpeed()
		price := getTicketPrice(speed)
		ticketType := getRandomTicketType()

		fmt.Println(createTicket(spaceline, speed, price, ticketType))
	}
}

func main() {
	N, err := strconv.Atoi(os.Args[1])
	if err == nil {
		fmt.Printf("%-20v %-4v %-10v %4v\n", "Spaceline", "Days", "Trip type", "Price")
		fmt.Println("============================================")
		generateTickets(N)
	} else {
		fmt.Println(err)
	}
}
