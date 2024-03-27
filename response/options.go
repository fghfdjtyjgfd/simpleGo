package response

import "net/http"

func Success(msg string) Response {
	return Response{
		Status:  http.StatusOK,
		Message: msg,
	}
}

func NotFound() Response {
	return Response{
		Status:  http.StatusNotFound,
		Message: "Not Found",
	}
}

func InternalServerError() Response {
	return Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
}

func Custom(status int, msg string) Response {
	return Response{
		Status:  status,
		Message: msg,
	}
}
