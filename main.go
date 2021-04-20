package main

import (
	"github.com/Torch-Labs/narwhal/cmd"
)

func main() {

	cmd.Execute()

	// config := config.Config{}

	// data, err := ioutil.ReadFile("/Users/sharithg/go/src/narwhal/.narwhal_config.yml")
	// if err != nil {
	// 	panic(err)
	// }

	// err = config.SetFromBytes(data)
	// if err != nil {
	// 	panic(err)
	// }

	// pr, err := config.Get("services")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(pr)
}
