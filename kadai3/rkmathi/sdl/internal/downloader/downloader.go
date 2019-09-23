package downloader

import (
	crand "crypto/rand"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"
)

type Downloader struct {
	parallel  int
	targetDir string
	url       string
	tmpDir    string
}

func CreateTmpDir() (string, error) {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	tmpDir := filepath.Join(os.TempDir(), fmt.Sprintf("tmp%d", rand.Int63()))
	if err := os.Mkdir(tmpDir, 0777); err != nil {
		return "", err
	}

	return tmpDir, nil
}

func RemoveTmpDir(tmpDir string) {
	_ = os.RemoveAll(tmpDir)
}

func NewDownloader(parallel int, targetDir, url, tmpDir string) *Downloader {
	return &Downloader{parallel, targetDir, url, tmpDir}
}

func (d *Downloader) Download() error {
	contentLength, err := d.getContentLength()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for i := 0; i < d.parallel; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := d.doPartialRequest(contentLength, i, (i+1 == d.parallel))
			if err != nil {
				_ = fmt.Errorf("%v\n", err)
			}
		}(i)
	}

	wg.Wait()

	r, _ := http.NewRequest("GET", d.url, nil)
	err = d.writeToFile(path.Base(r.URL.Path))
	if err != nil {
		return err
	}

	return nil
}

func (d *Downloader) getContentLength() (int64, error) {
	resp, err := http.Head(d.url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("status is not 200")
	}

	return resp.ContentLength, nil
}

func (d *Downloader) doPartialRequest(conLen int64, i int, last bool) error {
	div := conLen / int64(d.parallel)
	begin := int64(i) * div
	end := int64(0)
	if last {
		end = conLen - 1
	} else {
		end = int64(i+1)*div - 1
	}
	req, err := http.NewRequest("GET", d.url, nil)
	if err != nil {
		return err
	}

	h := fmt.Sprintf("bytes=%d-%d", begin, end)
	req.Header.Add("Range", h)

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	tmpFilePath := filepath.Join(d.tmpDir, strconv.Itoa(i))
	tmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	_, err = tmpFile.Write(body)
	if err != nil {
		return err
	}

	return nil
}

func (d *Downloader) writeToFile(filename string) error {
	targetPath := filepath.Join(d.targetDir, filename)
	f, err := os.OpenFile(targetPath, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := 0; i < d.parallel; i++ {
		bytes, err := ioutil.ReadFile(
			filepath.Join(d.tmpDir, fmt.Sprintf("%d", i)))
		if err != nil {
			return err
		}

		_, err = f.Write(bytes)
		if err != nil {
			return err
		}
	}

	fmt.Printf("save to %s\n", targetPath)

	return nil
}
