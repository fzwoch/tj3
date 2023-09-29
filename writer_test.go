// Copyright 2023 Florian Zwoch <fzwoch@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tj3

import (
	"bytes"
	"image"
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

func TestEncodeCMYK(t *testing.T) {
	m := image.NewCMYK(r)
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

func TestEncodeYuv420(t *testing.T) {
	m := image.NewYCbCr(r, image.YCbCrSubsampleRatio420)
	w := new(bytes.Buffer)
	err := Encode(w, m, &Options{Quality: DefaultQuality})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncodeYuv422(t *testing.T) {
	m := image.NewYCbCr(r, image.YCbCrSubsampleRatio422)
	w := new(bytes.Buffer)
	err := Encode(w, m, &Options{Quality: DefaultQuality})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncodeYuv444(t *testing.T) {
	m := image.NewYCbCr(r, image.YCbCrSubsampleRatio444)
	w := new(bytes.Buffer)
	err := Encode(w, m, &Options{Quality: DefaultQuality})
	if err != nil {
		t.Fatal(err)
	}
}
