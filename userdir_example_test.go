package userdir_test

import "fmt"
import "github.com/wilmoore/go-userdir"

func Example() {
	fmt.Println("Home Directory:", userdir.Home())
	fmt.Println("Authorized Keys:", userdir.AuthorizedKeys())
}
