package controller

import "fmt"

type IndexController struct {
}

func (c *IndexController) Index() {
	fmt.Println("router-index")
}
