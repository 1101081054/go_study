package main

import (
	"fmt"
	"errors"
	"os"
)

/**日志对外接口*/

//声明日志写入器接口
type LogWriter interface {
	Write(data interface{}) error
}

//日志器
type Logger struct {
	//这个日志器用到的日志写入器
	writeList []LogWriter
}

//注册一个日志写入器
func (l *Logger) RegisterWrite(writer LogWriter) {
	l.writeList = append(l.writeList, writer)
	fmt.Println("注册", writer)
}

//将一个data类型的数据写入日志
func (l *Logger) Log(data interface{}) {
	for _, writer := range l.writeList {
		//将日志输出到每一个写入器中
		writer.Write(data)
	}
}

//创建日志器的实例
func NewLogger() *Logger {
	return &Logger{}
}




/**文件写入器*/

//声明文件写入器
type fileWriter struct {
	file *os.File
}

//设置文件写入器写入的文件名
func (f *fileWriter) SetFile(filename string) (err error) {
	//如果文件已经打开，关闭前一个文件
	if f.file != nil {
		f.file.Close()
	}

	//创建一个文件并保存文件句柄
	f.file, err = os.Create(filename)

	//如果创建的过程出现错误，则返回错误
	return err
}

//实现LogWriter的Write()方法
func (f *fileWriter) Write(data interface{}) error {
	//日志文件可能没有创建成功
	if f.file == nil {
		//日志文件没有准备好
		return errors.New("file not created")
	}

	//将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)

	//将数据以字节数组写入文件中
	_, err := f.file.Write([]byte(str))

	return err
}

//创建文件写入器实例
func newFileWriter() *fileWriter {
	return &fileWriter{}
}





/**命令行写入器*/

//命令行写入器
type consoleWriter struct {

}

//实现LogWriter的Write() 方法
func (f *consoleWriter) Write(data interface{}) error {
	//将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)

	//将数据以字节数组写入命令行中
	_, err := os.Stdout.Write([]byte(str))

	return err
}

//创建命令行写入器实例
func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}






/**使用日志*/

//创建日志器
func createLogger() *Logger {
	//创建日志器
	l := NewLogger()

	//创建命令行写入器
	cw := newConsoleWriter()

	//注册命令行写入器到日志器中
	l.RegisterWrite(cw)

	//创建文件写入器
	fw := newFileWriter()

	//设置文件名
	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}

	//注册文件写入器到日志器
	l.RegisterWrite(fw)

	return l
}

func main(){
	//准备一个日志器
	l := createLogger()

	//写一个日志
	l.Log("hello")
}