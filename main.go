package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	sm "swaying-memory/swaying-memory"

	"github.com/hajimehoshi/ebiten/v2"
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

	ebiten.SetWindowSize(sm.ScreenWidth*2, sm.ScreenHeight*2)
	ebiten.SetWindowTitle("Swaying Memory")
	if err := ebiten.RunGame(&sm.Game{}); err != nil {
		log.Fatal(err)
	}
}
