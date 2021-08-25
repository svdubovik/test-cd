package main

import "testing"

func TestHello(t *testing.T) {
	got := buildSQL("test@test.com")
	want := "SELECT * FROM users WHERE email='test@test.com' AND Password='123456';"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
