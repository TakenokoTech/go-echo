package todos

type Repository interface {
	SelectAll() (todos []Todo)
	Insert(todo Todo)
}
