package Singleton

import "testing"

func TestGetInstance(t *testing.T) {
	single := GetInstance()
	t.Logf("Print single instance: %s", single.Show())
}
