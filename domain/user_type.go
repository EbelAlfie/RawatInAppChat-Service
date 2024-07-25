package domain

type User struct {
	userid   string `json:"id"`
	username string `json:"user"`
	password string `json:"pass"`
}

/*
Struct tags such as json:"artist"
specify what a field’s name should be when the struct’s
contents are serialized into JSON. Without them,
the JSON would use the struct’s
capitalized field names – a style not as common in JSON.
*/
