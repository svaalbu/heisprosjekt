package driver

const N_FLOORS = 4

// Number of buttons (and corresponding lamps) on a per-floor basis
const N_BUTTONS = 3

type MotorDirection int8

const (
	DirnDown = -1 + iota
	DirnStop
	DirnUp
)

type ButtonType int

const (
	ButtonCallUp = iota
	ButtonCallDown
	ButtonCommand
)
