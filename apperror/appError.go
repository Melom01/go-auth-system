package apperror

type AppCustomError struct {
	message    string
	statusCode int
	parameters map[string]interface{}
}

func (err AppCustomError) Error() string {
	return err.message
}

func (err AppCustomError) Status() int {
	return err.statusCode
}

func (err AppCustomError) Params() map[string]interface{} {
	return err.parameters
}