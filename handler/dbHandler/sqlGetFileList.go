package dbHandler

func (sl *SqliteHandler) GetFileList() []*FileMetadata {
	file_list := []*FileMetadata{}

	/* Query / id, name, size createAt */
	row, err := sl.Database.Query("SELECT id, name, size, createdAt FROM file_metadata")
	if err != nil {
		panic(err)
	}
	defer row.Close()

	/* Query한 결과를 List에 저장 */
	for row.Next() {
		var file FileMetadata

		row.Scan(&file.ID, &file.Name, &file.Size, &file.CreateAt)
		file_list = append(file_list, &file)
	}

	return file_list
}
