package files_test

import (
	"testing"

	"github.com/VladiTNT/gitfetch/pkg/files"
)

func TestGetFileExtension(t *testing.T) {
	testCases := []struct {
		Input  string
		Output string
	}{
		{"main.go", "go"},
		{"page.toolbar.html", "html"},
		{"help.txt", "txt"},
		{"car.nerd.fuck.shit.jpg", "jpg"},
		{".git", "git"},
		{"./nerd/fuckface/whyareyoureadingthis/bob.png", "png"},
	}

	for i, testCase := range testCases {
		result := files.GetFileExtension(testCase.Input)
		if result != testCase.Output {
			t.Logf("Test case %d failed: wanted '%s', got '%s'\n", i, testCase.Output, result)
		}
	}
}
