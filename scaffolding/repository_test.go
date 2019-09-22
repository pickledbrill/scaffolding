package scaffolding

import "testing"

func TestSetupRepository_ShouldCreateFileAndDirectory(t *testing.T) {
	// setup
	dir := "C:\\Example"
	files := make([]string, 3)
	files[0] = "test.txt"
	files[1] = "test01/"
	files[2] = "test02/"

	result, err := SetupRepository(dir, files)
	// validate
	if err != nil {
		t.Error(err)
	}
	if !result {
		t.Error("Something is wrong")
	}
}
