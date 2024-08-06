// Copyright 2023-2024 Florian Zwoch <fzwoch@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tj3

import (
	"bytes"
	_ "embed"
	"image/color"
	"image/jpeg"
	"testing"
)

var (
	//go:embed testdata/gray.jpg
	gray []byte

	//go:embed testdata/rgb.jpg
	rgb []byte

	//go:embed testdata/yuv420.jpg
	yuv420 []byte

	//go:embed testdata/yuv422.jpg
	yuv422 []byte

	//go:embed testdata/yuv444.jpg
	yuv444 []byte

	//go:embed testdata/cmyk.jpg
	cmyk []byte
)

func TestDecodeConfigSmallerThanBuffer(t *testing.T) {
	_, err := DecodeConfig(bytes.NewReader(gray[:1024]))
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecodeConfigTooShort(t *testing.T) {
	_, err := DecodeConfig(bytes.NewReader(gray[:128]))
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestDecodeConfigGray(t *testing.T) {
	cfg, err := DecodeConfig(bytes.NewReader(gray))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.ColorModel != color.GrayModel {
		t.Fatal("unexpected color model")
	}
	if cfg.Width != 320 {
		t.Fatal("unexpected width")
	}
	if cfg.Height != 240 {
		t.Fatal("unexpected height")
	}
}

func TestDecodeConfigRGB(t *testing.T) {
	cfg, err := DecodeConfig(bytes.NewReader(rgb))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.ColorModel != color.NRGBAModel {
		t.Fatal("unexpected color model")
	}
	if cfg.Width != 320 {
		t.Fatal("unexpected width")
	}
	if cfg.Height != 240 {
		t.Fatal("unexpected height")
	}
}

func TestDecodeConfigYUV420(t *testing.T) {
	cfg, err := DecodeConfig(bytes.NewReader(yuv420))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.ColorModel != color.YCbCrModel {
		t.Fatal("unexpected color model")
	}
	if cfg.Width != 320 {
		t.Fatal("unexpected width")
	}
	if cfg.Height != 240 {
		t.Fatal("unexpected height")
	}
}

func TestDecodeConfigYUV422(t *testing.T) {
	cfg, err := DecodeConfig(bytes.NewReader(yuv422))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.ColorModel != color.YCbCrModel {
		t.Fatal("unexpected color model")
	}
	if cfg.Width != 320 {
		t.Fatal("unexpected width")
	}
	if cfg.Height != 240 {
		t.Fatal("unexpected height")
	}
}

func TestDecodeConfigYUV444(t *testing.T) {
	cfg, err := DecodeConfig(bytes.NewReader(yuv444))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.ColorModel != color.YCbCrModel {
		t.Fatal("unexpected color model")
	}
	if cfg.Width != 320 {
		t.Fatal("unexpected width")
	}
	if cfg.Height != 240 {
		t.Fatal("unexpected height")
	}
}

func TestDecodeConfigCMYK(t *testing.T) {
	cfg, err := DecodeConfig(bytes.NewReader(cmyk))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.ColorModel != color.CMYKModel {
		t.Fatal("unexpected color model")
	}
	if cfg.Width != 320 {
		t.Fatal("unexpected width")
	}
	if cfg.Height != 240 {
		t.Fatal("unexpected height")
	}
}

func TestDecodeGray(t *testing.T) {
	img, err := Decode(bytes.NewReader(gray))
	if err != nil {
		t.Fatal(err)
	}
	if img.Bounds().Dx() != 320 {
		t.Fatal("unexpected width")
	}
	if img.Bounds().Dy() != 240 {
		t.Fatal("unexpected height")
	}
}

func TestDecodeRGB(t *testing.T) {
	img, err := Decode(bytes.NewReader(rgb))
	if err != nil {
		t.Fatal(err)
	}
	if img.Bounds().Dx() != 320 {
		t.Fatal("unexpected width")
	}
	if img.Bounds().Dy() != 240 {
		t.Fatal("unexpected height")
	}
}

func TestDecodeYUV420(t *testing.T) {
	img, err := Decode(bytes.NewReader(yuv420))
	if err != nil {
		t.Fatal(err)
	}
	if img.Bounds().Dx() != 320 {
		t.Fatal("unexpected width")
	}
	if img.Bounds().Dy() != 240 {
		t.Fatal("unexpected height")
	}
}

func TestDecodeYUV422(t *testing.T) {
	img, err := Decode(bytes.NewReader(yuv422))
	if err != nil {
		t.Fatal(err)
	}
	if img.Bounds().Dx() != 320 {
		t.Fatal("unexpected width")
	}
	if img.Bounds().Dy() != 240 {
		t.Fatal("unexpected height")
	}
}

func TestDecodeYUV444(t *testing.T) {
	img, err := Decode(bytes.NewReader(yuv444))
	if err != nil {
		t.Fatal(err)
	}
	if img.Bounds().Dx() != 320 {
		t.Fatal("unexpected width")
	}
	if img.Bounds().Dy() != 240 {
		t.Fatal("unexpected height")
	}
}

func TestDecodeCMYK(t *testing.T) {
	img, err := Decode(bytes.NewReader(cmyk))
	if err != nil {
		t.Fatal(err)
	}
	if img.Bounds().Dx() != 320 {
		t.Fatal("unexpected width")
	}
	if img.Bounds().Dy() != 240 {
		t.Fatal("unexpected height")
	}
}

func TestDecodeInvalidData(t *testing.T) {
	_, err := Decode(bytes.NewReader(yuv420[1024:]))
	if err == nil {
		t.Fatal("expected error")
	}
}

func BenchmarkDecodeYUV420_TurboJPEG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode(bytes.NewReader(yuv420))
	}
}

func BenchmarkDecodeYUV420_GoJPEG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jpeg.Decode(bytes.NewReader(yuv420))
	}
}
