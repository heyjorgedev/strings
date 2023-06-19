package strings

func For(str string) FluentString {
	return FluentString{str}
}

func Of(str string) FluentString {
	return FluentString{str}
}

type FluentString struct {
	str string
}

func (fs FluentString) Lower() FluentString {
	return For(Lower(fs.str))
}

func (fs FluentString) Upper() FluentString {
	return For(Upper(fs.str))
}

func (fs FluentString) Reverse() FluentString {
	return For(Reverse(fs.str))
}

func (fs FluentString) String() string {
	return fs.ToString()
}

func (fs FluentString) ToString() string {
	return fs.str
}

func (fs FluentString) Contains(needle string) bool {
	return Contains(fs.str, needle)
}

func (fs FluentString) ContainsSome(needles []string) bool {
	return ContainsSome(fs.str, needles)
}

func (fs FluentString) ContainsAll(needles []string) bool {
	return ContainsAll(fs.str, needles)
}

func (fs FluentString) Length() int {
	return Length(fs.str)
}

func (fs FluentString) After(after string) FluentString {
	return For(After(fs.str, after))
}

func (fs FluentString) Before(before string) FluentString {
	return For(Before(fs.str, before))
}
