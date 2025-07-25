// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Query struct {
}

type Todo struct {
	ID     int32  `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID int32  `json:"user_id"`
	User   *User  `json:"user"`
}

type User struct {
	ID       int32   `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Token    string  `json:"token"`
	Todos    []*Todo `json:"todos"`
}
