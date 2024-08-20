package helper

import "fmt"

type Record struct {
	id   int
	task string
	done bool
}

func (r Record) toString() []string {
	return []string{fmt.Sprintf("%v", r.id), r.task, fmt.Sprintf("%v", r.done)}
}
