package httputil

const (
	ErrorJsonMarshalFailed             = "error.json.marshal.failed"
	ErrorJsonDecodeFailed              = "error.json.decode.failed"
	ErrorContentTypeNotApplicationJson = "error.content.type.not.application.json"

	ContentTypeHeader          = "Content-Type"
	ContentTypeApplicationJson = "application/json"

	StatusCodeKey = "StatusCode"
)