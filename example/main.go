package main

import "flag"
import "fmt"
import "github.com/iampigeon/pigeon-go"

func main() {
	host := flag.String("host", "http://localhost:9000", "pigeon host address")
	fmt.Println("Pigeon example")
	cl, err := iampigeon.NewPigeon("12345", *host)
	if err != nil {
		panic(err)
	}

	s := cl.Subjects[0]
	fmt.Println(s.Name)

	payload := make(map[string]interface{})
	payload["foo"] = "bar"
	payload["baz"] = "zar"

	mqtt := &iampigeon.MQTT{
		Payload: payload,
	}

	chans := &iampigeon.Channels{MQTT: mqtt}

	m := &iampigeon.Message{
		SubjectName: s.Name,
		Channels:    chans,
	}

	res, err := cl.Deliver(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(2)
	fmt.Println(res)

	rm := new(iampigeon.ResponseMessages)

	if err := iampigeon.Decode(res.Data, rm); err != nil {
		panic(err)
	}

	fmt.Println(3)
	for _, msg := range rm.Messages {
		pr, err := cl.FetchStatusByID(msg.ID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", pr.Data["status"])

		pr, err = cl.CancelMessage(msg.ID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", pr.Data["status"])

		pr, err = cl.FetchStatusByID(msg.ID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", pr.Data["status"])
	}
}

/*
{
  "data": {
    "messages": [{
      "id": "id:01CRVFYX6TG8E7ETKRQ4H21S8R",
      "channel": "mqtt"
    }, {
      "id": "id:0YGFDLJUCFG8E7ETHJDHYTOSKFJ",
      "channel": "sms",
    }, {
      "error": "internal server error",
      "channel": "mandrill"
    }]
  }
}
*/
