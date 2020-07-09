package main

type _Bar struct {
	name string
	age int
}

type _Foo struct {
	a map[string]string
	b int
	c *string
	d map[_Bar]string
	e map[[3]string]string
}

func main() {
	a := make(map[string]string)
	a["a"] = "alpha"
	a["b"] = "beta"

	c := "hello"

	bar1 := _Bar{"a", 1}
	bar2 := _Bar{"b", 2}
	d := make(map[_Bar]string)
	d[bar1] = "bar1"
	d[bar2] = "bar2"

	e := make(map[[3]string]string)
	e[[3]string{"alpha", "beta", "gamma"}] = "greek letter"

	foo := _Foo{
		a: a,
		b: 2,
		c: &c,
		d: d,
		e: e,
	}

	Display("foo", foo)
}
