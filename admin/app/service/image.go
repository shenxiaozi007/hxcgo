package service

import (
	"github.com/huangxinchun/hxcgo/admin/core/opt"
	"strings"

	"github.com/disintegration/imaging"
)

type Image struct {
}

func (img *Image) Dir() string {
	return strings.Trim(opt.Config().UploadDir, "/") + "/images"
}

func (img *Image) TmpDir() string {
	return img.Dir() + "/tmp"
}

func (img *Image) Thumbnail(srcFilename string, dstFilename string, width int, height int, filter imaging.ResampleFilter) error {
	srcImage, err := imaging.Open(srcFilename)
	if err != nil {
		return err
	}
	dstImage := imaging.Thumbnail(srcImage, width, height, filter)
	return imaging.Save(dstImage, dstFilename)
}
