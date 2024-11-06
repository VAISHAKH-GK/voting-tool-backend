package user

type User struct {
	Id       string `json:"id,omitempty" form:"id,omitempty" bson:"id,omitempty"`
	Email    string `json:"email" form:"email" bson:"email"`
	Username string `json:"username" form:"username" bson:"username"`
	Password string `json:"password,omitempty" form:"password,omitempty" bson:"password,omitempty"`
}
