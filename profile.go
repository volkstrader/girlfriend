package girlfriend

const old = 25

type Profile struct {
	Age int
}

func (v Profile) Breakup() bool {
	return v.Age >= old
}

func (v Profile) Makeup(ok bool) bool {
	return ok
}
