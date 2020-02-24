package goroutine

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb"
)

//CopyFile
func Copy(from string, to string, limit int, offset int) (int, error) {
	fromFile, err := os.OpenFile(from, os.O_RDONLY, 0644)
	defer fromFile.Close()
	if err != nil {
		return 0, err
	}

	toFile, err := os.OpenFile(to, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	defer toFile.Close()
	if err != nil {
		return 0, err
	}

	fileInfo, err := fromFile.Stat()
	if err != nil {
		return 0, err
	}
	fileSize := int(fileInfo.Size())
	if offset >= fileSize {
		return 0, errors.New("Wrong Offset arg")
	}
	if fileSize == 0 {
		return 0, errors.New("Empty file")
	}
	if limit > fileSize-offset || limit == 0 {
		limit = fileSize - offset
	}

	bufferSize := 5 << 20
	buffer := make([]byte, bufferSize)
	bar := pb.StartNew(limit)
	written := 0
	currentOffset := offset

	for written != limit {
		readed, err := fromFile.ReadAt(buffer, int64(currentOffset))
		if err != nil && err != io.EOF {
			return written, err
		}
		currentOffset += readed
		if currentOffset > limit+offset {
			readed = limit + offset + readed - currentOffset
		}
		writed, err := toFile.Write(buffer[:readed])
		written += writed
		bar.Add(writed)

		if err != nil {
			return written, err
		}
	}
	bar.Finish()

	return 0, nil
}
