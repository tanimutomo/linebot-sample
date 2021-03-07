package db

var users = make(map[int]*User)

type User struct {
	ID         int
	LINEUserID string
	Name       string
}

func CreateUser(luid, name string) *User {
	u := &User{
		ID:         len(users) + 1,
		LINEUserID: luid,
		Name:       name,
	}
	users[u.ID] = u
	return u
}

func FindUser(id int) *User {
	u, ok := users[id]
	if !ok {
		return nil
	}
	return u
}
