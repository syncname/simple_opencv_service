package recognizers

import (
	"errors"
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

type FaceConfig struct {
	Ratio      float64
	Scalar     gocv.Scalar
	swapRGB    bool
	Pt         image.Point
	Confidence float32
}

type Facebox struct {
	Model   string
	Config  string
	FaceNet gocv.Net
	Conf    FaceConfig
}

type FaceBoxResult struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

func (f *Facebox) Close() error {
	return f.FaceNet.Close()
}

func (f *Facebox) ExtractFacesImg(img *gocv.Mat, coord []FaceBoxResult) []gocv.Mat {
	var faces = make([]gocv.Mat, 0, len(coord))
	for _, faceCoord := range coord {
		r := image.Rect(faceCoord.Left, faceCoord.Top, faceCoord.Right, faceCoord.Bottom)
		mat := img.Region(r)
		faces = append(faces, mat)
	}
	return faces
}

func (f *Facebox) GetFaces(img *gocv.Mat) ([]FaceBoxResult, error) {
	if img.Empty() {
		return nil, ErrEmptyImage
	}

	blob := gocv.BlobFromImage(*img, f.Conf.Ratio, f.Conf.Pt,
		f.Conf.Scalar, f.Conf.swapRGB, false)

	f.FaceNet.SetInput(blob, "data")

	var faces []FaceBoxResult
	var faceErrors []error
	faceImg := f.FaceNet.Forward("detection_out")
	for i := 0; i < faceImg.Total(); i += 7 {
		confidence := faceImg.GetFloatAt(0, i+2)
		if confidence > f.Conf.Confidence {
			left := int(faceImg.GetFloatAt(0, i+3) * float32(img.Cols()))
			top := int(faceImg.GetFloatAt(0, i+4) * float32(img.Rows()))
			right := int(faceImg.GetFloatAt(0, i+5) * float32(img.Cols()))
			bottom := int(faceImg.GetFloatAt(0, i+6) * float32(img.Rows()))
			r := image.Rect(left, top, right, bottom)
			if r.Max.X < img.Cols() && r.Max.Y < img.Rows() && r.Min.X > 0 && r.Min.Y > 0 {
				faces = append(faces, FaceBoxResult{
					Left:   left,
					Top:    top,
					Right:  right,
					Bottom: bottom,
				})
			} else {
				faceErrors = append(faceErrors, fmt.Errorf("facebox: bad coordinates rectangle: %v", r))
			}

		}

	}

	return faces, errors.Join(faceErrors...)
}

func NewFacebox(model, config string) (*Facebox, error) {

	faceNet := gocv.ReadNet(model, config)
	if faceNet.Empty() {
		return nil, fmt.Errorf("reading model error: %v %v\n", model, config)
	}

	var facebox = Facebox{
		Model:   model,
		Config:  config,
		FaceNet: faceNet,
		Conf: FaceConfig{
			Ratio:      1,
			Scalar:     gocv.Scalar{Val1: 104, Val2: 177, Val3: 123},
			swapRGB:    false,
			Pt:         image.Pt(300, 300),
			Confidence: 0.6,
		},
	}

	err := faceNet.SetPreferableBackend(gocv.NetBackendDefault)
	if err != nil {
		return nil, err
	}
	err = faceNet.SetPreferableTarget(gocv.NetTargetCPU)

	if err != nil {
		return nil, err
	}

	return &facebox, nil
}
