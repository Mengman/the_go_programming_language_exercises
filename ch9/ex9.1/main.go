package main

import (
	"github.com/Mengman/the_go_programming_language_exercises/ch9/ex9.1/bank"
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		bank.Deposit(100)
		log.Printf("deposit routine: balance %d\n", bank.Balance())
	}()


	go func() {
		wg.Add(1)
		defer wg.Done()
		if bank.Withdraw(100) {
			log.Printf("withdraw routine: withdraw success balance: %d\n", bank.Balance())
		} else {
			log.Printf("withdraw routine: withdraw fail balance: %d\n", bank.Balance())
		}
	}()

	wg.Wait()
	log.Printf("main routine: balance: %d\n", bank.Balance())
}
