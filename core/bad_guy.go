package core

import "time"

type BadGuy struct {
	currentLevel       int     // Current level of the game
	currentPosition    float64 // Current position of the bad guy
	nextPosition       float64 // Next position of the bad guy, will be lerped to
	shouldMoveTimer 	 time.Timer
	velocity           float64 // Speed of the bad guy
	velocityMultiplier float64 // Multiplier for the velocity -- will be calculated by the game somehow TBD lol
	shouldDropBomb     bool    // Should the bad guy drop a bomb? Will be determined by a timer
	bombDropTimer      time.Timer
}

func NewBadGuy(level int) *BadGuy {
	return &BadGuy{
		currentLevel:       level,
		currentPosition:    0,
		nextPosition:       0,
		velocity:           0,
		velocityMultiplier: 0,
		shouldDropBomb:     false,
	}
}
