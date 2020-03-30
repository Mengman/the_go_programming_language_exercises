package main

// Exercise 5.11: The instructor of the linear algebra course decides that calculus is now a
// prerequisite. Extend the topoSort func tion to report cycles.


import (
	"fmt"
	"os"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"intro to programming":  {"programming languages"},
	"computer organization": {"computer organization"},
}

var coursePrereqs = make(map[string][]string)

func main() {
	sortedCourses, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for i, course := range sortedCourses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items string) error

	visitAll = func(item string) error {
		if !seen[item] {
			seen[item] = true
			for _, it := range m[item] {
				err := flatPrereqs(item, it)
				if err != nil {
					return err
				}
				for k, vals := range m {
					if containsStr(vals, item) {
						err = flatPrereqs(k, it)
						if err != nil {
							return err
						}
					}
				}
				err = visitAll(it)
				if err != nil {
					return err
				}
			}
			order = append(order, item)
		}
		return nil
	}

	for key := range m {
		err := visitAll(key)
		if err != nil {
			return nil, err
		}
	}

	return order, nil
}

func containsStr(strs []string, s string) bool {
	for _, item := range strs {
		if item == s {
			return true
		}
	}
	return false
}

func flatPrereqs(course, prereq string) error {
	ps := coursePrereqs[prereq]
	if containsStr(ps, course) {
		return fmt.Errorf("Found cycle prerequistes '%s' <--> '%s'\n", course, prereq)
	}
	coursePrereqs[course] = append(coursePrereqs[course], prereq)
	return nil
}
