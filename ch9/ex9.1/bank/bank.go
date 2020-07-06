package bank

// Exercise 9.1: Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1
// program. The result should indicate whether the transaction succeeded or failed due to
// insufficient funds. The message sent to the monitor goroutine must contain both the amount to
// withdraw and a new channel over which the monitor goroutine can send the boolean result
// back to Withdraw.

type _WithDrawInfo struct {
	amount int
	success chan bool
}

var deposits = make(chan int)
var balances = make(chan int)
var withdraw = make(chan _WithDrawInfo)

func Deposit(amount int) { deposits <- amount }

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraw <- _WithDrawInfo{amount: amount, success: ch}
	return <-ch
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case info := <-withdraw:
			if balance < info.amount {
				info.success <- false
			} else {
				balance -= info.amount
				info.success <- true
			}
		case balances <- balance:

		}
	}
}

func init() {
	go teller()
}
