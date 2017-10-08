package handler_backend

import (
	"github.com/gin-gonic/gin"
	"path"
	"os"
	dm "git.vdo.space/foy/model"
	ml "git.vdo.space/berk/model"
	"github.com/satori/go.uuid"
	"errors"
	. "git.vdo.space/berk/consts"
	"io"
)

func FileUpload(c *gin.Context, pathstr string) *dm.ResultModel {

	filename := ""
	ext := ""
	rm := dm.ResultModel{}

	d := make(dm.Dungeons)

	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		panic(err)
	}

	if header != nil {

		filename = uuid.NewV4().String()

		ext = path.Ext(header.Filename)

		if !ml.ExtLimit(ext, ImageFileExts) {

			rm.Errors = append(rm.Errors, dm.NewError(errors.New("This file extra must be with .png .jpg .jpeg ! pls try again! ")))
			return &rm

		}

		out, err := os.OpenFile(os.Getenv(ResourcesPath) + pathstr + filename + ext, os.O_RDWR | os.O_CREATE, 0644) //打开文件

		if err != nil {
			rm.Errors = append(rm.Errors, dm.NewError(err))
			return &rm
		}

		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			rm.Errors = append(rm.Errors, dm.NewError(err))
			return &rm
		}
	}

	rm.Data = d
	d["file_name"] = filename + ext
	d["file_url"] = os.Getenv(ResourcesAddr) + pathstr + filename + ext
	return &rm
}
