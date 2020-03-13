package internship

import (
	"fmt"
)

const (
	LOW = iota
	MEDIUM
	HIGH
)

type contact struct {
	full_name string
	number    int
}

type task struct {
	id, priority int
	context      string
}

type contactList []contact
type taskList []task

func (t *taskList) addTask(tsk task) task {
	*t = append(*t, tsk)
	return tsk
}

func (t *taskList) removeTask(id int) bool {
	for i, v := range *t {
		if v.id == id {
			*t = append((*t)[:i], (*t)[i+1:]...)
			return true
		}
	}
	fmt.Println("No task with such ID")
	return false
}

func (t *taskList) editTask(id int, tsk task) bool {
	for i, v := range *t {
		if v.id == id {
			(*t)[i] = tsk
			return true
		}
	}
	fmt.Println("No task with such ID")
	return false
}

func (t *taskList) showTask(id int) bool {
	for _, v := range *t {
		if v.id == id {
			fmt.Printf("Task #%v\nPriority: %v\nContext: %v\n", v.id, v.priority, v.context)
			return true
		}
	}
	fmt.Println("No task with such ID")
	return false
}

func (t *taskList) showAllTasks() {
	for _, v := range *t {
		fmt.Printf("Task #%v\nPriority: %v\nContext: %v\n\n", v.id, v.priority, v.context)
	}
}

func (c *contactList) addContact(cnt contact) contact {
	*c = append(*c, cnt)
	return cnt
}

func (c *contactList) removeContact(number int) bool {
	for i, v := range *c {
		if v.number == number {
			*c = append((*c)[:i], (*c)[i+1:]...)
			return true
		}
	}
	fmt.Println("No contact with such number!")
	return false
}

func (c *contactList) editContact(number int, cnt contact) bool {
	for i, v := range *c {
		if v.number == number {
			(*c)[i] = cnt
			return true
		}
	}
	fmt.Println("Cannot edit contact with such number!")
	return false
}

func (c *contactList) showContact(number int) bool {
	for _, v := range *c {
		if v.number == number {
			fmt.Printf("Contact\nFull Name: %v\nNumber phone: %v\n", v.full_name, v.number)
			return true
		}
	}
	return false
}

func (c *contactList) showAllContacts() {
	for _, v := range *c {
		fmt.Printf("Contact\nFull Name: %v\nNumber: %v\n\n", v.full_name, v.number)
	}
}

func Structs() {
	project_tasks := taskList{}

	task1 := task{
		id:       0,
		priority: LOW,
		context:  "Make this task list!",
	}

	task2 := task{
		id:       1,
		priority: HIGH,
		context:  "Upgrade it!",
	}

	task3 := task{
		id:       2,
		priority: MEDIUM,
		context:  "Show it!",
	}

	project_tasks.addTask(task1)
	project_tasks.addTask(task2)
	project_tasks.addTask(task3)

	project_tasks.showAllTasks()

	project_tasks.removeTask(0)

	project_tasks.showAllTasks()

	project_tasks.editTask(2, task{
		id:       2,
		priority: HIGH,
		context:  "foobar!",
	})

	project_tasks.showAllTasks()

	my_contacts := contactList{}
	contact1 := contact{
		full_name: "Barak Obama",
		number:    909980898,
	}
	contact2 := contact{
		full_name: "Vladimir Putin",
		number:    970070707,
	}
	contact3 := contact{
		full_name: "Koronavirus",
		number:    712600202,
	}

	my_contacts.addContact(contact1)
	my_contacts.addContact(contact2)
	my_contacts.addContact(contact3)

	my_contacts.showAllContacts()
}
