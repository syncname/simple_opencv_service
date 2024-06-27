package recognizers

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

type Age struct {
	Model  string
	Config string
	Net    gocv.Net
}

func (a *Age) Close() error {
	return a.Net.Close()
}

func NewAge(model, config string) (*Age, error) {

	ageNet := gocv.ReadNet(model, config)
	if ageNet.Empty() {
		return nil, fmt.Errorf("%v %v %v\n", ErrModelReading, model, config)
	}

	var age = Age{
		Model:  model,
		Config: config,
		Net:    ageNet,
	}

	err := ageNet.SetPreferableBackend(gocv.NetBackendDefault)
	if err != nil {
		return nil, err
	}
	err = ageNet.SetPreferableTarget(gocv.NetTargetCPU)

	if err != nil {
		return nil, err
	}

	return &age, nil
}

func (a *Age) GetAge(face *gocv.Mat) string {
	scalar := gocv.NewScalar(0, 0, 0, 0)
	blob := gocv.BlobFromImage(*face, ratio, image.Pt(227, 227), scalar, swapRGB, false)
	a.Net.SetInput(blob, "")
	agePreds := a.Net.Forward("")
	_, _, _, ageLoc := gocv.MinMaxLoc(agePreds)
	return Ages[ageLoc.X]
}

func (r *Recognizer) GetAge(mat *gocv.Mat) ([]AgeResponse, error) {
	faces, errs := r.Facebox.GetFaces(mat)
	if len(faces) == 0 && errs != nil {
		return nil, errs
	}
	imgs := r.Facebox.ExtractFacesImg(mat, faces)

	var res = make([]AgeResponse, 0, len(faces))

	for i, img := range imgs {
		res = append(res, AgeResponse{
			Coordinates: faces[i],
			Age:         r.Age.GetAge(&img),
		})
	}

	return res, errs

}
