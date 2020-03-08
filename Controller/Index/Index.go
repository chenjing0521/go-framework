package Index

import "fmt"

func init() {

}

type IndexController struct {
}

func (c *IndexController) Index() {
	fmt.Println("test")
}
