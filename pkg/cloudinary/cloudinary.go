package cloudinary

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService struct {
	Cloudinary *cloudinary.Cloudinary
}

func NewCloudinaryService() *CloudinaryService {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatalf("❌ Gagal menginisialisasi Cloudinary: %v", err)
	}

	fmt.Println("✅ Cloudinary berhasil dikonfigurasi!")
	return &CloudinaryService{Cloudinary: cld}
}

func (cld *CloudinaryService) UploadImage(file io.Reader, filename string) (string, string, error) {
	ctx := context.Background()

	resp, err := cld.Cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: filename,
		Folder:   "be-catatain", // Changed folder structure
	})
	if err != nil {
		return "", "", fmt.Errorf("❌ Gagal upload gambar ke Cloudinary: %v", err)
	}

	return resp.SecureURL, resp.PublicID, nil
}

func (cld *CloudinaryService) DeleteImage(publicID string) error {
	ctx := context.Background()
	_, err := cld.Cloudinary.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return fmt.Errorf("failed to delete image: %v", err)
	}
	return nil
}

// Utility to parse the Cloudinary Image URL and extract PublicID to use for deletion
func (cld *CloudinaryService) GetPublicIDFromURL(imageURL string) string {
	parsedURL, err := url.Parse(imageURL)
	if err != nil {
		fmt.Println("❌ Gagal parsing URL:", err)
		return ""
	}
	path := parsedURL.Path
	segments := strings.Split(path, "/")

	if len(segments) > 0 {
		filenameWithExt := segments[len(segments)-1]
		publicID := strings.TrimSuffix(filenameWithExt, filepath.Ext(filenameWithExt))
		if len(segments) > 2 {
			// Include folder name (assuming 1 depth)
			folder := segments[len(segments)-2]
			publicID = folder + "/" + publicID
		}
		return publicID
	}

	return ""
}
