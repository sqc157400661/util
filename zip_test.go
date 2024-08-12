package helper

import (
	"os"
	"testing"
)

func TestZip(t *testing.T) {
	// 创建空文件 example.txt
	//f1, err := os.Create("./example1.txt")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//// 创建空文件 example2.txt
	//f2, err := os.Create("./example2.txt")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//// 创建空文件 example2.txt
	//f3, err := os.Create("./example3.txt")
	//if err != nil {
	//	t.Fatal(err)
	//}
	var err error
	f1, _ := os.Open("./test")
	var files = []*os.File{f1}
	dest := "test.zip"
	err = Zip(files, dest)
	if err != nil {
		t.Fatal(err)
	}
}
func TestUnzip(t *testing.T) {
	err := Unzip("test.zip", "./test/")
	if err != nil {
		t.Fatal(err)
	}
}
