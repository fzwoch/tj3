// Copyright 2023 Florian Zwoch <fzwoch@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tj3

// #include <turbojpeg.h>
import "C"
import (
	"errors"
	"image"
	"image/color"
	"io"
)

func Decode(r io.Reader) (image.Image, error) {
	ctx := C.tj3Init(C.TJINIT_DECOMPRESS)
	if ctx == nil {
		return nil, errors.New("tj3Init() failed")
	}
	defer C.tj3Destroy(ctx)

	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	ret := C.tj3DecompressHeader(ctx, (*C.uchar)(&b[0]), C.size_t(len(b)))
	if ret != 0 {
		return nil, errors.New(C.GoString(C.tj3GetErrorStr(ctx)))
	}

	if C.tj3Get(ctx, C.TJPARAM_PRECISION) != 8 {
		return nil, errors.New("image is not 8-bit")
	}

	bounds := image.Rectangle{
		Max: image.Point{
			X: int(C.tj3Get(ctx, C.TJPARAM_JPEGWIDTH)),
			Y: int(C.tj3Get(ctx, C.TJPARAM_JPEGHEIGHT)),
		},
	}

	switch C.tj3Get(ctx, C.TJPARAM_COLORSPACE) {
	case C.TJCS_GRAY:
		img := image.NewGray(bounds)
		ret = C.tj3Decompress8(ctx, (*C.uchar)(&b[0]), C.size_t(len(b)), (*C.uchar)(&img.Pix[0]), 0, C.TJPF_GRAY)
		if ret != 0 {
			return nil, errors.New(C.GoString(C.tj3GetErrorStr(ctx)))
		}
		return img, nil
	case C.TJCS_YCCK:
		img := image.NewCMYK(bounds)
		ret = C.tj3Decompress8(ctx, (*C.uchar)(&b[0]), C.size_t(len(b)), (*C.uchar)(&img.Pix[0]), 0, C.TJPF_CMYK)
		if ret != 0 {
			return nil, errors.New(C.GoString(C.tj3GetErrorStr(ctx)))
		}
		return img, nil
	case C.TJCS_RGB:
		img := image.NewNRGBA(bounds)
		ret = C.tj3Decompress8(ctx, (*C.uchar)(&b[0]), C.size_t(len(b)), (*C.uchar)(&img.Pix[0]), 0, C.TJPF_RGBA)
		if ret != 0 {
			return nil, errors.New(C.GoString(C.tj3GetErrorStr(ctx)))
		}
		return img, nil
	case C.TJCS_YCbCr:
		var ss image.YCbCrSubsampleRatio
		switch C.tj3Get(ctx, C.TJPARAM_SUBSAMP) {
		case C.TJSAMP_420:
			ss = image.YCbCrSubsampleRatio420
		case C.TJSAMP_422:
			ss = image.YCbCrSubsampleRatio422
		case C.TJSAMP_444:
			ss = image.YCbCrSubsampleRatio444
		}
		img := image.NewYCbCr(bounds, ss)
		ret = C.tj3DecompressToYUV8(ctx, (*C.uchar)(&b[0]), C.size_t(len(b)), (*C.uchar)(&img.Y[0]), 1)
		if ret != 0 {
			return nil, errors.New(C.GoString(C.tj3GetErrorStr(ctx)))
		}
		return img, nil
	}

	return nil, errors.New("")
}

func DecodeConfig(r io.Reader) (image.Config, error) {
	ctx := C.tj3Init(C.TJINIT_DECOMPRESS)
	if ctx == nil {
		return image.Config{}, errors.New("tj3Init() failed")
	}
	defer C.tj3Destroy(ctx)

	b := make([]byte, 4096)
	s, err := io.ReadFull(r, b)
	if err != nil && err != io.ErrUnexpectedEOF {
		return image.Config{}, err
	}

	ret := C.tj3DecompressHeader(ctx, (*C.uchar)(&b[0]), C.size_t(s))
	if ret != 0 {
		return image.Config{}, errors.New(C.GoString(C.tj3GetErrorStr(ctx)))
	}

	if C.tj3Get(ctx, C.TJPARAM_PRECISION) != 8 {
		return image.Config{}, errors.New("image is not 8-bit")
	}

	w := int(C.tj3Get(ctx, C.TJPARAM_JPEGWIDTH))
	h := int(C.tj3Get(ctx, C.TJPARAM_JPEGHEIGHT))

	switch C.tj3Get(ctx, C.TJPARAM_COLORSPACE) {
	case C.TJCS_GRAY:
		return image.Config{
			ColorModel: color.GrayModel,
			Width:      w,
			Height:     h,
		}, nil
	case C.TJCS_YCCK:
		return image.Config{
			ColorModel: color.CMYKModel,
			Width:      w,
			Height:     h,
		}, nil
	case C.TJCS_RGB:
		return image.Config{
			ColorModel: color.NRGBAModel,
			Width:      w,
			Height:     h,
		}, nil
	case C.TJCS_YCbCr:
		return image.Config{
			ColorModel: color.YCbCrModel,
			Width:      w,
			Height:     h,
		}, nil
	}

	return image.Config{}, errors.New("")
}

func init() {
	image.RegisterFormat("tj3", "\xff\xd8", Decode, DecodeConfig)
}
