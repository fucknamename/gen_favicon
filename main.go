package main

import (
	_ "embed"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"
	"strconv"
	"unicode"

	ico "github.com/Kodeworks/golang-image-ico"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

//go:embed font/ZiHunDaHei.ttf
var fontBytes []byte

const (
	width  = 32 // 宽
	height = 32 // 高
)

var (
	x, y          = 7, 23
	size  float64 = 30 // 字体默认大小
	text          = flag.String("t", "T", "favicon 图标文本, 最多3个字符")
	bclor         = flag.String("b", "ff9900", "背景颜色: ff9900 / 1f1f1f")
	fclor         = flag.String("f", "000000", "字体颜色: 000000 / ffffff")
)

func main() {
	flag.Parse()

	t := []rune(*text)
	count := len(t)
	if count == 1 {
		size = 30
		if isChinese(t[0]) {
			x, y = 2, 25
		} else {
			x, y = 5, 23
		}
	}
	if count == 2 {
		if isChinese(t[0]) || isChinese(t[1]) {
			size = 16
			x, y = 2, 21
		} else {
			size = 20
			x, y = 3, 21
		}
	}
	if count == 3 {
		if isChinese(t[0]) || isChinese(t[1]) || isChinese(t[2]) {
			size = 12
			x, y = 1, 20
		} else {
			size = 14
			x, y = 1, 20
		}
	}
	if count < 1 || count > 3 {
		fmt.Println("字符长度最多3位")
		return
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	bgcolor, err := hexToColor(*bclor)
	if err != nil {
		fmt.Println("背景色代码无效")
		return
	}

	ftcolor, err := hexToColor(*fclor)
	if err != nil {
		fmt.Println("字体颜色代码无效")
		return
	}

	// 设置背景颜色
	draw.Draw(img, img.Bounds(), &image.Uniform{bgcolor}, image.Point{}, draw.Src)

	// // 加载字体文件(如果要嵌入字体, 参考另一个项目: https://github.com/fucknamename/gen_logo)
	// fontBytes, err := os.ReadFile("./ZiHunDaHei.ttf")
	// if err != nil {
	// 	fmt.Println("读取字体失败")
	// 	return
	// }

	f, err := opentype.Parse(fontBytes)
	if err != nil {
		fmt.Println("字体解析失败")
		return
	}

	// 创建字体面并设置大小
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		fmt.Println("实例化字体对象失败")
		return
	}

	// 设置字体绘制器
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(ftcolor), // 设置字体颜色
		Face: face,
		Dot:  fixed.P(x, y),
	}

	// 绘制文字
	d.DrawString(*text)

	// // 将图像保存为 PNG 文件（用于调试）
	// file, err := os.Create("output.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// if err := png.Encode(file, img); err != nil {
	// 	log.Fatal(err)
	// }

	// 创建并写入 ICO 文件
	icoFile, _ := os.Create("favicon.ico")
	defer icoFile.Close()

	if err := ico.Encode(icoFile, img); err != nil {
		fmt.Println("生成图标失败")
	} else {
		fmt.Println("success")
	}
}

// 颜色代码转化
func hexToColor(hex string) (color.Color, error) {
	if len(hex) != 6 {
		return nil, fmt.Errorf("invalid length, expected 6 characters")
	}

	// 解析红色
	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return nil, err
	}

	// 解析绿色
	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return nil, err
	}

	// 解析蓝色
	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return nil, err
	}

	return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}, nil
}

// 判断是否为汉字
func isChinese(char rune) bool {
	return unicode.Is(unicode.Han, char)
}

// // 判断是否为字母（包括大写和小写）
// func isLetter(char rune) bool {
// 	return unicode.IsLetter(char)
// }
