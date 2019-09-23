package muget

type FileData struct {
	Name     string
	Size     uint
	Dirname  string
	FullName string
}

func Run(filePath, outPutPath string) error {
	//TODO: 一旦普通のダウンロード
	if err := DownloadFile(filePath, outPutPath); err != nil {
		return err
	}
	return nil
}
