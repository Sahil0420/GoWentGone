package helpers

import "fmt"

type Person struct {
	Name string
	Age  uint8
	Job  string
}

func (p Person) Unpack() (string, uint8, string) {
	return p.Name, p.Age, p.Job
}

func (p Person) Describe() string {
	return fmt.Sprintf("%s is %d years old and works as a %s", p.Name, p.Age, p.Job)
}

func (p *Person) UpdateJob(newJob string) {
	p.Job = newJob
}
