package puppy

import (
	"github.com/MayvenF/dog"
)

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}
