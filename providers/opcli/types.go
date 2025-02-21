package opcli

type NoteField struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type Note struct {
	Id     string      `json:"id"`
	Fields []NoteField `json:"fields"`
}

type Account struct {
	AccountUUID string `json:"account_uuid"`
	Email       string `json:"email"`
	Shorthand   string `json:"shorthand"`
	Url         string `json:"url"`
	UserUUID    string `json:"user_uuid"`
}
