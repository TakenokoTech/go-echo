package test

type test struct {
	Id     int    `json:"id"`
	Result string `json:"result"`
}

func NewTest(id int, result string) test {
	return test{
		Id:     id,
		Result: result,
	}
}
