package fileoperate

import "testing"

func TestCopyFile(t *testing.T) {
	src := "D:/workspace/libiary/ForTest/go_code/go_test_init/fileoperate/test/src/srcData.js"
	dst := "D:/workspace/libiary/ForTest/go_code/go_test_init/fileoperate/test/dst/srcData.js"
	if err := CopyFile(src, dst); err != nil {
		t.Fatalf("copy file failed error:%s", err)
	}
}

func TestCopyDir(t *testing.T) {
	src := "D:/workspace/libiary/ForTest/go_code/go_test_init/fileoperate/test/srcFolder"
	dst := "D:/workspace/libiary/ForTest/go_code/go_test_init/fileoperate/test/dstFolder"
	if err := CopyDir(src, dst); err != nil {
		t.Fatalf("copy dir failed error:%s:", err)
	}
}
