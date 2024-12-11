package database

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

func NewS3Client(bucketName string) (*S3Client, error) {
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

func (cli *S3Client) UploadProject(projectId string, projectStructure []byte) error {
	const op = "Services.FileStorage.UploadNewProject"
	var mapStructure map[string]interface{}
	if err := json.Unmarshal(projectStructure, &mapStructure); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	fmt.Println(mapStructure)
	cli.wg.Add(1)
	go cli.uploadByStructure(projectId, mapStructure)
	return nil
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

func (d *S3Client) downloadFile(key string) (string, error) {
	const op = "Database.FileStorage.DownloadFile"

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

func (d *S3Client) DownloadProject(projectName string) ([]byte, error) {
	const op = "Database.FileStorage.DownloadProject"

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

func (d *S3Client) insertIntoStructure(root *map[string]interface{}, path string, content string) {
	parts := strings.Split(path, "/")
	current := root

	for i, part := range parts {
		if i == len(parts)-1 {
			(*current)[part] = content
		} else {
			if _, exists := (*current)[part]; !exists {
				(*current)[part] = make(map[string]interface{})
			}

			nextMap, ok := (*current)[part].(map[string]interface{})
			if !ok {
				nextMap = make(map[string]interface{})
				(*current)[part] = nextMap
			}
			current = &nextMap
		}
	}
}

func (s *S3Client) DeleteFolder(folderPath string) error {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucketName),
		Prefix: aws.String(folderPath),
	}

	result, err := s.client.ListObjectsV2(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to list objects: %v", err)
	}

	for _, item := range result.Contents {
		s.wg.Add(1)

		go func(key *string) {
			defer s.wg.Done()
			_, err := s.client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
				Bucket: aws.String(s.bucketName),
				Key:    key,
			})
			if err != nil {
				log.Printf("failed to delete object %s: %v", *key, err)
			} else {
				log.Printf("successfully deleted object %s", *key)
			}
		}(item.Key)
	}

	s.wg.Wait()

	return nil
}
