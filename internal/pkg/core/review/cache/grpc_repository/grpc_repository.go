package grpc_repository

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/cache"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/models"
	pb "gitlab.ozon.dev/Woofka/movie-review-system/pkg/repository_api"
)

func New(client pb.RepositoryClient) cache.Interface {
	return &Repository{client}
}

type Repository struct {
	client pb.RepositoryClient
}

func (r *Repository) List(ctx context.Context, limit, offset uint, orderDesc bool) ([]*models.Review, error) {
	response, err := r.client.ListReview(ctx, &pb.ListReviewRequest{
		Limit:     uint64(limit),
		Offset:    uint64(offset),
		OrderDesc: orderDesc,
	})
	if err != nil {
		return nil, errors.Wrap(err, "remote repository List error")
	}

	reviews := make([]*models.Review, 0, len(response.Reviews))
	for _, responseReview := range response.Reviews {
		reviews = append(reviews, &models.Review{
			Id:         uint(responseReview.GetId()),
			Reviewer:   responseReview.GetReviewer(),
			MovieTitle: responseReview.GetMovieTitle(),
			Text:       responseReview.GetText(),
			Rating:     uint8(responseReview.GetRating()),
		})
	}

	return reviews, nil
}

func (r *Repository) Add(ctx context.Context, review *models.Review) error {
	_, err := r.client.CreateReview(ctx, &pb.CreateReviewRequest{
		Review: &pb.Review{
			Reviewer:   review.Reviewer,
			MovieTitle: review.MovieTitle,
			Text:       review.Text,
			Rating:     uint32(review.Rating),
		},
	})
	if err != nil {
		return errors.Wrap(err, "remote repository Add error")
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, id uint) (*models.Review, error) {
	response, err := r.client.GetReview(ctx, &pb.GetReviewRequest{Id: uint64(id)})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, errors.Wrap(cache.ErrReviewNotExists, err.Error())
		}
		return nil, errors.Wrap(err, "remote repository Get error")
	}

	review := response.GetReview()
	return &models.Review{
		Id:         uint(review.GetId()),
		Reviewer:   review.GetReviewer(),
		MovieTitle: review.GetMovieTitle(),
		Text:       review.GetText(),
		Rating:     uint8(review.GetRating()),
	}, nil
}

func (r *Repository) Update(ctx context.Context, review *models.Review) error {
	_, err := r.client.UpdateReview(ctx, &pb.UpdateReviewRequest{
		Review: &pb.Review{
			Reviewer:   review.Reviewer,
			MovieTitle: review.MovieTitle,
			Text:       review.Text,
			Rating:     uint32(review.Rating),
			Id:         uint64(review.Id),
		},
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return errors.Wrap(cache.ErrReviewNotExists, err.Error())
		}
		return errors.Wrap(err, "remote repository Update error")
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id uint) error {
	_, err := r.client.DeleteReview(ctx, &pb.DeleteReviewRequest{Id: uint64(id)})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return errors.Wrap(cache.ErrReviewNotExists, err.Error())
		}
		return errors.Wrap(err, "remote repository Delete error")
	}

	return nil
}
