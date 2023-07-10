package shared

// a function to create a *[]*string
func StringSlice(s string) *[]*string {
	return &[]*string{&s}
}
