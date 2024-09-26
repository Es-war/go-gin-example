package file

import (
	"io"
	"mime/multipart"
	"os"
	"path"
)

// 获取文件大小（以字节为单位）
func GetSize(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)

	return len(content), err
}

// 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// 检查文件是否存在
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	// 如果 err 为 nil，表示文件或目录存在；
	// 如果 err 不为 nil，则可以通过 os.IsNotExist(err) 来检查该错误是否表示文件不存在
	return os.IsNotExist(err)
}

// 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	// 如果路径存在且有权限，返回 false（即没有权限错误）。
	// 如果路径不存在，os.Stat 将返回一个错误，
	//     该错误不属于权限错误，因此 os.IsPermission(err) 返回 false
	return os.IsPermission(err)
}

// 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// 打开指定名称的文件，使用指定的标志和权限
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}