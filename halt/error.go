package helper

//Generic Error handling
func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
