package service

import (
	"crypto/sha1"
	"douyin/app/config"
	"douyin/app/dao"
	"douyin/pkg/iox"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"mime"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	hashFn    = sha1.New
	coverExt  = ".jpg"
	coverMIME = mime.TypeByExtension(coverExt)
)

func init() {
	// check ffmpeg installed correctly
	cmd := exec.Command("ffmpeg", "-version")
	if out, err := cmd.CombinedOutput(); err != nil {
		panic("ffmpeg isn't installed: " + string(out))
	}
}

func PublishVideo(userId int64, title string, file *multipart.FileHeader) error {
	videoExt, err := iox.GetExtension(file)
	if err != nil {
		return err
	}
	if videoExt == "" {
		return errors.New("failed to determine MIME")
	}
	videoMIME := mime.TypeByExtension(videoExt)

	storagePath := config.Val.Static.Filepath

	// don't concat videoExt
	videoFileNoExt := fmt.Sprintf("%d_%s", userId, generateUUID())
	// concat videoExt here
	videoPath := filepath.Join(storagePath, videoFileNoExt+videoExt)

	h, err := iox.HashAndSaveFile(hashFn, file, videoPath)
	if err != nil {
		clearFiles(videoPath)
		return err
	}

	coverFilename := fmt.Sprintf("%s_cover%s", videoFileNoExt, coverExt)
	coverFilePath := filepath.Join(storagePath, coverFilename)
	if err := extractCoverFrame(videoPath, coverFilePath); err != nil {
		clearFiles(videoPath, coverFilePath)
		return err
	}

	coverHash, err := iox.HashFile(hashFn, coverFilePath)
	if err != nil {
		clearFiles(videoPath, coverFilePath)
		return err
	}

	mediaFile := dao.MediaFile{
		Key:  videoFileNoExt + videoExt,
		MIME: videoMIME,
		SHA1: hex.EncodeToString(*h),
	}
	coverFile := dao.MediaFile{
		Key:  coverFilename,
		MIME: coverMIME,
		SHA1: hex.EncodeToString(*coverHash),
	}
	err = dao.SaveVideoFile(userId, title, &mediaFile, &coverFile)
	if err != nil {
		clearFiles(videoPath, coverFilePath)
		return err
	}
	return nil
}

func generateUUID() string {
	return strings.Replace(uuid.NewString(), "-", "", 4)
}

func clearFiles(files ...string) {
	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			// file exists
			// remove it
			_ = os.Remove(file)
		}
	}
}

func extractCoverFrame(videoPath string, coverPath string) error {
	// capture the first frame
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vframes", "1", coverPath)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
