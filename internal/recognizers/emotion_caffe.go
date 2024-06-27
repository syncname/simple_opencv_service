package recognizers

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

type EmotionCaffe struct {
	Model  string
	Config string
	Net    gocv.Net
}

func (e *EmotionCaffe) Close() error {
	return e.Net.Close()
}

func NewEmotionCaffe(model, config string) (*EmotionCaffe, error) {

	emotionNet := gocv.ReadNet(model, config)
	if emotionNet.Empty() {
		return nil, fmt.Errorf("%v %v %v\n", ErrModelReading, model, config)
	}

	var emotion = EmotionCaffe{
		Model:  model,
		Config: config,
		Net:    emotionNet,
	}

	err := emotionNet.SetPreferableBackend(gocv.NetBackendDefault)
	if err != nil {
		return nil, err
	}
	err = emotionNet.SetPreferableTarget(gocv.NetTargetCPU)

	if err != nil {
		return nil, err
	}

	return &emotion, nil
}

func (e *EmotionCaffe) GetEmotion(face *gocv.Mat) string {
	scalar := gocv.NewScalar(0, 0, 0, 0)
	blob := gocv.BlobFromImage(*face, ratio, image.Pt(227, 227), scalar, swapRGB, false)
	e.Net.SetInput(blob, "")
	emoPreds := e.Net.Forward("")
	_, _, _, emoLoc := gocv.MinMaxLoc(emoPreds)
	return EmotionsCaffe[emoLoc.X]
}

func (r *Recognizer) GetCaffeEmotion(mat *gocv.Mat) ([]EmotionResponse, error) {
	faces, errs := r.Facebox.GetFaces(mat)
	if len(faces) == 0 && errs != nil {
		return nil, errs
	}
	imgs := r.Facebox.ExtractFacesImg(mat, faces)

	var res = make([]EmotionResponse, 0, len(faces))

	for i, img := range imgs {
		res = append(res, EmotionResponse{
			Coordinates: faces[i],
			Emotion:     r.EmotionCaffe.GetEmotion(&img),
		})
	}

	return res, errs

}
