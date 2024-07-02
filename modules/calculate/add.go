package calculate

const Foo = 10

type Player struct {
	attackPower int // private field
	Name        string
}

// private function
func add(a, b int) int {
	return a + b
}

// public function
func Add(a, b int) int {
	return a + b
}
