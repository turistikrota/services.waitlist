package res

type Response interface{}

type response struct{}

func New() Response {
	return &response{}
}
