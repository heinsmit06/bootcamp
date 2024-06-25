package bootcamp

type Lock struct {
	// Define fields here
	locked bool
}

func (l *Lock) Lock() bool {
	if l.locked == true {
		return false
	} else {
		l.locked = true
		return true
	}
}

func (l *Lock) Unlock() bool {
	if l.locked == false {
		return false
	} else {
		l.locked = false
		return true
	}
}

func (l Lock) IsLocked() bool {
	return l.locked
}

func NewLock() Lock {
	lock := Lock{
		locked: false,
	}
	return lock
}
