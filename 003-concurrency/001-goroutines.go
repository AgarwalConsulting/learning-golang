package main

import (
	"github.com/AgarwalConsulting/learning-golang/003-concurrency/functions"
	"time"
)

func main() {
	go functions.WomenRelatedCrimeStats(2012)
	go functions.WomenRelatedCrimeStats(2013)
	go functions.WomenRelatedCrimeStats(2014)
	go functions.WomenRelatedCrimeStats(2015)

	time.Sleep(5 * time.Second)
}
