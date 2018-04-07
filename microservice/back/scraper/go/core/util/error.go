package util

// Errable represents an interface to check Errors
type Errable interface {
	// ToJsonQ get JsonQuery from string
	CheckError(err error)
}

// CheckError is un panic when err not is nil
func CheckError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
