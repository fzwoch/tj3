// Copyright 2023-2024 Florian Zwoch <fzwoch@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tj3

import (
	"bytes"
	"encoding/base64"
	"image/color"
	"image/jpeg"
	"testing"
)

func TestDecodeConfigSmallerThanBuffer(t *testing.T) {
	b, _ := base64.StdEncoding.DecodeString(gray)
	_, err := DecodeConfig(bytes.NewReader(b[:1024]))
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecodeConfigTooShort(t *testing.T) {
	b, _ := base64.StdEncoding.DecodeString(gray)
	_, err := DecodeConfig(bytes.NewReader(b[:128]))
	if err == nil {
		t.Fatal(err)
	}
}

func TestDecodeConfigGray(t *testing.T) {
	b, _ := base64.StdEncoding.DecodeString(gray)
	cfg, err := DecodeConfig(bytes.NewReader(b))
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
	b, _ := base64.StdEncoding.DecodeString(rgb)
	cfg, err := DecodeConfig(bytes.NewReader(b))
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
	b, _ := base64.StdEncoding.DecodeString(yuv420)
	cfg, err := DecodeConfig(bytes.NewReader(b))
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
	b, _ := base64.StdEncoding.DecodeString(yuv422)
	cfg, err := DecodeConfig(bytes.NewReader(b))
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
	b, _ := base64.StdEncoding.DecodeString(yuv444)
	cfg, err := DecodeConfig(bytes.NewReader(b))
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
	b, _ := base64.StdEncoding.DecodeString(cmyk)
	cfg, err := DecodeConfig(bytes.NewReader(b))
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
	b, _ := base64.StdEncoding.DecodeString(gray)

	img, err := Decode(bytes.NewReader(b))
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
	b, _ := base64.StdEncoding.DecodeString(rgb)

	img, err := Decode(bytes.NewReader(b))
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
	b, _ := base64.StdEncoding.DecodeString(yuv420)

	img, err := Decode(bytes.NewReader(b))
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
	b, _ := base64.StdEncoding.DecodeString(yuv422)

	img, err := Decode(bytes.NewReader(b))
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
	b, _ := base64.StdEncoding.DecodeString(yuv444)

	img, err := Decode(bytes.NewReader(b))
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
	b, _ := base64.StdEncoding.DecodeString(cmyk)

	img, err := Decode(bytes.NewReader(b))
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

func BenchmarkDecodeYUV420_TurboJPEG(b *testing.B) {
	r, _ := base64.StdEncoding.DecodeString(yuv420)
	for i := 0; i < b.N; i++ {
		Decode(bytes.NewReader(r))
	}
}

func BenchmarkDecodeYUV420_GoJPEG(b *testing.B) {
	r, _ := base64.StdEncoding.DecodeString(yuv420)
	for i := 0; i < b.N; i++ {
		jpeg.Decode(bytes.NewReader(r))
	}
}

const gray = `
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
gDuFGFw//9k=`

const rgb = `
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
AJX/AP7no/4bp/6pd/5X/wD7noo/4bp/6pd/5X//ALno/wCG6f8Aql3/AJX/AP7nqa0g8+Qpu24G
c4zU1p+3F58hT/hWO3Azn+3s/wDtvU1p+3F58hT/AIVjtwM5/t7P/tvVn+zv+m3/AI7/APXqz/w2
x/1TT/yu/wD3PVn/AIbY/wCqaf8Ald/+56P7O/6bf+O//Xrxn9pv4x/8Lo/4R/8A4p3+wf7G+0/8
vv2nzvO8r/YTbjyvfO7tjnxn9pv4x/8AC6P+Ef8A+Kd/sH+xvtP/AC+/afO87yv9hNuPK987u2OT
+zv+m3/jv/168Z/s7/pt/wCO/wD168Z/s7/pt/47/wDXr1l/gDeJNqcT+LdLxpdubi8mjXzY4Y0t
jLI7BGLhFkUoDsO5drgfOit6y/wBvEm1OJ/Ful40u3NxeTRr5scMaWxlkdgjFwiyKUB2Hcu1wPnR
W9Zf4A3iTanE/i3S8aXbm4vJo182OGNLYyyOwRi4RZFKA7DuXa4HzorPuP2f76eaCDQfEH9oySwy
FfN09oRcSq4jRIAGZ5Ed/MAkZUULE7NtAp9x+z/fTzQQaD4g/tGSWGQr5untCLiVXEaJAAzPIjv5
gEjKihYnZtoFPuP2f76eaCDQfEH9oySwyFfN09oRcSq4jRIAGZ5Ed/MAkZUULE7NtAqeH9n0Lpz3
Wp+MotKlhvrazuLe505i9ubh4fJkk8t28qNkklbdJtG6NU6udk8P7PoXTnutT8ZRaVLDfW1ncW9z
pzF7c3Dw+TJJ5bt5UbJJK26TaN0ap1c7J4f2fQunPdan4yi0qWG+trO4t7nTmL25uHh8mSTy3byo
2SSVt0m0bo1Tq52Zlz8BtbQmzh1WN9XJR47Gazlti0b+WFBaULiX96pK42KAQ0isNtZlz8BtbQmz
h1WN9XJR47Gazlti0b+WFBaULiX96pK42KAQ0isNtZlz8BtbQmzh1WN9XJR47Gazlti0b+WFBaUL
iX96pK42KAQ0isNtfUvwT8aP8LvhFZ+E7vSH1VdDvJ7S4voJ1jiika6Z284uMRcTxbFyzPv+YR4b
b9S/BPxo/wALvhFZ+E7vSH1VdDvJ7S4voJ1jiika6Z284uMRcTxbFyzPv+YR4bb9S/BPxo/wu+EV
n4Tu9IfVV0O8ntLi+gnWOKKRrpnbzi4xFxPFsXLM+/5hHhtvVP8AHeb/AIRK31+18FXeoGe+WJbS
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`

const yuv420 = `
/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAUDBAQEAwUEBAQFBQUGBwwIBwcHBw8LCwkMEQ8SEhEP
ERETFhwXExQaFRERGCEYGh0dHx8fExciJCIeJBweHx7/2wBDAQUFBQcGBw4ICA4eFBEUHh4eHh4e
Hh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh7/wAARCADwAUADASIA
AhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQA
AAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3
ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWm
p6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEA
AwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSEx
BhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElK
U1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3
uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD6rooo
oAKKKKACqGuf8ei/9dB/I1fqhrn/AB6L/wBdB/I187xZ/wAibEf4TWh/ERjUUUV/Nx7AUUUUAFcv
4t/5CUf/AFxH82rqK5fxb/yEo/8AriP5tX6X4S/8lFH/AAS/I/PvE3/kRS/xRMeiiiv6lP5wCiii
gArnLv8A4+5v+ujfzNdHXOXf/H3N/wBdG/ma/MPE/wD3Sh/if5HZgviZFRRRX42eiFFFFAFDxD/y
Bp/+A/8AoQrkK6/xD/yBp/8AgP8A6EK5Cv6U8G/+RHV/6+y/9Igf154Af8k5X/6/S/8ASKYUUUV+
sH7gFFFFAGXr3/LH/gX9Ky61Ne/5Y/8AAv6Vl1/PPHX/ACPq/wD27/6RE9bDfwkFFFFfJG4UUUUA
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
JFIAhKbk8mP94A7hRhcOAf/Z`

const yuv422 = `
/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAUDBAQEAwUEBAQFBQUGBwwIBwcHBw8LCwkMEQ8SEhEP
ERETFhwXExQaFRERGCEYGh0dHx8fExciJCIeJBweHx7/2wBDAQUFBQcGBw4ICA4eFBEUHh4eHh4e
Hh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh7/wAARCADwAUADASEA
AhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQA
AAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3
ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWm
p6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEA
AwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSEx
BhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElK
U1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3
uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD6rooA
KKACqGuf8ei/9dB/I187xZ/yJsR/hNaH8RGNRX83HsBRQAVy/i3/AJCUf/XEfzav0vwl/wCSij/g
l+R+feJv/Iil/iiY9Ff1KfzgFFABXOXf/H3N/wBdG/ma/MPE/wD3Sh/if5HZgviZFRX42eiFFAFD
xD/yBp/+A/8AoQrkK/pTwb/5EdX/AK+y/wDSIH9eeAH/ACTlf/r9L/0imFFfrB+4BRQBl69/yx/4
F/Ssuv5546/5H1f/ALd/9Iiethv4SCivkjcKKAOP8Rf8hmf/AID/AOgis+v7S4V/5EeC/wCvVP8A
9IR/JPEn/I4xf/Xyf/pTCivePFCigD9NqK/kw90KKACqGuf8ei/9dB/I187xZ/yJsR/hNaH8RGNR
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
`

const yuv444 = `
/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAUDBAQEAwUEBAQFBQUGBwwIBwcHBw8LCwkMEQ8SEhEP
ERETFhwXExQaFRERGCEYGh0dHx8fExciJCIeJBweHx7/2wBDAQUFBQcGBw4ICA4eFBEUHh4eHh4e
Hh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh7/wAARCADwAUADAREA
AhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQA
AAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3
ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWm
p6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEA
AwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSEx
BhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElK
U1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3
uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD6roAK
ACgAoAKAKGuf8ei/9dB/I1+eeJf/ACKqf/Xxf+kzOvB/G/Qxq/Dz0goAKACgAoA5fxb/AMhKP/ri
P5tX9LeDP/Ijq/8AX2X/AKRA/n3xY/5HFP8A69r/ANKmY9frR+YhQAUAFABQBzl3/wAfc3/XRv5m
v5iz/wD5GuJ/6+T/APSmezS+BeiIq8k0CgAoAKAKHiH/AJA0/wDwH/0IV934Zf8AJT4b/t//ANNz
P0vwe/5LLB/9xP8A01M5Cv6xP7hCgAoAKACgDL17/lj/AMC/pX5N4pf8wv8A2/8A+2HfgvtfIy6/
JTuCgAoAKACgDj/EX/IZn/4D/wCgiv608Mf+SXwv/b//AKcmfzH4h/8AJRYj/tz/ANIiZ9feHxYU
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`

const cmyk = `
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`
