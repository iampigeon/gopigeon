package main

import "fmt"
import "github.com/iampigeon/pigeon-go"

func main() {
	fmt.Println("vim-go")
	cl, err := iampigeon.NewPigeon("12345")
	if err != nil {
		panic(err)
	}

	s := cl.Subjects[0]
	fmt.Println(s.Name)

	payload := make(map[string]interface{})
	payload["foo"] = "bar"
	payload["baz"] = "zar"

	mqtt := &iampigeon.MQTT{
		Topic:   "POCPIGEON",
		Payload: payload,
	}

	m := &iampigeon.Message{
		SubjectName: s.Name,
		MQTT:        mqtt,
	}

	res, err := cl.Deliver(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(2)
	fmt.Println(res)

}
