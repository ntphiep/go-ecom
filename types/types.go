package types


type User struct {
    ID          int    `json:"id"`
    FirstName   string `json:"firstName"`
    LastName    string `json:"lastName"`
    Email       string `json:"email"`
    Password    string `json:"-"`
    CreatedAt   string `json:"createdAt"`
}

type RegisterUserPayload struct {
    FirstName   string `json:"firstName"` 
    LastName    string `json:"lastName"`
    Email       string `json:"email"`
    Password    string `json:"password"`
}
