package req

type Request interface {
	Join() *JoinRequest
	Leave() *LeaveRequest
}

type request struct{}

func New() Request {
	return &request{}
}

func (r *request) Join() *JoinRequest {
	return &JoinRequest{}
}

func (r *request) Leave() *LeaveRequest {
	return &LeaveRequest{}
}
