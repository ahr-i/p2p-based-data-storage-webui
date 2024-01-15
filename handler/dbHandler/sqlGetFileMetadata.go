package dbHandler

func (sl *SqliteHandler) GetFileMetadata(id int) *FileMetadata {
	/* Query / path, hash */
	statement, err := sl.Database.Prepare("SELECT path, hash FROM file_metadata WHERE id = ?")
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	/* Query한 결과를 Metadata Block에 저장 */
	file_metadata := FileMetadata{}
	err_ := statement.QueryRow(id).Scan(&file_metadata.Path, &file_metadata.Hash)
	if err_ != nil {
		panic(err)
	}

	return &file_metadata
}
