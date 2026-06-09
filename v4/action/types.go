package action

type Code int

const (
	SessionTimeout Code = 1008
)

// Status 操作结果
type Status struct {
	Result  int    `json:"Result"`
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func (r *Status) Ok() bool {
	return r.Code == 0
}

func (r *Status) Is(expectedCode Code) bool {
	return r.Code == expectedCode
}

func Join[T ~string](src []T, sep string, defaultValue T) string {
	if len(src) == 0 {
		return string(defaultValue)
	}

	var (
		sep_ = []byte(sep)
		// preallocate for len(sep) + assume at least 1 character
		out = make([]byte, 0, (1+len(sep_))*len(src))
	)
	for _, s := range src {
		out = append(out, s...)
		out = append(out, sep_...)
	}

	return string(out[:len(out)-len(sep)])
}
