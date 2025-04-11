package main

import (
	"fmt"

	game_lib "github.com/NikosGour/Game-Engine/src/lib"
	platform "github.com/NikosGour/Game-Engine/src/lib/platform"
	platform_types "github.com/NikosGour/Game-Engine/src/lib/platform/types"
	log "github.com/NikosGour/logging/src"

	"github.com/NikosGour/testbed/src/build"
)

func main() {
	fmt.Printf("DEBUG_MODE = %t\n", build.DEBUG_MODE)
	game_lib.Add(2, 3)
	plat_state := &platform_types.PlatformState{}
	err := platform.Platform_startup(plat_state, "Nikos", 0, 0, 800, 600)
	if err != nil {
		log.Fatal("%s", err)
	}

}
