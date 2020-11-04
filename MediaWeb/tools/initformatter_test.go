package tools

import (
	"MediaWeb/tools/initformatter"
	"MediaWeb/tools/uuid"
	"testing"
)

func TestInitFormatter(t *testing.T)  {
	f, err := initformatter.New("../dbopts/dbconfig.ini")
	if err != nil {
		println(err)
	} else {
		f.Print()
	}
}

func TestGetuuid(t *testing.T)  {
	uuid.GenerateUUID()
}