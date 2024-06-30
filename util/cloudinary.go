package util

import (
	"context"
	"mime/multipart"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryProvider interface {
	Upload(ctx context.Context, file multipart.File, params uploader.UploadParams) (*uploader.UploadResult, error)
}

type cloudinaryProviderImpl struct {
	cloud *cloudinary.Cloudinary
}

type CloudinaryProviderOpts struct {
	CloudinaryName      string
	CloudinaryAPIKey    string
	CloudinaryAPISecret string
}

func NewCloudinaryProvider(opts CloudinaryProviderOpts) (*cloudinaryProviderImpl, error) {
	c, err := cloudinary.NewFromParams(opts.CloudinaryName, opts.CloudinaryAPIKey, opts.CloudinaryAPISecret)
	if err != nil {
		return nil, apperror.Wrap(err)
	}
	return &cloudinaryProviderImpl{
		cloud: c,
	}, nil
}

func (p *cloudinaryProviderImpl) Upload(ctx context.Context, file multipart.File, params uploader.UploadParams) (*uploader.UploadResult, error) {
	res, err := p.cloud.Upload.Upload(ctx, file, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}
