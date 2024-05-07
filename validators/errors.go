package validators

type Error string

func (err Error) Error() string {
	return string(err)
}

var (
	ErrorPasswordNotStrong Error = " Password must be at least 8 characters long and contain at least one lowercase and uppercase letters and numbers. Special characters like $, #, @, !, %, ^, &, *, (,) "
)
