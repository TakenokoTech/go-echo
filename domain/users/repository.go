package users

type Repository interface {
	SelectAll() (user []User)
	SelectById(id string) (user User)
	Insert(user User)
}
