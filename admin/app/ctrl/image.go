package ctrl

import (
	"github.com/huangxinchun/hxcgo/admin/app/service"
	"github.com/huangxinchun/hxcgo/admin/core/opt"
	"github.com/huangxinchun/hxcgo/admin/core/uuid"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"

	e "github.com/huangxinchun/hxcgo/admin/app/err"

	"github.com/gin-gonic/gin"
)

type Image struct {
	mimeTypes []string
	maxSize   int64
	idFactory *uuid.IDFactory
}

func NewImage() *Image {
	idFactory, err := uuid.New(int64(opt.Config().Node))
	if err != nil {
		panic(fmt.Sprintf("image err: %s", err.Error()))
	}
	img := &Image{
		mimeTypes: []string{
			"image/gif",
			"image/jpeg",
			"image/jpg",
			"image/png",
		},
		maxSize:   1024 * 1024 * 10, //10M
		idFactory: idFactory,
	}

	go img.gc()
	return img
}

func (img *Image) gc() {
	ticker := time.NewTicker(7 * 24 * time.Hour)
	for {
		select {
		case <-ticker.C:
			imageService := &service.Image{}
			files, err := ioutil.ReadDir(imageService.TmpDir())
			if err != nil {
				log.Println("read dir err: ", err)
				continue
			}

			yesterday := time.Now().AddDate(0, 0, -1).Format("20060102")
			for _, file := range files {
				if strings.Compare(yesterday, file.Name()) > 0 {
					os.RemoveAll(fmt.Sprintf("%s/%s", imageService.TmpDir(), file.Name()))
				}
			}
		}
	}
}

func (img *Image) setHeader(c *gin.Context) {
	c.Request.Header.Add("Access-Control-Allow-Origin", "*")
	c.Request.Header.Add(
		"Access-Control-Allow-Methods",
		"OPTIONS, HEAD, GET, POST, DELETE",
	)
	c.Request.Header.Add(
		"Access-Control-Allow-Headers",
		"Content-Type, Content-Range, Content-Disposition",
	)
}

func (img *Image) validateSize(size int64) error {
	if size == 0 {
		return e.ERequest
	}
	if size > img.maxSize {
		return e.ELimitSize
	}
	return nil
}

func (img *Image) validateMIMEType(mimeType string) error {
	isAllow := false
	for _, mime := range img.mimeTypes {
		if mimeType == mime {
			isAllow = true
			break
		}
	}
	if !isAllow {
		return e.EMIMEType
	}
	return nil
}

func (img *Image) createImageFile(filename string, buf []byte) error {
	fd, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		return err
	}
	_, err = fd.Write(buf)
	fd.Close()
	return err
}

func (img *Image) Empty(c *gin.Context) {
	img.setHeader(c)
}

func (img *Image) parseParam(c *gin.Context) (*multipart.FileHeader, int, int, error) {
	//获取宽高 size=120x120
	sizes := strings.Split(c.Query("size"), "x")
	if len(sizes) != 2 {
		return nil, 0, 0, e.EParam
	}
	width, err := strconv.Atoi(sizes[0])
	if err != nil || width <= 0 {
		return nil, 0, 0, e.EParam
	}
	height, err := strconv.Atoi(sizes[1])
	if err != nil || height <= 0 {
		return nil, 0, 0, e.EParam
	}

	//解析上传文件
	fh, err := c.FormFile("imageFile")
	if err != nil {
		return nil, 0, 0, e.EUpload
	}

	if img.validateSize(fh.Size) != nil {
		return nil, 0, 0, e.ELimitSize
	}

	return fh, width, height, nil
}

func (img *Image) Upload(c *gin.Context) {
	//设置头部信息
	img.setHeader(c)
	fh, width, height, err := img.parseParam(c)
	if err != nil {
		Json(c, err, nil)
		return
	}

	imageFile, err := fh.Open()
	if err != nil {
		Json(c, e.EUpload, nil)
		return
	}
	defer imageFile.Close()

	buf, err := ioutil.ReadAll(imageFile)
	if err != nil {
		Json(c, e.EUpload, nil)
		return
	}

	//验证 mime type
	fileType := http.DetectContentType(buf)
	if img.validateMIMEType(fileType) != nil {
		Json(c, e.EMIMEType, nil)
		return
	}
	//验证扩展名
	_, err = mime.ExtensionsByType(fileType)
	if err != nil {
		Json(c, e.EMIMEType, nil)
		return
	}

	imageService := &service.Image{}

	//临时上传目录
	uploadDir := fmt.Sprintf("%s/%s", imageService.TmpDir(), time.Now().Format("20060102"))
	err = os.MkdirAll(uploadDir, 0660)
	if err != nil {
		Json(c, e.EUpload, nil)
		return
	}

	filename := img.idFactory.String()
	ext := filepath.Ext(fh.Filename)
	//保存原图
	srcImageName := fmt.Sprintf("%s/%s%s", uploadDir, filename, ext)
	err = img.createImageFile(srcImageName, buf)

	//保存预览缩略图
	previewImageName := fmt.Sprintf("%s/%s_%dx%d%s", uploadDir, filename, width, height, ext)
	err = imageService.Thumbnail(srcImageName, previewImageName, width, height, imaging.Lanczos)
	if err != nil {
		log.Println("save preview image err: ", err)
		Json(c, e.EUpload, nil)
		return
	}

	Json(c, nil, gin.H{
		"src":   srcImageName,
		"thumb": "/" + previewImageName,
	})

}
