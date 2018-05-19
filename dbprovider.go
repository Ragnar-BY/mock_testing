package main

// DBProvider is for DB.
type DBProvider struct {
	DB Database
}

// AddValue adds key:value to db.
func (dp *DBProvider) AddValue(key string, value string) error {
	return dp.DB.Write(key, value)
}

// ReadValue read value by the key.
func (dp *DBProvider) ReadValue(key string) (string, error) {
	return dp.DB.Read(key)
}
