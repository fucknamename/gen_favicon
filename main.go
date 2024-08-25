package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"

	ico "github.com/Kodeworks/golang-image-ico"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

const width, height = 32, 32

func main() {
	var text string
	flag.StringVar(&text, "t", "T", "ico text.")
	flag.Parse()

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 设置背景颜色为白色
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// 加载字体文件(如果要嵌入字体, 参考另一个项目: https://github.com/fucknamename/gen_logo)
	fontBytes, err := os.ReadFile("./ZiHunDaHei.ttf")
	if err != nil {
		log.Fatalf("failed to read font file: %v", err)
	}

	f, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}

	// 创建字体面并设置大小
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create font face: %v", err)
	}

	// 设置字体绘制器
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black), // 设置字体颜色为黑色
		Face: face,
		Dot:  fixed.P(8, 24),
	}

	// 绘制文字
	d.DrawString(text)

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
	icoFile, err := os.Create("favicon.ico")
	if err != nil {
		log.Fatal(err)
	}
	defer icoFile.Close()

	if err := ico.Encode(icoFile, img); err != nil {
		log.Fatal(err)
	}
}
