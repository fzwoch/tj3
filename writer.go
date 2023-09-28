// Copyright 2023 Florian Zwoch <fzwoch@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tj3

// #include <turbojpeg.h>
import "C"
import (
	"errors"
	"image"
	"io"
	"runtime"
)

const DefaultQuality = 75

type Options struct {
	Quality int
}

func Encode(w io.Writer, m image.Image, o *Options) error {
	ctx := C.tj3Init(C.TJINIT_COMPRESS)
	if ctx == nil {
		return errors.New("tj3Init() failed")
	}
	defer C.tj3Destroy(ctx)

	if o.Quality < 0 || o.Quality > 100 {
		return errors.New("invalid quality option")
	}
	q := o.Quality
	if q == 0 {
		q = DefaultQuality
	}

	C.tj3Set(ctx, C.TJPARAM_NOREALLOC, 1)
	C.tj3Set(ctx, C.TJPARAM_QUALITY, C.int(q))

	var (
		pin runtime.Pinner
		s   C.size_t
	)

	switch m := m.(type) {
	case *image.Gray:
		C.tj3Set(ctx, C.TJPARAM_SUBSAMP, C.TJSAMP_GRAY)

		s = C.tj3JPEGBufSize(C.int(m.Rect.Dx()), C.int(m.Rect.Dy()), C.TJSAMP_GRAY)

		b := make([]byte, int(s))
		t := (*C.uchar)(&b[0])

		pin.Pin(t)
		ret := C.tj3Compress8(ctx, (*C.uchar)(&m.Pix[0]), C.int(m.Rect.Dx()), C.int(m.Stride), C.int(m.Rect.Dy()), C.TJPF_GRAY, &t, &s)
		pin.Unpin()
		if ret != 0 {
			return errors.New(C.GoString(C.tj3GetErrorStr(ctx)))
		}
		_, err := w.Write(b[:s])
		if err != nil {
			return err
		}
		return nil
	case *image.CMYK:
		C.tj3Set(ctx, C.TJPARAM_SUBSAMP, C.TJSAMP_444)

		s = C.tj3JPEGBufSize(C.int(m.Rect.Dx()), C.int(m.Rect.Dy()), C.TJSAMP_444)

		b := make([]byte, int(s))
		t := (*C.uchar)(&b[0])

		pin.Pin(t)
		ret := C.tj3Compress8(ctx, (*C.uchar)(&m.Pix[0]), C.int(m.Rect.Dx()), C.int(m.Stride), C.int(m.Rect.Dy()), C.TJPF_CMYK, &t, &s)
		pin.Unpin()
		if ret != 0 {
			return errors.New(C.GoString(C.tj3GetErrorStr(ctx)))
		}
		_, err := w.Write(b[:s])
		if err != nil {
			return err
		}
		return nil
	case *image.NRGBA:
		C.tj3Set(ctx, C.TJPARAM_SUBSAMP, C.TJSAMP_444)

		s = C.tj3JPEGBufSize(C.int(m.Rect.Dx()), C.int(m.Rect.Dy()), C.TJSAMP_444)

		b := make([]byte, int(s))
		t := (*C.uchar)(&b[0])

		pin.Pin(t)
		ret := C.tj3Compress8(ctx, (*C.uchar)(&m.Pix[0]), C.int(m.Rect.Dx()), C.int(m.Stride), C.int(m.Rect.Dy()), C.TJPF_RGBA, &t, &s)
		pin.Unpin()
		if ret != 0 {
			return errors.New(C.GoString(C.tj3GetErrorStr(ctx)))
		}
		_, err := w.Write(b[:s])
		if err != nil {
			return err
		}
		return nil
	case *image.YCbCr:
		var ss C.int

		switch m.SubsampleRatio {
		case image.YCbCrSubsampleRatio420:
			ss = C.TJSAMP_420
		case image.YCbCrSubsampleRatio422:
			ss = C.TJSAMP_422
		case image.YCbCrSubsampleRatio444:
			ss = C.TJSAMP_444
		}

		C.tj3Set(ctx, C.TJPARAM_SUBSAMP, ss)

		s = C.tj3JPEGBufSize(C.int(m.Rect.Dx()), C.int(m.Rect.Dy()), ss)

		b := make([]byte, int(s))
		t := (*C.uchar)(&b[0])

		planes := []*C.uchar{
			(*C.uchar)(&m.Y[0]),
			(*C.uchar)(&m.Cb[0]),
			(*C.uchar)(&m.Cr[0]),
		}
		strides := []C.int{
			C.int(m.YStride),
			C.int(m.CStride),
			C.int(m.CStride),
		}

		pin.Pin(t)
		pin.Pin(planes[0])
		pin.Pin(planes[1])
		pin.Pin(planes[2])
		ret := C.tj3CompressFromYUVPlanes8(ctx, &planes[0], C.int(m.Rect.Dx()), &strides[0], C.int(m.Rect.Dy()), &t, &s)
		pin.Unpin()
		if ret != 0 {
			return errors.New(C.GoString(C.tj3GetErrorStr(ctx)))
		}
		_, err := w.Write(b[:s])
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("")
}
