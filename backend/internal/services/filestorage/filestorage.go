package filestorage

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client     *s3.Client
	bucketName string
	wg         *sync.WaitGroup
}

type S3Service struct {
	s3Uploader   StorageUploder
	s3Downloader StorageDownloader
}

type StorageUploder interface {
	UploadProject(projectId string, projectStructure []byte) error
}

type StorageDownloader interface {
	DownloadProject(projectName string) ([]byte, error)
}

func New(
	storageUploader StorageUploder,
	storageProviderr StorageDownloader,
) *S3Service {
	return &S3Service{
		s3Uploader:   storageUploader,
		s3Downloader: storageProviderr,
	}
}

func (s3Serv *S3Service) SaveNewProject(projectId string, projectStructure []byte) error {
	const op = "Services.FileStorage.SaveNewProject"
	err := s3Serv.s3Uploader.UploadProject(projectId, projectStructure)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s3Serv *S3Service) GetProject(projectId string) ([]byte, error) {
	const op = "Services.FileStorage.GetProject"
	prjStructure, err := s3Serv.s3Downloader.DownloadProject(projectId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return prjStructure, nil
}
