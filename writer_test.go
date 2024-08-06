// Copyright 2023-2024 Florian Zwoch <fzwoch@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tj3

import (
	"bytes"
	"image"
	"image/jpeg"
	"testing"
)

var r = image.Rectangle{
	Min: image.Point{0, 0},
	Max: image.Point{320, 240},
}

func TestEncodeNilOptions(t *testing.T) {
	m := image.NewGray(r)
	w := new(bytes.Buffer)
	err := Encode(w, m, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncodeInvalidQualityLow(t *testing.T) {
	m := image.NewGray(r)
	w := new(bytes.Buffer)
	err := Encode(w, m, &Options{Quality: 0})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEncodeInvalidQualityHigh(t *testing.T) {
	m := image.NewGray(r)
	w := new(bytes.Buffer)
	err := Encode(w, m, &Options{Quality: 101})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEncodeGray(t *testing.T) {
	m := image.NewGray(r)
	w := new(bytes.Buffer)
	err := Encode(w, m, &Options{Quality: DefaultQuality})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncodeRGB(t *testing.T) {
	m := image.NewNRGBA(r)
	w := new(bytes.Buffer)
	err := Encode(w, m, &Options{Quality: DefaultQuality})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncodeYUV420(t *testing.T) {
	m := image.NewYCbCr(r, image.YCbCrSubsampleRatio420)
	w := new(bytes.Buffer)
	err := Encode(w, m, &Options{Quality: DefaultQuality})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncodeYUV422(t *testing.T) {
	m := image.NewYCbCr(r, image.YCbCrSubsampleRatio422)
	w := new(bytes.Buffer)
	err := Encode(w, m, &Options{Quality: DefaultQuality})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncodeYUV444(t *testing.T) {
	m := image.NewYCbCr(r, image.YCbCrSubsampleRatio444)
	w := new(bytes.Buffer)
	err := Encode(w, m, &Options{Quality: DefaultQuality})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncodeCMYK(t *testing.T) {
	m := image.NewCMYK(r)
	w := new(bytes.Buffer)
	err := Encode(w, m, &Options{Quality: DefaultQuality})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncodeInvalidFormat(t *testing.T) {
	m := image.NewAlpha(r)
	w := new(bytes.Buffer)
	err := Encode(w, m, &Options{Quality: DefaultQuality})
	if err == nil {
		t.Fatal("expected error")
	}
}

func BenchmarkEncodeYUV420_TurboJPEG(b *testing.B) {
	m := image.NewYCbCr(r, image.YCbCrSubsampleRatio420)
	w := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		Encode(w, m, &Options{Quality: DefaultQuality})
	}
}

func BenchmarkEncodeYUV420_GoJPEG(b *testing.B) {
	m := image.NewYCbCr(r, image.YCbCrSubsampleRatio420)
	w := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		jpeg.Encode(w, m, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}
}
