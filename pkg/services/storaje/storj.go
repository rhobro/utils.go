/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: utils.go
 * File Name: storj.go
 * Last Modified: 24/05/2021, 12:11
 */

package storaje

import (
	"context"
	"io"
	"storj.io/uplink"
)

var C *uplink.Project

func Init(satelliteAddr, apiKey string) error {
	// open access grant
	access, err := uplink.RequestAccessWithPassphrase(context.Background(), satelliteAddr, apiKey, "")
	if err != nil {
		return err
	}

	// open project
	C, err = uplink.OpenProject(context.Background(), access)
	if err != nil {
		return err
	}
	return nil
}

func ListBuckets() *uplink.BucketIterator {
	return C.ListBuckets(context.Background(), nil)
}

func CreateBucket(name string) error {
	_, err := C.CreateBucket(context.Background(), name)
	return err
}

func EnsureBucket(name string) error {
	_, err := C.EnsureBucket(context.Background(), name)
	return err
}

func Download(bucket, id string) (io.Reader, error) {
	return DownloadWithOptions(bucket, id, nil)
}

func DownloadWithOptions(bucket, id string, opt *uplink.DownloadOptions) (io.Reader, error) {
	if C == nil {
		panic("Storj client is nil")
	}

	dwl, err := C.DownloadObject(context.Background(), bucket, id, opt)
	return dwl, err
}

func Upload(rd io.Reader, bucket, id string) error {
	return UploadWithOptions(rd, bucket, id, nil)
}

func UploadWithOptions(rd io.Reader, bucket, id string, opt *uplink.UploadOptions) error {
	if C == nil {
		panic("Storj client is nil")
	}

	upl, err := C.UploadObject(context.Background(), bucket, id, opt)
	if err != nil {
		return err
	}

	// write data
	_, err = io.Copy(upl, rd)
	if err != nil {
		return err
	}
	err = upl.Commit()
	if err != nil {
		return err
	}

	return nil
}

func ListObjects(bucket string) *uplink.ObjectIterator {
	return ListObjectsWithOptions(bucket, nil)
}

func ListObjectsWithOptions(bucket string, opt *uplink.ListObjectsOptions) *uplink.ObjectIterator {
	return C.ListObjects(context.Background(), "videos", opt)
}

func DeleteObject(bucket, key string) (*uplink.Object, error) {
	return C.DeleteObject(context.Background(), "videos", "composite")
}
