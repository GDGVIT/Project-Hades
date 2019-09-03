package views

type Token struct {
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
	Err     string `json:"err"`
}

type Msg struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
