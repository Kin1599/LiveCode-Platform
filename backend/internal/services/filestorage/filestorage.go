package filestorage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client     *s3.Client
	bucketName string
	wg         *sync.WaitGroup
}

func New(bucketName string) (*S3Client, error) {
	const op = "Services.FileStorage.New"

	configFiles := []string{"./configs/config", "./configs/credentials"}

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigFiles(configFiles),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	client := s3.NewFromConfig(cfg)

	return &S3Client{
		client:     client,
		bucketName: bucketName,
		wg:         new(sync.WaitGroup),
	}, nil
}

func (u *S3Client) uploadFile(path string, content string) {
	defer u.wg.Done()

	input := &s3.PutObjectInput{
		Bucket:      aws.String(u.bucketName),
		Key:         aws.String(path),
		Body:        strings.NewReader(content),
		ContentType: aws.String("text/plain"),
	}

	if _, err := u.client.PutObject(context.Background(), input); err != nil {
		return
	}
}

func (u *S3Client) uploadByStructure(basePath string, data interface{}) {
	defer u.wg.Done()

	switch v := data.(type) {
	case string:

		u.wg.Add(1)
		go u.uploadFile(basePath, v)

	case map[string]interface{}:
		for key, value := range v {
			newPath := fmt.Sprintf("%s/%s", basePath, key)
			u.wg.Add(1)
			go u.uploadByStructure(newPath, value)
		}
	}
}

func (cli *S3Client) UploadProject(projectId string, projectStructure []byte) error {
	const op = "Services.FileStorage.UploadNewProject"
	var mapStructure map[string]interface{}
	if err := json.Unmarshal(projectStructure, &mapStructure); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	cli.wg.Add(1)
	go cli.uploadByStructure(projectId, mapStructure)
	return nil
}

func (d *S3Client) DownloadProject(projectName string) ([]byte, error) {
	const op = "Services.FileStorage.DownloadProject"

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(d.bucketName),
		Prefix: aws.String(projectName + "/"),
	}

	result, err := d.client.ListObjectsV2(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	project := make(map[string]interface{})

	var mu sync.Mutex
	for _, item := range result.Contents {
		path := *aws.String(*item.Key)
		if strings.HasSuffix(path, "/") {
			continue
		}

		d.wg.Add(1)
		go func(path string) {
			defer d.wg.Done()
			content, err := d.downloadFile(path)
			if err != nil {
				log.Printf("failed to download file %q: %v", path, err)
				return
			}

			relativePath := strings.TrimPrefix(path, projectName+"/")
			mu.Lock()
			d.insertIntoStructure(&project, relativePath, content)
			mu.Unlock()

		}(path)
	}

	d.wg.Wait()

	projectJson, err := json.Marshal(project)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	projectStructure := []byte(projectJson)

	return projectStructure, nil
}

func (d *S3Client) downloadFile(key string) (string, error) {
	const op = "Services.FileStorage.DownloadFile"

	input := &s3.GetObjectInput{
		Bucket: aws.String(d.bucketName),
		Key:    aws.String(key),
	}

	result, err := d.client.GetObject(context.Background(), input)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer result.Body.Close()

	content, err := io.ReadAll(result.Body)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return string(content), nil
}
