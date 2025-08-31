package data

import (
	"testing"
)

func TestData(t *testing.T) {
	data := Data{Conn("../.env")}
	defer data.DB.Close()
}
