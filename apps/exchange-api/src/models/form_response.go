package models

type TemplateResponse struct {
	Message    string      `json:"message"`
	StatusCode interface{} `json:"status_code"`
	Data       interface{} `json:"data"`
}

func SetResponse(msg string, values ...interface{}) interface{} {
	var code interface{}
	var data interface{}

	switch len(values) {
	case 0:
		code, data = 200, nil
	case 1:
		code, data = values[0], nil
	default:
		code, data = values[0], values[1]
	}

	return TemplateResponse{
		Message:    msg,
		StatusCode: code,
		Data:       data,
	}
}

func SuccessResponse(msg string, data ...interface{}) interface{} {
	var setData interface{}
	if len(data) == 0 {
		setData = nil
	} else {
		setData = data[0]
	}
	return TemplateResponse{
		Message:    msg,
		StatusCode: 200,
		Data:       setData,
	}
}

func FailResponse(msg string) interface{} {
	return TemplateResponse{
		Message:    msg,
		StatusCode: 400,
		Data:       nil,
	}
}

func ErrResponse(err error) interface{} {
	return TemplateResponse{
		Message:    err.Error(),
		StatusCode: 500,
		Data:       nil,
	}
}
