package request

type CreateTaskReq struct {
	Content string `bind:"required" json:"content"`
}
