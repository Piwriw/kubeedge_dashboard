package md5

import (
	"crypto/md5"
	"fmt"
)

const bufferSize = 65536

// MD5sum returns MD5 checksum of filename
func MD5sum(data []byte) (string, error) {
	//if info, err := os.Stat(filename); err != nil {
	//	return "", err
	//} else if info.IsDir() {
	//	return "", nil
	//}
	//
	//file, err := os.Open(filename)
	//if err != nil {
	//	return "", err
	//}
	//defer file.Close()

	hash := md5.New()
	//for buf, reader := make([]byte, bufferSize), bufio.NewReader(file); ; {
	//	n, err := reader.Read(buf)
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		}
	//		return "", err
	//	}

	hash.Write(data)
	//}

	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, nil
}
