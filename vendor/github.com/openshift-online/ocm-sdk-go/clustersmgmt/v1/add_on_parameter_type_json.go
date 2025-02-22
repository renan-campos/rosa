/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalAddOnParameter writes a value of the 'add_on_parameter' type to the given writer.
func MarshalAddOnParameter(object *AddOnParameter, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeAddOnParameter(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// writeAddOnParameter writes a value of the 'add_on_parameter' type to the given stream.
func writeAddOnParameter(object *AddOnParameter, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	stream.WriteObjectField("kind")
	if object.bitmap_&1 != 0 {
		stream.WriteString(AddOnParameterLinkKind)
	} else {
		stream.WriteString(AddOnParameterKind)
	}
	count++
	if object.bitmap_&2 != 0 {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(object.id)
		count++
	}
	if object.bitmap_&4 != 0 {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(object.href)
		count++
	}
	var present_ bool
	present_ = object.bitmap_&8 != 0 && object.addon != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("addon")
		writeAddOn(object.addon, stream)
		count++
	}
	present_ = object.bitmap_&16 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("default_value")
		stream.WriteString(object.defaultValue)
		count++
	}
	present_ = object.bitmap_&32 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("description")
		stream.WriteString(object.description)
		count++
	}
	present_ = object.bitmap_&64 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("editable")
		stream.WriteBool(object.editable)
		count++
	}
	present_ = object.bitmap_&128 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("enabled")
		stream.WriteBool(object.enabled)
		count++
	}
	present_ = object.bitmap_&256 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("name")
		stream.WriteString(object.name)
		count++
	}
	present_ = object.bitmap_&512 != 0 && object.options != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("options")
		writeAddOnParameterOptionList(object.options, stream)
		count++
	}
	present_ = object.bitmap_&1024 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("required")
		stream.WriteBool(object.required)
		count++
	}
	present_ = object.bitmap_&2048 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("validation")
		stream.WriteString(object.validation)
		count++
	}
	present_ = object.bitmap_&4096 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("validation_err_msg")
		stream.WriteString(object.validationErrMsg)
		count++
	}
	present_ = object.bitmap_&8192 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("value_type")
		stream.WriteString(object.valueType)
	}
	stream.WriteObjectEnd()
}

// UnmarshalAddOnParameter reads a value of the 'add_on_parameter' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalAddOnParameter(source interface{}) (object *AddOnParameter, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readAddOnParameter(iterator)
	err = iterator.Error
	return
}

// readAddOnParameter reads a value of the 'add_on_parameter' type from the given iterator.
func readAddOnParameter(iterator *jsoniter.Iterator) *AddOnParameter {
	object := &AddOnParameter{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			if value == AddOnParameterLinkKind {
				object.bitmap_ |= 1
			}
		case "id":
			object.id = iterator.ReadString()
			object.bitmap_ |= 2
		case "href":
			object.href = iterator.ReadString()
			object.bitmap_ |= 4
		case "addon":
			value := readAddOn(iterator)
			object.addon = value
			object.bitmap_ |= 8
		case "default_value":
			value := iterator.ReadString()
			object.defaultValue = value
			object.bitmap_ |= 16
		case "description":
			value := iterator.ReadString()
			object.description = value
			object.bitmap_ |= 32
		case "editable":
			value := iterator.ReadBool()
			object.editable = value
			object.bitmap_ |= 64
		case "enabled":
			value := iterator.ReadBool()
			object.enabled = value
			object.bitmap_ |= 128
		case "name":
			value := iterator.ReadString()
			object.name = value
			object.bitmap_ |= 256
		case "options":
			value := readAddOnParameterOptionList(iterator)
			object.options = value
			object.bitmap_ |= 512
		case "required":
			value := iterator.ReadBool()
			object.required = value
			object.bitmap_ |= 1024
		case "validation":
			value := iterator.ReadString()
			object.validation = value
			object.bitmap_ |= 2048
		case "validation_err_msg":
			value := iterator.ReadString()
			object.validationErrMsg = value
			object.bitmap_ |= 4096
		case "value_type":
			value := iterator.ReadString()
			object.valueType = value
			object.bitmap_ |= 8192
		default:
			iterator.ReadAny()
		}
	}
	return object
}
