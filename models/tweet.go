package models

type Tweet struct {
	Id    int    `form:"id" json:"id" xml:"id"  binding:"required"`
	Title string `form:"title" json:"title" xml:"title"  binding:"required"`
}
