package responses

type Response struct {
	Data      interface{} `json:"data"`
	RecordSet interface{} `json:"record_set"`
	Message   string      `json:"message"`
}

type ErrorMessage struct {
	Name  string      `json:"name"`
	Desc  string      `json:"desc"`
	Error interface{} `json:"error"`
}
