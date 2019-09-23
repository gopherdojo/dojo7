package muget

import "context"

type FileData struct {
	Name     string
	Size     uint
	Dirname  string
	FullName string
}

func Run(url, outPutPath string) error {
	//TODO: 一旦普通のダウンロード
	size, err := CheckRanges(context.Background(), url)
	if err != nil {
		return err
	}

	if err := DownloadFile(url, outPutPath, 0, size); err != nil {
		return err
	}
	return nil
}
