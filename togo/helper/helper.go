package helper

func Validate(firstName string) bool {
	return len(firstName) >= 2
}
