package recognizers

import (
	"cvrecognizer/internal/config"
	"gocv.io/x/gocv"
)

type Recognizer struct {
	*Facebox
	*Age
	*EmotionCaffe
	*EmotionONNX
	*Gender
}

func New(cfg *config.Config) (*Recognizer, error) {
	facebox, err := NewFacebox(cfg.Facebox.Model, cfg.Facebox.Config)
	if err != nil {
		return nil, err
	}

	age, err := NewAge(cfg.Age.Model, cfg.Age.Config)
	if err != nil {
		return nil, err
	}
	emotionCaffe, err := NewEmotionCaffe(cfg.EmotionCaffe.Model, cfg.EmotionCaffe.Config)
	if err != nil {
		return nil, err
	}
	emotionONNX, err := NewEmotionONNX(cfg.EmotionOnnx.Model)
	if err != nil {
		return nil, err
	}
	gender, err := NewGender(cfg.Gender.Model, cfg.Gender.Config)
	if err != nil {
		return nil, err
	}

	r := &Recognizer{
		Facebox:      facebox,
		Age:          age,
		EmotionCaffe: emotionCaffe,
		EmotionONNX:  emotionONNX,
		Gender:       gender,
	}

	return r, nil
}

func (r *Recognizer) GetFullIno(mat *gocv.Mat) ([]FullInfo, error) {
	faces, errs := r.Facebox.GetFaces(mat)
	if len(faces) == 0 && errs != nil {
		return nil, errs
	}
	imgs := r.Facebox.ExtractFacesImg(mat, faces)

	var res = make([]FullInfo, 0, len(faces))

	for i, img := range imgs {
		res = append(res, FullInfo{
			Coordinates:  faces[i],
			Gender:       r.Gender.GetGender(&img),
			Age:          r.Age.GetAge(&img),
			EmotionCaffe: r.EmotionCaffe.GetEmotion(&img),
			EmotionOnnx:  r.EmotionONNX.GetEmotion(&img),
		})
	}

	return res, errs

}
