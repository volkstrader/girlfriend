package girlfriend

const old = 24

type Profile struct {
	Age int
}

func (v Profile) Breakup() bool {
	return v.Age > old
}
