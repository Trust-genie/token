package helper

//generic error handling
//Will panic if error is generated
func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
