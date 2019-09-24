package muget

import (
	"fmt"
	"io"
	"os"
)

func MergeFiles(count int, fileName, ext string) (err error) {

	fh, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() {
		err = fh.Close()
		if err != nil {
			return
		}
	}()

	var f string
	for i := 0; i < count; i++ {
		f = fmt.Sprintf("./%d%s", i, ext)
		openFile, err := os.Open(f)
		if err != nil {
			return err
		}

		_, err = io.Copy(fh, openFile)
		if err != nil {
			return err
		}

		err = openFile.Close()
		if err != nil {
			return err
		}

		// remove a file in download location for join
		err = os.Remove(f)
		if err != nil {
			return err
		}
	}

	return
}
