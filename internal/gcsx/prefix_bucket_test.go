// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gcsx_test

import (
	"io/ioutil"
	"testing"

	"golang.org/x/net/context"

	"github.com/googlecloudplatform/gcsfuse/internal/gcsx"
	"github.com/jacobsa/gcloud/gcs"
	"github.com/jacobsa/gcloud/gcs/gcsfake"
	"github.com/jacobsa/gcloud/gcs/gcsutil"
	. "github.com/jacobsa/ogletest"
	"github.com/jacobsa/timeutil"
)

func TestPrefixBucket(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Boilerplate
////////////////////////////////////////////////////////////////////////

type PrefixBucketTest struct {
	ctx     context.Context
	prefix  string
	wrapped gcs.Bucket
	bucket  gcs.Bucket
}

var _ SetUpInterface = &PrefixBucketTest{}

func init() { RegisterTestSuite(&PrefixBucketTest{}) }

func (t *PrefixBucketTest) SetUp(ti *TestInfo) {
	var err error

	t.ctx = ti.Ctx
	t.prefix = "foo_"
	t.wrapped = gcsfake.NewFakeBucket(timeutil.RealClock(), "some_bucket")

	t.bucket, err = gcsx.NewPrefixBucket(t.prefix, t.wrapped)
	AssertEq(nil, err)
}

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *PrefixBucketTest) Name() {
	ExpectEq(t.wrapped.Name(), t.bucket.Name())
}

func (t *PrefixBucketTest) NewReader() {
	var err error
	suffix := "taco"
	name := t.prefix + suffix
	contents := "foobar"

	// Create an object through the back door.
	_, err = gcsutil.CreateObject(t.ctx, t.wrapped, name, []byte(contents))
	AssertEq(nil, err)

	// Read it through the prefix bucket.
	rc, err := t.bucket.NewReader(
		t.ctx,
		&gcs.ReadObjectRequest{
			Name: suffix,
		})

	AssertEq(nil, err)
	defer rc.Close()

	actual, err := ioutil.ReadAll(rc)
	AssertEq(nil, err)
	ExpectEq(contents, string(actual))
}

func (t *PrefixBucketTest) CreateObject() {
	AddFailure("TODO")
}

func (t *PrefixBucketTest) CopyObject() {
	AddFailure("TODO")
}

func (t *PrefixBucketTest) ComposeObject() {
	AddFailure("TODO")
}

func (t *PrefixBucketTest) StatObject() {
	AddFailure("TODO")
}

func (t *PrefixBucketTest) ListObjects() {
	AddFailure("TODO")
}

func (t *PrefixBucketTest) UpdateObject() {
	AddFailure("TODO")
}

func (t *PrefixBucketTest) DeleteObject() {
	AddFailure("TODO")
}
