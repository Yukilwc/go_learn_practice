package fileoperate

import "testing"

func TestCopyFile(t *testing.T) {
	src := "D:/workspace/libiary/ForTest/go_code/go_test_init/fileoperate/test/src/srcData.js"
	dst := "D:/workspace/libiary/ForTest/go_code/go_test_init/fileoperate/test/dst/srcData.js"
	if err := CopyFile(src, dst); err != nil {
		t.Fatalf("copy failed error:%s", err)
	}
}
