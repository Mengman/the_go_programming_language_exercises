package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%s%v", string(u.op), u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("%v %s %v", b.x, string(b.op), b.y)
}

func (c call) String() string {
	argsStrs := make([]string, len(c.args))
	for i, arg := range c.args {
		argsStrs[i] = arg.String()
	}
	return fmt.Sprintf("%s(%v)", c.fn, strings.Join(argsStrs, ", "))
}
