package auto

import "api/models"

var users = []models.User{
	models.User{Nickname: "john", Email: "johndoe@email.com", Password: "123456789"},
}

var posts = []models.Post{
	models.Post{Title: "john", Content: "john"},
}
