package todo

/*TodoList- */
type TodoList struct {
	ID          int    `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
}

/*UsersList-*/
type UsersList struct {
	ID     int
	ListId int
	UserId int
}

/*TodoItem-*/
type TodoItem struct {
	ID          int    `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Done        bool   `json: "done"`
}

/*ListsItem-*/
type ListsItem struct {
	ID     int
	ListId int
	ItemId int
}
