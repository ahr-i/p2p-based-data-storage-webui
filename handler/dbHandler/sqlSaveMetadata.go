package dbHandler

func (wh *SqliteHandler) SaveFileMetadata(file_metadata FileMetadata) int {
	/* Insert / File의 Metadata 저장 */
	statement, err := wh.Database.Prepare("INSERT INTO file_metadata (name, path, size, hash) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}

	/* Insert의 결과 */
	result, err := statement.Exec(file_metadata.Name, file_metadata.Path, file_metadata.Size, file_metadata.Hash)
	if err != nil {
		panic(err)
	}

	id, _ := result.LastInsertId() // 저장한 File의 ID값

	return int(id)
}
