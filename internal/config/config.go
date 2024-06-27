package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
	Facebox struct {
		Model  string `json:"model"`
		Config string `json:"config"`
	} `json:"facebox"`
	EmotionCaffe struct {
		Model  string `json:"model"`
		Config string `json:"config"`
	} `json:"emotion_caffe"`
	EmotionOnnx struct {
		Model string `json:"model"`
	} `json:"emotion_onnx"`
	Gender struct {
		Model  string `json:"model"`
		Config string `json:"config"`
	} `json:"gender"`
	Age struct {
		Model  string `json:"model"`
		Config string `json:"config"`
	} `json:"age"`
}

func ReadConfig(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil

}
