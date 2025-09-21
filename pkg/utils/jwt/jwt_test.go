package jwt

import "testing"

func TestGenAndParseToken(t *testing.T) {
	token, err := GenerateToken("testAdmin@qq.com")
	if err != nil {
		t.Fatal(err)
	}
	claims, err := ParseToken(token)
	if err != nil {
		t.Fatal(err)
	}

	if claims.Username != "testAdmin@qq.com" {
		t.Fatalf("expect username testAdmin@qq.com, got %s", claims.Username)
	}
}