package test

//自定义协议外壳
type Shell struct {
	Header string
	Kernel Protocol
}
//协议核心
type Protocol struct {
	Mark        int
	Function    string
	DataRequest interface{}
	DataResponse Response
}
//回复模板
type Response struct {
	Status int
	Msg string
    Data []interface{}
}
