package storage

import "fmt"

var lastId = uint(0)

type Review struct {
	id         uint
	reviewer   string
	movieTitle string
	text       string
	rating     uint8
}

func NewReview(reviewer, movieTitle, text string, rating uint8) (*Review, error) {
	r := Review{}
	if err := r.SetReviewer(reviewer); err != nil {
		return nil, err
	}
	if err := r.SetMovieTitle(movieTitle); err != nil {
		return nil, err
	}
	if err := r.SetText(text); err != nil {
		return nil, err
	}
	if err := r.SetRating(rating); err != nil {
		return nil, err
	}
	lastId++
	r.id = lastId
	return &r, nil
}

func (r *Review) SetReviewer(reviewer string) error {
	if len(reviewer) == 0 || len(reviewer) > 32 {
		return fmt.Errorf("invalid reviewer length: %d. Should be 1..32", len(reviewer))
	}
	r.reviewer = reviewer
	return nil
}

func (r *Review) SetMovieTitle(title string) error {
	if len(title) == 0 || len(title) > 50 {
		return fmt.Errorf("invalid movie title length: %d. Should be 1..50", len(title))
	}
	r.movieTitle = title
	return nil
}

func (r *Review) SetText(text string) error {
	if len(text) == 0 || len(text) > 200 {
		return fmt.Errorf("invalid text length: %d. Should be 1..200", len(text))
	}
	r.text = text
	return nil
}

func (r *Review) SetRating(rating uint8) error {
	if rating > 10 {
		return fmt.Errorf("invalid rating value: %d. Should be 0..10", rating)
	}
	r.rating = rating
	return nil
}

func (r Review) GetId() uint {
	return r.id
}

func (r Review) GetReviewer() string {
	return r.reviewer
}

func (r Review) GetMovieTitle() string {
	return r.movieTitle
}

func (r Review) GetText() string {
	return r.text
}

func (r Review) GetRating() uint8 {
	return r.rating
}

func (r Review) String() string {
	return fmt.Sprintf("#%d %s\n%d/10 by %s\n%s", r.id, r.movieTitle, r.rating, r.reviewer, r.text)
}
