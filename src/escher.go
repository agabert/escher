package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var sourceDirectory = "/space"
	var destinationDirectory = "/space3/escher/trunk"

	if _, isDebug := os.LookupEnv("DEBUG"); isDebug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug logging enabled.")
	}

	logrus.Info("main(): starting")

	var fileWalker = func(fullName string, fileInfo os.FileInfo, fileError error) error {
		if fileError != nil {
			logrus.Fatal(fileError)
			return fileError
		}

		if fileInfo.IsDir() {
			logrus.Debug("Processing Directory: ", fullName)
		} else {
			/*
			 * (I) check if the file is a normal file
			 */
			sourceFileStat, err := os.Stat(fullName)
			if err != nil {
				return nil
			}
			if !sourceFileStat.Mode().IsRegular() {
				return nil
			}
			sourceLinkStat, err := os.Lstat(fullName)
			if sourceLinkStat.Mode()&os.ModeSymlink != 0 {
				return nil
			}

			/*
			 * (II) check if the file has an extension that we are looking for
			 */
			extension := strings.TrimPrefix(strings.ToLower(filepath.Ext(fullName)), ".")
			if extension == "" {
				return nil
			}
			if extension == "jpeg" {
				extension = "jpg"
			}
			if extension == "tiff" {
				extension = "tif"
			}
			if filterExtension(extension) {
				return nil
			}
			if filterPath(fullName) {
				return nil
			}

			/*
			 * (III) check for minimum and maximum file size
			 */
			fhandle, err := os.Stat(fullName)
			if err != nil {
				return nil
			}
			if fhandle.Size() < 20000 {
				return nil
			}

			if fhandle.Size() > 2000000000 {
				return nil
			}

			/*
			 * (IV) prepare the fingerprint by reading the file once
			 */
			fingerPrint := getFileChecksum(fullName)
			if fingerPrint == "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" {
				// empty file
				return nil
			}
			if fingerPrint == "a8b7b3b3de85c0309fa72d15a91876e8b4638e8c3412bc783d94cf422025917d" {
				// NSFW ;)
				return nil
			}

			/*
			 * (V) import the file to the trunk directory based on fingerprint
			 */
			logrus.Debug("Importing: ", fullName, " (", fingerPrint, ")")
			dstFile := fmt.Sprintf("%s/%s.%s", destinationDirectory, fingerPrint, extension)
			sourceFileHandle, err := os.Open(fullName)
			if err == nil {
				destinationFileHandle, err := os.Create(dstFile)
				if err == nil {
					logrus.Info("Creating: ", dstFile, " [", fullName, "]")
					io.Copy(destinationFileHandle, sourceFileHandle)
					destinationFileHandle.Close()
					sourceFileHandle.Close()
				}
			}
		}

		return nil
	}

	if err := filepath.Walk(sourceDirectory, fileWalker); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("main(): exiting")

	os.Exit(0)
}
