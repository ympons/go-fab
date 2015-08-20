package main

import (
	"fmt"

	fab ".."
)

func main() {
	fab := fab.NewFab()
	fmt.Println(fab.Paint("Making my outputs fabulous!!!"))
}
