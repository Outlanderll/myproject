package underly

func SqlClose() {
	err := db.Close()
	if err != nil {
		return
	}
}
