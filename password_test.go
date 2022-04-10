package _crypto

import "testing"

func TestHashPassword(t *testing.T) {
	pass := "123456"
	r, err := HashPassword("123456")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("hashed password:", r)
	err = VerifyPassword(r, pass+"2")
	if err != nil {
		t.Error(err)
		return
	} else {
		t.Log("密码校验通过")
	}
}
