package _const

var (
	TodoItemCreatedStatus = 0
	TodoItemDoingStatus   = 1
	TodoItemDoneStatus    = 2
	TodoItemDeletedStatus = 3
)

var (
	TodoItemStatusMap = map[int]string{
		TodoItemCreatedStatus: "created",
		TodoItemDoingStatus:   "doing",
		TodoItemDoneStatus:    "done",
		TodoItemDeletedStatus: "deleted",
	}
)

func GetTodoItemStatusText(status int) string {
	text, exist := TodoItemStatusMap[status]
	if exist {
		return text
	}
	return "unknown"
}
