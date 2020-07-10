package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), 1)
}

func formatStruct(v reflect.Value) string {
	b := &bytes.Buffer{}
	b.WriteByte('{')
	for i := 0; i < v.Type().NumField(); i++ {
		if i != 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(b, "%s: %s", v.Type().Field(i).Name, formatAtom(v.Field(i)))
	}

	b.WriteByte('}')

	return b.String()
}

func formatArray(v reflect.Value) string {
	b := &bytes.Buffer{}
	b.WriteByte('[')
	for i := 0; i < v.Type().Len(); i++ {
		if i != 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(b, "%s", formatAtom(v.Index(i)))
	}
	b.WriteByte(']')
	return b.String()
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'e', 7, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'e', 7, 64)
	case reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("%g", v.Complex())
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		return formatStruct(v)
	case reflect.Array:
		return formatArray(v)
	default:
		return v.Type().String() + " value"
	}
}

func display(path string, v reflect.Value, level int) {
	const maxLevel = 3
	if level > maxLevel {
		fmt.Printf("exceeded maximun level")
		fmt.Printf("%s = %s\n", path, formatAtom(v))
		return
	}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), level + 1)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), level+1)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key), level+1)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), level+1)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s=nil", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), level+1)
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
