package Models

type Player struct {
	Id          int
	Name        string
	ShirtNumber int
	Address     *Address
}
type Address struct {
	Country string
	State   string
	City    string
}
