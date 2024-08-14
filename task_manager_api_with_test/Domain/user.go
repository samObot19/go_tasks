package domain

type User struct {
    UserName string `bson:"username,omitempty" json:"username,omitempty"`
    Password string `bson:"password,omitempty" json:"password,omitempty"`
    Role     string `bson:"role,omitempty" json:"role,omitempty"`
}
