package main

import (
	"github.com/secsy/goftp"
	"io"
	"os"
)

type ftpConnect struct {
	client *goftp.Client
}

func (client ftpConnect) login(url string, username string, password string) error {
	var err error
	client.client, err = goftp.DialConfig(goftp.Config{
		User:     username,
		Password: password,
	}, url)
	if err != nil {
		return err
	}
	return nil
}

func (client ftpConnect) CreateFile(fileName string) error {
	err := client.client.Retrieve(fileName, nil)
	if err != nil {
		return err
	}
	return nil
}

func (client ftpConnect) ReadFile(fileName string) (io.Reader, error) {
	var reader io.Reader
	err := client.client.Store(fileName, reader)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func (client ftpConnect) WriteFile(fileName string, writer io.Writer) error {
	err := client.client.Retrieve(fileName, writer)
	if err != nil {
		return err
	}
	return nil
}

func (client ftpConnect) DeleteFile(fileName string) error {
	err := client.client.Delete(fileName)
	if err != nil {
		return err
	}
	return nil
}

func (client ftpConnect) CreateDir(pathName string) error {
	_, err := client.client.Mkdir(pathName)
	if err != nil {
		return err
	}
	return nil
}

func (client ftpConnect) ReadDir(pathName string) ([]os.FileInfo, error) {
	res, err := client.client.ReadDir(pathName)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (client ftpConnect) DeleteDir(pathName string) error {
	err := client.client.Rmdir(pathName)
	if err != nil {
		return err
	}
	return nil
}

func (client ftpConnect) close() error {
	err := client.client.Close()
	if err != nil {
		return err
	}
	return nil
}
