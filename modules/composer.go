package modules

import (
	"fmt"

	"github.com/chenjiandongx/oscar/fixtures"
)

type composer struct{}

func (mod *composer) Name() string {
	return ModComposer
}

func (mod *composer) Display() {
	stage := "Installing"

	for _, idx := range GenRandomIndex(10, 100, len(fixtures.Composer)) {
		fmt.Printf("  - %s ", stage)
		PrintGreen("%s ", fixtures.Composer[idx])
		PrintYellow("(%s)", GenPackageVersion())
		fmt.Println(": Loading from cache")
		SleepInMills(100, 2000)
	}

	PrintGreen("Writing lock file\n")
	PrintGreen("Generating autoload files\n")
}

func NewComposerModule() Moduler {
	return &composer{}
}
