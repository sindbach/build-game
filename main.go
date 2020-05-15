
package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten"
	"github.com/sindbach/build-game/builders"
)

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	ebiten.SetWindowSize(builders.ScreenWidth*2, builders.ScreenHeight*2)
	ebiten.SetWindowTitle("Rise Of The Fish")
	if err := ebiten.RunGame(&builders.Game{}); err != nil {
		log.Fatal(err)
	}
}