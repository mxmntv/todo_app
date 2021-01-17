package repository

/*Auther - */
type Auther interface {
}

/*Lister - */
type Lister interface {
}

/*Itemer - */
type Itemer interface {
}

/*Repository - */
type Repository struct {
	Auther
	Lister
	Itemer
}

/*NewRepository - */
func NewRepository() *Repository {
	return &Repository{}
}
