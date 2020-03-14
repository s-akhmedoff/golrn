package internship

import (
	"fmt"
	"strconv"
)

const (
	LOW_RANK = iota
	MEDIUM_RANK
	HIGH_RANK
)

type Getters interface {
	getProgLanguage() string
	getStatus() string
	getTask() string
}

type Developer struct {
	progLanguage, status string
	currentTask          task
}

type Helper struct {
	Developer
	rank int
}

func NewDeveloper(prog_language, status string, task task) *Developer {
	return &Developer{
		progLanguage: prog_language,
		status:       status,
		currentTask:  task,
	}
}

func NewHelper(prog_language, status string, task task, rank int) *Helper {
	return &Helper{
		Developer: Developer{
			progLanguage: prog_language,
			status:       status,
			currentTask:  task,
		},
		rank: rank,
	}
}

func (d Developer) getProgLanguage() string {
	return d.progLanguage
}

func (d Developer) getStatus() string {
	return d.status
}

func (d Developer) getTask() string {
	return d.currentTask.context
}

func (d *Developer) setProgLanguage(prog_language string) (old, new string) {
	defer func() {
		d.progLanguage = prog_language
	}()
	return d.progLanguage, prog_language
}

func (d *Developer) setStatus(status string) (old, new string) {
	defer func() {
		d.status = status
	}()
	return d.status, status
}

func (d *Developer) setTask(task task) (old, new task) {
	defer func() {
		d.currentTask = task
	}()
	return d.currentTask, task
}

func (d Developer) String() string {
	return "Developer codes on " + d.progLanguage + "\nHe is " + d.status + " now\nHis task is: " + d.currentTask.context + "\n"
}

func (h Helper) String() string {
	return "Helper codes on " + h.progLanguage + "\nHe has rank: " + strconv.Itoa(h.rank) + "\nHe is " + h.status + " now\nHis task is: " + h.currentTask.context + "\n"
}

func getProgLanguage(g Getters) string {
	return g.getProgLanguage()
}

func getStatus(g Getters) string {
	return g.getStatus()
}

func getTask(g Getters) string {
	return g.getTask()
}

func OOP() {
	dev1 := NewDeveloper("C++", "free", task{context: "Build project and Debug"})
	helper1 := NewHelper("C", "free", task{context: "Write Makefile"}, MEDIUM_RANK)
	fmt.Println(getProgLanguage(helper1), getProgLanguage(dev1))
	fmt.Println(dev1)
	fmt.Println(helper1)
	fmt.Println(getTask(dev1), getTask(helper1))

}
