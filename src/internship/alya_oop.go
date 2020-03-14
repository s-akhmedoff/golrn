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
	getTask() task
}

type Developer struct {
	progLanguage, status string
	currentTask          task
}

type Helper struct {
	Developer
	rank int
}

func NewDeveloper(params ...interface{}) *Developer {
	developer := new(Developer)

	developer.progLanguage = params[0].(string)
	developer.status = params[1].(string)
	developer.currentTask = params[2].(task)

	return developer
}

func NewHelper(params ...interface{}) *Helper {
	helper := new(Helper)

	helper.progLanguage = params[0].(string)
	helper.status = params[1].(string)
	helper.currentTask = params[2].(task)
	helper.rank = params[3].(int)

	return helper
}

func (d Developer) getProgLanguage() string {
	return d.progLanguage
}

func (d Developer) getStatus() string {
	return d.status
}

func (d Developer) getTask() task {
	return d.currentTask
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

func getTask(g Getters) task {
	return g.getTask()
}

func OOP() {
	dev1 := NewDeveloper("C++", "free", task{context: "Build project and Debug"})
	helper1 := NewHelper("C", "free", task{context: "Write Makefile"}, MEDIUM_RANK)
	fmt.Println(dev1, helper1)

}
