package uploadprovider

import (
	"context"
	"ielts/common"
)

type UpLoadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
