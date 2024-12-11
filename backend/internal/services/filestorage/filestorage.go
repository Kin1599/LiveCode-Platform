package filestorage

import (
	"fmt"
)

type S3Service struct {
	s3Uploader   StorageUploder
	s3Downloader StorageDownloader
}

type StorageUploder interface {
	UploadProject(projectId string, projectStructure []byte) error
}

type StorageDownloader interface {
	DownloadProject(projectName string) ([]byte, error)
	DeleteFolder(folderPath string) error
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

func (s3Serv *S3Service) DeleteProject(projectId string) error {
	const op = "Services.FileStorage.DeleteProject"
	err := s3Serv.s3Downloader.DeleteFolder(projectId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}