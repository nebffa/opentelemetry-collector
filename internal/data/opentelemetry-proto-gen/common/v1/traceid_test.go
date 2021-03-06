// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTraceID(t *testing.T) {
	tid := NewTraceID(nil)
	assert.EqualValues(t, []byte(nil), tid.id)
	assert.EqualValues(t, 0, tid.Size())

	b := []byte{1, 2, 3}
	tid = NewTraceID(b)
	assert.EqualValues(t, b, tid.id)
	assert.EqualValues(t, b, tid.Bytes())
	assert.EqualValues(t, 3, tid.Size())
}

func TestTraceIDHexString(t *testing.T) {
	tid := NewTraceID(nil)
	assert.EqualValues(t, "", tid.HexString())

	tid = NewTraceID([]byte{})
	assert.EqualValues(t, "", tid.HexString())

	tid = NewTraceID([]byte{0x12, 0x23, 0xAD})
	assert.EqualValues(t, "1223ad", tid.HexString())
}

func TestTraceIDEqual(t *testing.T) {
	tid := NewTraceID(nil)
	assert.True(t, tid.Equal(tid))
	assert.True(t, tid.Equal(NewTraceID(nil)))
	assert.True(t, tid.Equal(NewTraceID([]byte{})))
	assert.False(t, tid.Equal(NewTraceID([]byte{1})))

	tid = NewTraceID([]byte{0x12, 0x23, 0xAD})
	assert.False(t, tid.Equal(NewTraceID(nil)))
	assert.True(t, tid.Equal(tid))
	assert.True(t, tid.Equal(NewTraceID([]byte{0x12, 0x23, 0xAD})))
}

func TestTraceIDCompare(t *testing.T) {
	tid := NewTraceID(nil)
	assert.EqualValues(t, 0, tid.Compare(tid))
	assert.EqualValues(t, 0, tid.Compare(NewTraceID(nil)))
	assert.EqualValues(t, 0, tid.Compare(NewTraceID([]byte{})))
	assert.EqualValues(t, -1, tid.Compare(NewTraceID([]byte{1})))

	tid = NewTraceID([]byte{0x12, 0x23, 0xAD})
	assert.EqualValues(t, 1, tid.Compare(NewTraceID(nil)))
	assert.EqualValues(t, 0, tid.Compare(tid))
	assert.EqualValues(t, 0, tid.Compare(NewTraceID([]byte{0x12, 0x23, 0xAD})))
	assert.EqualValues(t, 1, tid.Compare(NewTraceID([]byte{0x12, 0x23, 0xAC})))
	assert.EqualValues(t, -1, tid.Compare(NewTraceID([]byte{0x12, 0x23, 0xAE})))
	assert.EqualValues(t, 1, tid.Compare(NewTraceID([]byte{0x12, 0x23})))
	assert.EqualValues(t, -1, tid.Compare(NewTraceID([]byte{0x12, 0x24})))
}

func TestTraceIDMarshal(t *testing.T) {
	buf := make([]byte, 10)

	tid := NewTraceID(nil)
	n, err := tid.MarshalTo(buf)
	assert.EqualValues(t, 0, n)
	assert.NoError(t, err)

	tid = NewTraceID([]byte{0x12, 0x23, 0xAD})
	n, err = tid.MarshalTo(buf)
	assert.EqualValues(t, 3, n)
	assert.EqualValues(t, []byte{0x12, 0x23, 0xAD}, buf[0:3])
	assert.NoError(t, err)

	_, err = tid.MarshalTo(buf[0:1])
	assert.Error(t, err)
}

func TestTraceIDMarshalJSON(t *testing.T) {
	tid := NewTraceID(nil)
	json, err := tid.MarshalJSON()
	assert.EqualValues(t, []byte(`""`), json)
	assert.NoError(t, err)

	tid = NewTraceID([]byte{0x12, 0x23, 0xAD})
	json, err = tid.MarshalJSON()
	assert.EqualValues(t, []byte(`"1223ad"`), json)
	assert.NoError(t, err)
}

func TestTraceIDUnmarshal(t *testing.T) {
	buf := []byte{0x12, 0x23, 0xAD}

	tid := TraceID{}
	err := tid.Unmarshal(buf[0:3])
	assert.NoError(t, err)
	assert.EqualValues(t, []byte{0x12, 0x23, 0xAD}, tid.id)

	err = tid.Unmarshal(buf[0:0])
	assert.NoError(t, err)
	assert.EqualValues(t, []byte{}, tid.id)

	tid = TraceID{}
	err = tid.Unmarshal(nil)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(nil), tid.id)
}

func TestTraceIDUnmarshalJSON(t *testing.T) {
	tid := TraceID{}
	err := tid.UnmarshalJSON([]byte(`""`))
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(nil), tid.id)

	err = tid.UnmarshalJSON([]byte(`"1234"`))
	assert.NoError(t, err)
	assert.EqualValues(t, []byte{0x12, 0x34}, tid.id)

	err = tid.UnmarshalJSON([]byte(`1234`))
	assert.Error(t, err)

	err = tid.UnmarshalJSON([]byte(`"nothex"`))
	assert.Error(t, err)

	err = tid.UnmarshalJSON([]byte(`"1"`))
	assert.Error(t, err)

	err = tid.UnmarshalJSON([]byte(`"123"`))
	assert.Error(t, err)

	err = tid.UnmarshalJSON([]byte(`"`))
	assert.Error(t, err)
}
