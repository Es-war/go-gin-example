package upload

import (
	"os"
	"path"
	"log"
	"fmt"
	"strings"
	"mime/multipart"

	"github.com/Es-war/go-gin-example/pkg/file"
	"github.com/Es-war/go-gin-example/pkg/setting"
	"github.com/Es-war/go-gin-example/pkg/logging"
	"github.com/Es-war/go-gin-example/pkg/util"
)

// 拼接图片的完整 URL，包含前缀和存储路径
// 例如： http://127.0.0.1:8000/  upload/images/  d41d8cd98f00b204e9800998ecf8427e.jpg
func GetImageFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

// 通过 MD5 对图片文件名进行加密，保留文件扩展名
// 例如： d41d8cd98f00b204e9800998ecf8427e.jpg
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// 返回图片的相对保存路径
// 例如： upload/images/
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

// 返回图片的完整保存路径，
// 结合运行时的根路径和图片的相对路径。
// 例如： runtime/ upload/images/
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// 检查图片文件的扩展名是否合法
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		// strings.EqualFold 专门用于不区分大小写的比较
		if strings.EqualFold(allowExt, ext) {
			return true
		}
	}

	return false
}

// 检查图片的大小是否在允许的范围内
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

// 检查图片保存的文件夹是否存在并且具有适当的权限
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}