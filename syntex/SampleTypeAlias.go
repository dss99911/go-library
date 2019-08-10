package syntex

type String string
type Int int

func (s String) toInt() Int {
	return Int(1)
}

func (s Int) toString() String {
	return String(s)
}

func SampleAlias() {
	str := String("dsf")
	str.toInt().toString()
}
