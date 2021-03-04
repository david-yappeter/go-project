package service

import (
	"context"
	"fmt"
)

//GdriveViewLink View Link
func GdriveViewLink(ctx context.Context, fileID string) (string, error) {
	return fmt.Sprintf("https://drive.google.com/uc?export=view&id=%s", fileID), nil
}

//GdriveDownloadLink Download Link
func GdriveDownloadLink(ctx context.Context, fileID string) (string, error) {
	return fmt.Sprintf("https://drive.google.com/uc?export=download&id=%s", fileID), nil
}
