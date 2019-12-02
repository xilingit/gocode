package main

import (
	"fmt"
	"io/ioutil"
)

/*
	文件分为文本文件和二进制文件

	打开和关闭文件
		os.Open()函数能够打开一个文件，返回一个*File和一个err，对
		得到的文件示例调用close()方法能够关闭文件

	读取文件
		func (f *File)Read(b []byte) (n int,err error)
		接收一个字符切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回0和io.EOF
		可以用for循环来读取文件中的所有数据，先读到一个切片中，读完后添加到一个总的切片中，最后转string

		bufio读取文件
		在file基础上封装了一层API，支持更多功能

		ioutil读取整个文件
		io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入

	文件写入操作
		os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能
		func OpenFile(name string,flag int,perm FileMode)(*File,err){}
		name:要打开的文件名
		flag:打开文件的模式，只写，创建文件，只读，读写，清空，追加
		perm:文件权限，一个八进制数，r:04	w:02	x:01

		file.Write和WriteString

		bufio.NewWriter

		ioutil.WriteFile
*/
func main() {
	//打开的是该项目下的文件
	// f, err := os.Open("./oldboy/day1120file/main.go")
	// if err != nil {
	// 	fmt.Println("open file err,err:", err)
	// 	return
	// }
	// defer f.Close()
	// var content []byte
	// var temp = make([]byte, 128)
	// for {
	// 	n, err := f.Read(temp)
	// 	if err == io.EOF {
	// 		fmt.Println("文件读完了")
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("read file failed,err:", err)
	// 		return
	// 	}
	// 	content = append(content, temp[:n]...)
	// }
	// fmt.Println(string(content))

	// reader := bufio.NewReader(f)
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		fmt.Println("文件读完了")
	// 		return
	// 	}
	// 	if err != nil {
	// 		fmt.Println("read file failed,err:", err)
	// 		return
	// 	}
	// 	fmt.Println(line)
	// }

	content, err := ioutil.ReadFile("./oldboy/day1120file/main.go")
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}
	fmt.Println(string(content))
}
