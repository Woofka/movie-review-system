package models

import "fmt"

type Review struct {
	Id         uint
	Reviewer   string
	MovieTitle string
	Text       string
	Rating     uint8
}

func (r Review) String() string {
	return fmt.Sprintf("#%d %s\n%d/10 by %s\n%s", r.Id, r.MovieTitle, r.Rating, r.Reviewer, r.Text)
}
