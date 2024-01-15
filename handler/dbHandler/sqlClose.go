package dbHandler

func (sl *SqliteHandler) Close() {
	sl.Database.Close()
}
