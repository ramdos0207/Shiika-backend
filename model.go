package main

type User struct {
	Name     string `json:"name,omitempty" db:"name"`
	Password string `json:"password,omitempty"  db:"password"`
}
type Kaminoku struct {
	Id      string `json:"id,omitempty" db:"id"`
	Content string `json:"name,omitempty" db:"content"`
	Userid  string `json:"userid,omitempty" db:"userid"`
}
type Simonoku struct {
	Id         string `json:"id,omitempty" db:"id"`
	Content    string `json:"name,omitempty" db:"content"`
	KaminokuId string `json:"kaminokuid,omitempty" db:"kaminokuid"`
	Userid     string `json:"userid,omitempty" db:"userid"`
}
type LoginRequestBody struct {
	Username string `json:"username,omitempty" form:"username"`
	Password string `json:"password,omitempty" form:"password"`
}

// for post kaminoku's request body's binding
type KaminokuReq struct {
	Content string `json:"content,omitempty"`
}

// kaminoku for response
type KaminokuRes struct {
	Id       string `json:"id,omitempty"`
	Content  string `json:"name,omitempty"`
	UserName string `json:"username,omitempty"`
}

// for post simonoku's request body's binding
type SimonokuReq struct {
	Content string `json:"content,omitempty"`
}

type SimonokuRes struct {
	Id         string `json:"id,omitempty"`
	Content    string `json:"name,omitempty"`
	KaminokuId int    `json:"kaminokuid,omitempty"`
	UserName   string `json:"username,omitempty"`
}

type TankaRes struct {
	Kaminoku KaminokuRes
	Simonoku []SimonokuRes
}
