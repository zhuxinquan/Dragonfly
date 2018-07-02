/*
 * Copyright 1999-2018 Alibaba Group.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/go-check/check"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type TypesSuite struct{}

func init() {
	check.Suite(&TypesSuite{})
}

func (suite *TypesSuite) SetUpTest(c *check.C) {
	rand.Seed(time.Now().UnixNano())
}

// Testing BaseResponse

func (suite *TypesSuite) TestNewBaseResponse(c *check.C) {
	code := rand.Intn(100)
	msg := strconv.Itoa(rand.Int())
	res := NewBaseResponse(code, msg)
	c.Assert(res.Code, check.Equals, code)
	c.Assert(res.Msg, check.Equals, msg)
}

func (suite *TypesSuite) TestBaseResponse_IsSuccess(c *check.C) {
	var cases = []struct {
		code     int
		expected bool
	}{
		// [1]
		{1, true},
		// [2, n)
		{rand.Intn(10000) + 2, false},
		// (-n, 0]
		{-rand.Intn(10000), false},
	}

	var res *BaseResponse
	for _, cc := range cases {
		res = NewBaseResponse(cc.code, "")
		c.Assert(res.IsSuccess(), check.Equals, cc.expected)
	}
}
