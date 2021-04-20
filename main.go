/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Torch-Labs/narwhal/config"
)

func main() {
	config := config.Config{}

	// dirname, err := os.UserHomeDir()
	// if err != nil {
	// 	panic(err)
	// }

	data, err := ioutil.ReadFile("/Users/sharithg/go/src/narwhal/.narwhal_config.yml")
	if err != nil {
		panic(err)
	}

	err = config.SetFromBytes(data)
	if err != nil {
		panic(err)
	}

	pr, err := config.Get("services")
	if err != nil {
		panic(err)
	}
	fmt.Println(pr)
}
