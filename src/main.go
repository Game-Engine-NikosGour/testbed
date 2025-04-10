package main

import (
	"fmt"

	game_lib "github.com/NikosGour/Game-Engine/src/lib"
	"github.com/NikosGour/Game-Engine/src/lib/platform"

	"github.com/NikosGour/testbed/src/build"
)

func main() {
	fmt.Printf("DEBUG_MODE = %t\n", build.DEBUG_MODE)
	game_lib.Add(2, 3)
	plat_state := &platform.PlatformState{}
	platform.Platform_startup(plat_state, "Nikos", 0, 0, 800, 600)

}
