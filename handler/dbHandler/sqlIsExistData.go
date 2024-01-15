package dbHandler

func (sl *SqliteHandler) IsExistData(hash string) bool {
	/* Query / hash Data의 존재 여부 */
	statement, err := sl.Database.Prepare("SELECT EXISTS(SELECT 1 FROM file_metadata WHERE hash = ?)")
	if err != nil {
		panic(err)
	}

	/* Query한 결과를 저장 */
	var result bool
	err_ := statement.QueryRow(hash).Scan(&result)
	if err_ != nil {
		panic(err)
	}

	return result
}
