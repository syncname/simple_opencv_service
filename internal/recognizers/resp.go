package recognizers

var Genders = []string{"Male", "Female"}
var Ages = []string{"0-2", "3-7", "8-12", "13-20", "20-36", "37-47", "48-55", "56-100"}
var EmotionsONNX = []string{"Neutral", "Happy", "Surprise", "Sad", "Anger", "Disgust", "Fear", "Contempt"}
var EmotionsCaffe = []string{"Angry", "Disgust", "Fear", "Happy", "Neutral", "Sad", "Surprise"}

type EmotionResponse struct {
	Coordinates FaceBoxResult `json:"coordinates"`
	Emotion     string        `json:"emotion"`
}

type AgeResponse struct {
	Coordinates FaceBoxResult `json:"coordinates"`
	Age         string        `json:"age"`
}

type GenderResponse struct {
	Coordinates FaceBoxResult `json:"coordinates"`
	Gender      string        `json:"age"`
}

type FullInfo struct {
	Coordinates  FaceBoxResult `json:"coordinates"`
	EmotionCaffe string        `json:"emotion_caffe"`
	EmotionOnnx  string        `json:"emotion_onnx"`
	Age          string        `json:"age"`
	Gender       string        `json:"gender"`
}
