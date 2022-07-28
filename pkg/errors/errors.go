package errors

type Response struct {
	Error Error `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func CreateResponse(err error) Response {
	var resp Response

	switch err.(type) {
	case *DuplicateDna:
		resp.Error.Code = 409
		resp.Error.Message = err.Error()
	case *NotMutant:
		resp.Error.Code = 403
		resp.Error.Message = err.Error()
	case *MethodNotAllowed:
		resp.Error.Code = 405
		resp.Error.Message = err.Error()
	case *BadRequest:
		resp.Error.Code = 400
		resp.Error.Message = err.Error()
	default:
		resp.Error.Code = 500
		resp.Error.Message = err.Error()
	}
	return resp
}
