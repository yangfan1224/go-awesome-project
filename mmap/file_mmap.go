package mmap

import (
	"os"
	"syscall"
)

const defaultMapSize = 10 * (1 << 5)

func WriteFile() error {
	file, err := os.OpenFile("nmmp.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	if _, err := file.WriteString("yangfan"); err != nil{
		return err
	}

	return file.Close()
}

func MmapFile() ([] byte, error){
	var data []byte
	file, err := os.OpenFile("nmmp.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	if data, err = syscall.Mmap(int(file.Fd()), 0, len("yangfan"), syscall.PROT_READ, syscall.MAP_SHARED); err != nil {
		return nil, err
	}
	return data, nil
}