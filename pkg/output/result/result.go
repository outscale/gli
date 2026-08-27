package result

type Result struct {
	// Iter defines the iteration counter in a --watch/--wait call
	Iter        int
	SingleEntry bool
	Ok          any
	Error       error
}

func Error(err error) Result {
	return Result{Error: err}
}

func New(old Result, newOk any) Result {
	return Result{Ok: newOk, Iter: old.Iter, SingleEntry: old.SingleEntry}
}
