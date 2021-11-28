package sysError

// ReceiveFileError 以表单形式提交文件至服务器时 服务器接收文件失败的系统错误
type ReceiveFileError struct {
	RealError error
}

func (r *ReceiveFileError) Error() string {
	return r.RealError.Error()
}
