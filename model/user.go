package model

type User struct {
	Name      string
	Email     string
	Sexo      string
	Age       int
	Phone     int
	Residence Residence
}

type Residence struct {
	Street  string
	City    string
	Country string
}
