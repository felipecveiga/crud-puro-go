package model

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Sexo      string    `json:"sexo"`
	Age       int       `json:"age"`
	Phone     int       `json:"phone"`
	Residence Residence `json:"residence"`
}

type Residence struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	Number  int    `json:"number"`
}
