package dbHandler

import "time"

type DBHandler interface {
	GetFileList() []*FileMetadata
	SaveFileMetadata(file_metadata FileMetadata) int
	GetFileMetadata(id int) *FileMetadata
	IsExistData(hash string) bool
	GetStorageInfo() int
	Close()
}

type FileMetadata struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Path     string    `json:"path"`
	Size     int       `json:"size"`
	Hash     string    `json:"hash"`
	CreateAt time.Time `json:"create_at"`
}

func CreateDBHandler(db_path string) DBHandler {
	return CreateSqliteHandler(db_path)
}
