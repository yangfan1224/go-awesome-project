package bbolt

import "testing"

func TestImportBbolt(t *testing.T) {
	ImportBbolt()
}

func TestUpdateBolt(t *testing.T) {
	err := UpdateBolt()
	if err != nil{
		t.Fatalf("GetBolt failed. err is %s " , err)
	}
}

func TestGetBolt(t *testing.T) {
	err := GetBolt()
	if err != nil{
		t.Fatalf("GetBolt failed. err is %s " , err)
	}
}
