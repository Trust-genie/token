package prog

func Halt(err error) {
	if err != nil {
		panic(err)
	}
}
