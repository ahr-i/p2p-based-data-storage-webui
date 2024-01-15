package dbHandler

func (sl *SqliteHandler) GetStorageInfo() int {
	var storage_size int = 0

	/* Query / size */
	row, err := sl.Database.Query("SELECT size FROM file_metadata")
	if err != nil {
		panic(err)
	}
	defer row.Close()

	/* QUery한 결과를 저장 */
	for row.Next() {
		var size int

		row.Scan(&size)
		storage_size += size
	}

	return storage_size
}
