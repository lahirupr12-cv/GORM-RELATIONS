package Utils

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
