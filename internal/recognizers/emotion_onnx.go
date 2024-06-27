package recognizers

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

type EmotionONNX struct {
	Model string
	Net   gocv.Net
}

func (e *EmotionONNX) Close() error {
	return e.Net.Close()
}

func NewEmotionONNX(model string) (*EmotionONNX, error) {

	emotionNet := gocv.ReadNetFromONNX(model)
	if emotionNet.Empty() {
		return nil, fmt.Errorf("%v %v\n", ErrModelReading, model)
	}

	var emotion = EmotionONNX{
		Model: model,
		Net:   emotionNet,
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

func (e *EmotionONNX) GetEmotion(face *gocv.Mat) string {

	meanVal := gocv.Scalar{Val1: 78.4263377603,
		Val2: 87.7689143744,
		Val3: 114.895847746}

	//создавать нужно так
	var imgFER = gocv.NewMat()

	gocv.CvtColor(*face, &imgFER, gocv.ColorBGRToGray)
	imgFER.ConvertTo(&imgFER, gocv.MatTypeCV32FC1)
	blobFER := gocv.BlobFromImage(imgFER, ratio, image.Pt(64, 64),
		meanVal, false, swapRGB)

	e.Net.SetInput(blobFER, "")
	output := e.Net.Forward("")

	_, _, _, emoONNXLoc := gocv.MinMaxLoc(output)
	return EmotionsONNX[emoONNXLoc.X]
}

func (r *Recognizer) GetOnnxEmotion(mat *gocv.Mat) ([]EmotionResponse, error) {
	faces, errs := r.Facebox.GetFaces(mat)
	if len(faces) == 0 && errs != nil {
		return nil, errs
	}
	imgs := r.Facebox.ExtractFacesImg(mat, faces)

	var res = make([]EmotionResponse, 0, len(faces))

	for i, img := range imgs {
		res = append(res, EmotionResponse{
			Coordinates: faces[i],
			Emotion:     r.EmotionONNX.GetEmotion(&img),
		})
	}

	return res, errs

}
