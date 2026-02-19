package result

type Result struct {
	SingleEntry bool
	Ok          any
	Error       error
}
