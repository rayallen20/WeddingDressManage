package wdmError

// DBError 数据库错误
type DBError struct {
	Message string
}

func (d DBError) Error() string {
	return d.Message
}

// FileTypeError 上传文件类型错误
type FileTypeError struct {
	Message string
}

func (f FileTypeError) Error() string {
	return f.Message
}

// SaveFileError 保存文件错误
type SaveFileError struct {
	Message string
}

func (s SaveFileError) Error() string {
	return s.Message
}
