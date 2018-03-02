// Copyright (c) 2018-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package utils

import (
	"testing"

	"github.com/mattermost/mattermost-server/model"
)

func TestCheckMandatoryS3Fields(t *testing.T) {
	cfg := model.FileSettings{}

	err := CheckMandatoryS3Fields(&cfg)
	if err == nil || err.Message != "api.admin.test_s3.missing_s3_bucket" {
		t.Fatal("should've failed with missing s3 bucket")
	}

	cfg.AmazonS3Bucket = "test-mm"
	err = CheckMandatoryS3Fields(&cfg)
	if err == nil || err.Message != "api.admin.test_s3.missing_s3_endpoint" {
		t.Fatal("should've failed with missing s3 endpoint")
	}

	cfg.AmazonS3Endpoint = "s3.newendpoint.com"
	err = CheckMandatoryS3Fields(&cfg)
	if err == nil || err.Message != "api.admin.test_s3.missing_s3_region" {
		t.Fatal("should've failed with missing s3 region")
	}

}
