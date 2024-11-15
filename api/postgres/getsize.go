package postgres

func GetNumMessages() (uint64, error) {
	var result []int
	err := Database.Select(&result, "SELECT max(id) from messages;")
	if err != nil {
		return 0, err
	}
	return uint64(result[0]), nil
}
