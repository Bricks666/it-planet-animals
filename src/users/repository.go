package users

var A = 15

//reform:users
type User struct {
	id        uint   `reform:"id, pk"`
	email     string `reform:"email"`
	password  string `reform:"password"`
	firstName string `reform:"firstName"`
	lastName  string `reform:"lastName"`
}
