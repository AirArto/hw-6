package goroutine

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	toFile, err := os.OpenFile("testFromFile", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Errorf("\n\t%s", err)
	}
	buffer := make([]byte, 50<<20)
	_, err = toFile.Write(buffer)
	toFile.Close()

	_, err = Copy("testFromFile", "testToFile", 5<<20, 2<<20)
	if err != nil {
		t.Errorf("\n\t%s", err)
	}

	toFile, err = os.OpenFile("testToFile", os.O_RDONLY, 0644)
	fileInfo, err := toFile.Stat()
	fileSize := fileInfo.Size()
	toFile.Close()
	if fileSize != 5<<20 {
		t.Errorf("\n\t%s", "Copy Error")
	}

	_, err = Copy("testFromFile", "testToFile", 12<<20, 0)
	if err != nil {
		t.Errorf("\n\t%s", err)
	}

	toFile, err = os.OpenFile("testToFile", os.O_RDONLY, 0644)
	fileInfo, err = toFile.Stat()
	fileSize = fileInfo.Size()
	toFile.Close()
	if fileSize != 12<<20 {
		t.Errorf("\n\t%s", "Copy Error")
	}

	_, err = Copy("testFromFile", "testToFile", 0, 0)
	if err != nil {
		t.Errorf("\n\t%s", err)
	}

	toFile, err = os.OpenFile("testToFile", os.O_RDONLY, 0644)
	fileInfo, err = toFile.Stat()
	fileSize = fileInfo.Size()
	toFile.Close()
	if fileSize != 50<<20 {
		t.Errorf("\n\t%s", "Copy Error")
	}
}
