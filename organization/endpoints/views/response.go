package views

type Token struct {
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
	Err     error  `json:"err,omitempty"`
}

type Msg struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
