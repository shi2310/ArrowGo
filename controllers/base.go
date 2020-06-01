package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"path"

	"ArrowGo/models"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Upload 公用上传文件方法
func Upload(context *gin.Context, field string) (*models.File, error) {
	// 获取FormData中文件
	file, header, err := context.Request.FormFile(field)
	if err != nil {
		return nil, err
	}
	// 获取文件的MD5
	md5 := md5.New()
	io.Copy(md5, file)
	fileMd5 := hex.EncodeToString(md5.Sum(nil))
	v, err := models.GetFileByMD5(fileMd5)
	if err != nil {
		// 没有获取到文件，则上传文件
		folder := "./upload"
		_, err := os.Stat(folder)
		if err != nil {
			if os.IsNotExist(err) {
				os.Mkdir(folder, os.ModePerm)
			}
		}
		oldFileName := header.Filename
		ext := path.Ext(oldFileName)
		_uuid := uuid.NewV4()
		filename := _uuid.String() + ext
		filepath := path.Join(folder, filename)

		if err := context.SaveUploadedFile(header, filepath); err != nil {
			return nil, errors.New("文件保存出错")
		}

		v = &models.File{Name: filename, URL: filepath, FileMd5: fileMd5}
		if err := models.AddFile(v); err != nil {
			return nil, errors.New("submit file " + oldFileName + " record failed")
		}

	}
	return v, nil
}
