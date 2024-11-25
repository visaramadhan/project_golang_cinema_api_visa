package cinemas

type Cinema struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Seats    int    `json:"seats"`
}
