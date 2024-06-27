package recognizers

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

type Gender struct {
	Model  string
	Config string
	Net    gocv.Net
}

func (g *Gender) Close() error {
	return g.Net.Close()
}

func NewGender(model, config string) (*Gender, error) {

	genderNet := gocv.ReadNet(model, config)
	if genderNet.Empty() {
		return nil, fmt.Errorf("%v %v %v\n", ErrModelReading, model, config)
	}

	var emotion = Gender{
		Model:  model,
		Config: config,
		Net:    genderNet,
	}

	err := genderNet.SetPreferableBackend(gocv.NetBackendDefault)
	if err != nil {
		return nil, err
	}
	err = genderNet.SetPreferableTarget(gocv.NetTargetCPU)

	if err != nil {
		return nil, err
	}

	return &emotion, nil
}

func (g *Gender) GetGender(face *gocv.Mat) string {
	scalar := gocv.NewScalar(0, 0, 0, 0)
	blob := gocv.BlobFromImage(*face, ratio, image.Pt(227, 227), scalar, swapRGB, false)
	g.Net.SetInput(blob, "")
	genderPreds := g.Net.Forward("")
	_, _, _, ageLoc := gocv.MinMaxLoc(genderPreds)
	return Genders[ageLoc.X]
}

func (r *Recognizer) GetGender(mat *gocv.Mat) ([]GenderResponse, error) {
	faces, errs := r.Facebox.GetFaces(mat)
	if len(faces) == 0 && errs != nil {
		return nil, errs
	}
	imgs := r.Facebox.ExtractFacesImg(mat, faces)

	var res = make([]GenderResponse, 0, len(faces))

	for i, img := range imgs {
		res = append(res, GenderResponse{
			Coordinates: faces[i],
			Gender:      r.Gender.GetGender(&img),
		})
	}

	return res, errs

}
