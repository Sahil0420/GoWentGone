package helpers

type Person struct {
	Name string
	Age  uint8
	Job  string
}

func (p Person) Unpack() (string, uint8, string) {
	return p.Name, p.Age, p.Job
}
