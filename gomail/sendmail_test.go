package gomail

import "testing"

func TestSendMail(t *testing.T) {
	if err := SendMail(); err != nil{
		t.Fatalf("SendMail failed. error is %s", err)
	}
}
