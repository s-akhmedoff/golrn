package internship

import "testing"

var tempList taskList

func BenchmarkTaskList_AddTask(b *testing.B) {
	tempTask := task{id: 0, priority: LOW, context: "null"}
	for i := 0; i < b.N; i++ {
		tempList.AddTask(tempTask)
	}
}

func TestTaskList_AddTask(t *testing.T) {
	tempTask := task{id: 0, priority: LOW, context: "null"}
	result := tempList.AddTask(tempTask)

	if result != tempTask {
		t.Errorf("taskList.addTask(%v) = %v; but expected %v\n", tempTask, result, tempTask)
	}
}

func TestTaskList_RemoveTask(t *testing.T) {
	tempTask := task{id: 0, priority: LOW, context: "null"}
	tempList.AddTask(tempTask)
	result := tempList.RemoveTask(0); if result != nil {
		t.Errorf("%v", result)
	}
}
