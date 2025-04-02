package main

import (
	"fmt"

	game_lib "github.com/NikosGour/Game-Engine/src/lib"
	"github.com/NikosGour/testbed/src/build"
)

func main() {
	fmt.Printf("DEBUG_MODE = %t\n", build.DEBUG_MODE)
	game_lib.Add(2, 3)

}
