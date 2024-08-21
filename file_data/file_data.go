package filedata

import "htmx-learning/repository"

type FileData interface {
	CheckNewFileRealTime()
}

type fileData struct {
	dataTempRepo repository.Repository
}

func NewFileData(dataTempRepo repository.Repository) FileData {
	return fileData{dataTempRepo}
}
