package services

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

const uploadDir = "uploads"
const tempDir = "temp"

type FileSystemService interface {
	SaveToTemp(file multipart.File, filename string) (string, error)
	CleanTemp()
	MoveToPermanent(filenames []string, folder string) error
	Delete(filenames []string, folder string) error
	getPermanentPath(folder string) string
	getNewFileName(filename string) (string, error)
}

type fileSystemService struct{}

func NewFileSystemService() FileSystemService {
	return &fileSystemService{}
}

func (s *fileSystemService) SaveToTemp(file multipart.File, filename string) (string, error) {
	dir := filepath.Join(uploadDir, tempDir)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Printf("failed to create temp directory: %v", err)
		return "", err
	}

	newFilename, err := s.getNewFileName(filename)
	if err != nil {
		log.Printf("failed to generate new filename: %v", err)
		return "", err
	}

	tempFilePath := filepath.Join(dir, newFilename)
	outFile, err := os.Create(tempFilePath)
	if err != nil {
		log.Printf("failed to create temporary file: %v", err)
		return "", err
	}

	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			log.Printf("failed to close temporary file: %v", err)
		}
	}(outFile)

	if _, err := io.Copy(outFile, file); err != nil {
		log.Printf("failed to write file contents: %v", err)
		return "", err
	}

	return newFilename, nil
}

func (s *fileSystemService) CleanTemp() {
	tempDirPath := filepath.Join(uploadDir, tempDir)
	err := os.RemoveAll(tempDirPath)
	if err != nil {
		log.Printf("failed to clean temporary upload directory: %v", err)
	}
}

func (s *fileSystemService) MoveToPermanent(filenames []string, folder string) error {
	permanentDir := s.getPermanentPath(folder)
	err := os.MkdirAll(permanentDir, os.ModePerm)
	if err != nil {
		log.Printf("failed to create permanent upload directory: %v", err)
		return err
	}

	for _, filename := range filenames {
		tempFilePath := filepath.Join(uploadDir, tempDir, filename)
		permanentFilePath := filepath.Join(permanentDir, filename)

		if _, err := os.Stat(tempFilePath); os.IsNotExist(err) {
			log.Printf("temporary file does not exist: %s", tempFilePath)
			continue
		}

		if err := os.Rename(tempFilePath, permanentFilePath); err != nil {
			log.Printf("failed to move file from %s to %s: %v", tempFilePath, permanentFilePath, err)
			continue
		}
	}

	return nil
}

func (s *fileSystemService) Delete(filename []string, folder string) error {
	permanentDir := s.getPermanentPath(folder)

	for _, file := range filename {
		filePath := filepath.Join(permanentDir, file)
		if err := os.Remove(filePath); err != nil {
			log.Printf("failed to delete file %s: %v", filePath, err)
			return err
		}
	}

	return nil
}

func (s *fileSystemService) getPermanentPath(folder string) string {
	return filepath.Join(uploadDir, folder)
}

func (s *fileSystemService) getNewFileName(filename string) (string, error) {
	randomBytes := make([]byte, 8)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	randomPrefix := hex.EncodeToString(randomBytes)
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	newName := name + "_" + randomPrefix + ext

	return newName, nil
}
