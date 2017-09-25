// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package sqlstore

import (
	"testing"

	"github.com/mattermost/mattermost-server/model"
"github.com/mattermost/mattermost-server/store"
)

func TestLicenseStoreSave(t *testing.T) {
	ss := Setup()

	l1 := model.LicenseRecord{}
	l1.Id = model.NewId()
	l1.Bytes = "junk"

	if err := (<-ss.License().Save(&l1)).Err; err != nil {
		t.Fatal("couldn't save license record", err)
	}

	if err := (<-ss.License().Save(&l1)).Err; err != nil {
		t.Fatal("shouldn't fail on trying to save existing license record", err)
	}

	l1.Id = ""

	if err := (<-ss.License().Save(&l1)).Err; err == nil {
		t.Fatal("should fail on invalid license", err)
	}
}

func TestLicenseStoreGet(t *testing.T) {
	ss := Setup()

	l1 := model.LicenseRecord{}
	l1.Id = model.NewId()
	l1.Bytes = "junk"

	store.Must(ss.License().Save(&l1))

	if r := <-ss.License().Get(l1.Id); r.Err != nil {
		t.Fatal("couldn't get license", r.Err)
	} else {
		if r.Data.(*model.LicenseRecord).Bytes != l1.Bytes {
			t.Fatal("license bytes didn't match")
		}
	}

	if err := (<-ss.License().Get("missing")).Err; err == nil {
		t.Fatal("should fail on get license", err)
	}
}
