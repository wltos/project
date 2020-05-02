package serverchan

type rsp struct {
	ErrNo   int64  `json:"errNo"`
	ErrMsg  string `json:"errMsg"`
	DataSet string `json:"dataSet"`
}
