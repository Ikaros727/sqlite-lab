package types

const (
	StatusCodeSuccess    = "success"
	StatusMessageSuccess = "成功"
)

func StatusSuccess() Status {
	return Status{
		Code:    StatusCodeSuccess,
		Message: StatusMessageSuccess,
	}
}
