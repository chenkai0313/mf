package validate

type SayHelloReq struct {
	Content string `json:"content"  validate:"required"`
}
