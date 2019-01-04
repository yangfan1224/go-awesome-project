package image

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/pkg/errors"
	"image"
	"image/jpeg"
	"math"
	"os"
)

const DEFAULT_MAX_WIDTH float64 = 54

const DEFAULT_MAX_HEIGHT float64 = 66

// 计算图片缩放后的尺寸
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// 生成缩略图
func MakeThumbnail(imagePath, savePath string) (string, error) {

	file, _ := os.Open(imagePath)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", errors.Wrap(err,"image.Decode error")
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := calculateRatioFit(width, height)

	fmt.Println("width = ", width, " height = ", height)
	fmt.Println("w = ", w, " h = ", h)

	// 调用resize库进行图片缩放
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	// 需要保存的文件
	imgfile, _ := os.Create(savePath)
	defer imgfile.Close()

	// 以PNG格式保存文件
	err = jpeg.Encode(imgfile, m, &jpeg.Options{Quality:100})
	if err != nil {
		errors.Wrap(err,"png.Encode error")
	}

	return "", nil
}
