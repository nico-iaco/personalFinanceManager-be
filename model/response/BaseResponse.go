package response

type BaseResponse[T any] struct {
	ResponseCode int    `json:"statusCode"`
	Body         T      `json:"body"`
	Message      string `json:"message"`
}
