package main

import (
	"desafioPG/controller"
	"fmt"
)

func main() {
	// conf, err := config.GetConf()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(conf)

	dadosArquivos, err := controller.ReadFiles()
	if err != nil {
		fmt.Println(err)
		return
	}
	for index, line := range dadosArquivos {
		fmt.Printf("[%d] %s \n", index, line)
	}
}
