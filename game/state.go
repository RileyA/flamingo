package game

import "time"

type State interface {
	Init(*Game)
	Update(time.Duration) (bool, error)
	Deinit()
}
