package httpstarter

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	ChangeTracking
}

var AnonymousUser User = User{ID: 0, Email: "anonymous@dt.com"}
