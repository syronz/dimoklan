// package gp has all contsants related to gameplay, most important one is Rate which is the speed of the game
// As much is the rate is bigger the game would be slower. In case the rate is 0 means the game speed is instant
// which is not make sence and it affect the game status

package gp

import "time"

const (
	Rate = time.Second

	// move types
	relocate = "relocate"
	attack   = "attack"
	capture  = "capture"

	// direction
	Source      = "S"
	Destination = "D"
)

func MoveTypes() []string {
	return []string{relocate, attack, capture}
}

func Directions() []string {
	return []string{Source, Destination}
}
