package myhttp

type Result struct {
	Dto        any
	StatusCode int
}

var EmptyResult = Result{}

func NewOkResult(status int, dto any) Result {
	return Result{
		StatusCode: status,
		Dto:        dto,
	}
}
