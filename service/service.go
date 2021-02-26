package service

import (
	"encoding/csv"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/resty.v1"
)

// Service for Pokemon requests
type Service struct {
	logger *logrus.Logger
	client *resty.Client
	csvr   *os.File
	csvw   *csv.Writer
}
