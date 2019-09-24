package muget

import (
	"fmt"
	"io"
	"os"
)

func BindwithFiles(count int, ext string) error {

	fmt.Println("\nbinding with files...")

	fh, err := os.Create("gogogo"+ ext)
	if err != nil {
		return err
	}
	defer fh.Close()

	var f string
	for i := 0; i < count; i++ {
		f = fmt.Sprintf("./%d%s", i,ext)
		subfp, err := os.Open(f)
		if err != nil {
			return err
		}

		io.Copy(fh, subfp)

		// Not use defer
		subfp.Close()

		// remove a file in download location for join
		if err := os.Remove(f); err != nil {
			return err
		}
	}

	fmt.Println("Complete")

	return nil
}
