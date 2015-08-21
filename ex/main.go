package main

import (
	"fmt"

	fab ".."
)

func main() {
	fab := fab.NewSuperFab()
	fmt.Println(fab.Paint("Making my outputs fabulous!!!"))
}
