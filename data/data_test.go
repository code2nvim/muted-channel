package data

import (
	"testing"
)

func TestData(t *testing.T) {
	data := Database{Conn("../.env")}
	defer data.DB.Close()
}
