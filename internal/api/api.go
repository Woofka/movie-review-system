package api

import (
	"context"

	"github.com/pkg/errors"
	reviewPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review"
	cachePkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/cache"
	"gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review/models"
	pb "gitlab.ozon.dev/Woofka/movie-review-system/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(review reviewPkg.Interface) pb.AdminServer {
	return &implementation{
		review: review,
	}
}

type implementation struct {
	pb.UnimplementedAdminServer
	review reviewPkg.Interface
}

func (i *implementation) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	err := i.review.Create(ctx, &models.Review{
		Reviewer:   req.Review.GetReviewer(),
		MovieTitle: req.Review.GetMovieTitle(),
		Text:       req.Review.GetText(),
		Rating:     uint8(req.Review.GetRating()),
	})
	if err != nil {
		if errors.Is(err, reviewPkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateReviewResponse{}, nil
}

func (i *implementation) GetReview(ctx context.Context, req *pb.GetReviewRequest) (*pb.GetReviewResponse, error) {
	review, err := i.review.Get(ctx, uint(req.GetId()))
	if err != nil {
		if errors.Is(err, cachePkg.ErrReviewNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetReviewResponse{
		Review: &pb.Review{
			Id:         uint64(review.Id),
			Reviewer:   review.Reviewer,
			MovieTitle: review.MovieTitle,
			Text:       review.Text,
			Rating:     uint32(review.Rating),
		},
	}, nil
}

func (i *implementation) UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.UpdateReviewResponse, error) {
	err := i.review.Update(ctx, &models.Review{
		Id:         uint(req.Review.GetId()),
		Reviewer:   req.Review.GetReviewer(),
		MovieTitle: req.Review.GetMovieTitle(),
		Text:       req.Review.GetText(),
		Rating:     uint8(req.Review.GetRating()),
	})
	if err != nil {
		if errors.Is(err, reviewPkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateReviewResponse{}, nil
}

func (i *implementation) DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*pb.DeleteReviewResponse, error) {
	err := i.review.Delete(ctx, uint(req.GetId()))
	if err != nil {
		if errors.Is(err, cachePkg.ErrReviewNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteReviewResponse{}, nil
}

func (i *implementation) ListReview(ctx context.Context, _ *pb.ListReviewRequest) (*pb.ListReviewResponse, error) {
	reviews, err := i.review.List(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result := make([]*pb.Review, 0, len(reviews))
	for _, review := range reviews {
		result = append(result, &pb.Review{
			Id:         uint64(review.Id),
			Reviewer:   review.Reviewer,
			MovieTitle: review.MovieTitle,
			Text:       review.Text,
			Rating:     uint32(review.Rating),
		})
	}

	return &pb.ListReviewResponse{
		Reviews: result,
	}, nil
}
