package upload

import (
	"be-hoatieu/pkg/file"
	"be-hoatieu/pkg/setting"
	"be-hoatieu/pkg/utils"
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// GetImageFullUrl get the full access path
func GetImageFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}
func CheckFileExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.FileAllowExts {
		if strings.Contains(strings.ToUpper(ext), strings.ToUpper(allowExt)) {
			return true
		}
	}
	return false
}

// CheckFileSize kiểm tra kích thước file
func CheckFileSize(f *multipart.FileHeader) bool {
	return int(f.Size) <= setting.AppSetting.FileMaxSize
}

// GetImageName get image name
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)

	return fileName + ext
}

// GetImagePath get save path
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

// GetImageFullPath get full save path
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	strArrayImg := [3]string{".jpg", ".jpeg", ".png"}
	ext := file.GetExt(fileName)
	for _, allowExt := range strArrayImg {
		if strings.Contains(strings.ToUpper(ext), strings.ToUpper(allowExt)) {
			return true
		}
	}

	return false
}

// CheckImageSize check image size
func CheckImageSize(f *multipart.FileHeader) bool {
	return int(f.Size) <= setting.AppSetting.ImageMaxSize
}

// CheckImage check if the file exists
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
