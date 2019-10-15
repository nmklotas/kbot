package command

import "testing"

func TestRecognizesBotCommand(t *testing.T) {
	result := IsBotCommand("kbot list")
	if !result {
		t.Errorf("kbot command not recognized")
	}
}
