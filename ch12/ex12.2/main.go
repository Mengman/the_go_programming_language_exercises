package main

type Cycle struct {
	Value int
	Tail *Cycle
}

func main() {
	var c Cycle
	c = Cycle{42, &c}
	Display("c", c)
}
