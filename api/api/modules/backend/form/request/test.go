package request

type TestReq struct {
	Content string `json:"content" validate:"required,max=150"`
}
