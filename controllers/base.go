package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"ArrowGo/models"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// FormUpload Form上传文件方法
func FormUpload(context *gin.Context, field string) (*models.File, error) {
	// 获取FormData中文件
	header, err := context.FormFile(field)
	if err != nil {
		return nil, err
	}
	oldFileName := header.Filename
	ext := path.Ext(oldFileName)

	file, err := header.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	// if err := context.SaveUploadedFile(header, filepath); err != nil {
	// 	return nil, errors.New("文件保存出错")
	// }
	return saveFile(ext, bytes)
}

// Base64Upload Base64上传文件方法
func Base64Upload(context *gin.Context) (*models.File, error) {
	bytes, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		return nil, err
	}
	strs := strings.Split(string(bytes), ",")
	head := strs[0]
	body := strs[1]
	start := strings.LastIndex(head, "/")
	end := strings.LastIndex(head, ";")
	ext := head[start+1 : end]

	return saveFile(ext, []byte(body))
}

// saveFile 私有方法
func saveFile(ext string, bytes []byte) (*models.File, error) {
	// 获取文件MD5
	md5 := md5.New()
	md5.Write(bytes)
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

		// 重设文件名
		_uuid := uuid.NewV4()
		filename := _uuid.String() + ext
		filepath := path.Join(folder, filename)

		if err = ioutil.WriteFile(filepath, bytes, 0666); err != nil {
			return nil, errors.New("文件保存出错")
		}

		v = &models.File{Name: filename, URL: filepath, FileMd5: fileMd5}
		if err := models.AddFile(v); err != nil {
			return nil, errors.New("submit file record failed")
		}
	}
	return v, nil
}
