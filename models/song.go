package models

import (
  "github.com/jinzhu/gorm"
)

type Song struct {
  gorm.Model

  Id     string `jsaon:"Id"`

  Title  string `json:"title"`
  Artist string `json:"artist"`
  Album  string `json:"album"`

  Votes  int    `json:"score"`
}
