package trails

import (
	"cloud.google.com/go/storage"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
	"io"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

var (
	storageClient *storage.Client
)

var bucket = "your bucket name"
var urlName = "https://storage.googleapis.com"

// HandleFileUploadToBucket uploads a file to a bucket
func HandleFileUploadToBucket(c *fiber.Ctx, folder *string) (*string, error) {
	var err error
	ctx := c.Context()
	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("./google_could.json"))
	if err != nil {
		return nil, err
	}

	file, err := c.FormFile("image")
	if err != nil {
		return nil, err
	}

	// Set the maximum allowed size to 5MB (5 * 1024 * 1024 bytes)
	maxSize := int64(5 * 1024 * 1024)
	if file.Size > maxSize {
		return nil, errors.New("file size exceeds the maximum allowed size")
	}
	// Validate the file type (assuming you want to allow only specific image types)
	validImageTypes := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		// Add more allowed types here
	}
	ext := filepath.Ext(file.Filename)
	if !validImageTypes[ext] {
		return nil, errors.New("invalid image type")
	}

	fileHeader, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer fileHeader.Close()

	var objectName string

	if folder != nil {
		objectName = *folder + "/" + time.Now().Format("20060102150405") + ext
	} else {
		objectName = time.Now().Format("20060102150405") + "_" + file.Filename
	}

	sw := storageClient.Bucket(bucket).Object(objectName).NewWriter(ctx)

	if _, errCopy := io.Copy(sw, fileHeader); errCopy != nil {
		return nil, errCopy
	}

	if errSw := sw.Close(); errSw != nil {
		return nil, errSw
	}

	u, err := url.Parse("/" + bucket + "/" + objectName)
	if err != nil {
		return nil, err
	}

	link := urlName + u.EscapedPath()
	fmt.Println(link)
	return &link, nil

}

func HandleDeleteImage(c *fiber.Ctx, imageName string) error {
	ctx := c.Context()
	storageClientD, err := storage.NewClient(ctx, option.WithCredentialsFile("./google_could.json"))
	if err != nil {
		return err
	}
	name := strings.Replace(imageName, urlName+"/"+bucket+"/", "", -1)
	obj := storageClientD.Bucket(bucket).Object(name)
	if errDelete := obj.Delete(ctx); errDelete != nil {
		return errDelete

	}
	return nil
}
