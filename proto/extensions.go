// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2010 Google Inc.  All rights reserved.
// http://code.google.com/p/goprotobuf/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package proto


/*
 * Types and routines for supporting protocol buffer extensions.
 */

import (
	"os"
	"reflect"
	"unsafe"
)

// ExtensionRange represents a range of message extensions for a protocol buffer.
// Used in code generated by the protocol compiler.
type ExtensionRange struct {
	Start, End int32 // both inclusive
}

// extendableProto is an interface implemented by any protocol buffer that may be extended.
type extendableProto interface {
	ExtensionRangeArray() []ExtensionRange
	ExtensionMap() map[int32][]byte
}

// ExtensionDesc represents an extension specification.
// Used in generated code from the protocol compiler.
type ExtensionDesc struct {
	ExtendedType  interface{} // nil pointer to the type that is being extended
	ExtensionType interface{} // nil pointer to the extension type
	Field         int32       // field number
	Tag           string      // PB(...) tag style
}

// isExtensionField returns true iff the given field number is in an extension range.
func isExtensionField(pb extendableProto, field int32) bool {
	for _, er := range pb.ExtensionRangeArray() {
		if er.Start <= field && field <= er.End {
			return true
		}
	}
	return false
}

// checkExtensionTypes checks that the given extension is valid for pb.
func checkExtensionTypes(pb extendableProto, extension *ExtensionDesc) os.Error {
	// Check the extended type.
	if a, b := reflect.Typeof(pb), reflect.Typeof(extension.ExtendedType); a != b {
		return os.NewError("bad extended type; " + b.String() + " does not extend " + a.String())
	}
	// Check the range.
	if !isExtensionField(pb, extension.Field) {
		return os.NewError("bad extension number; not in declared ranges")
	}
	return nil
}

// HasExtension returns whether the given extension is present in pb.
func HasExtension(pb extendableProto, extension *ExtensionDesc) bool {
	// TODO: Check types, field numbers, etc.?
	_, ok := pb.ExtensionMap()[extension.Field]
	return ok
}

// ClearExtension removes the given extension from pb.
func ClearExtension(pb extendableProto, extension *ExtensionDesc) {
	// TODO: Check types, field numbers, etc.?
	pb.ExtensionMap()[extension.Field] = nil, false
}

// GetExtension parses and returns the given extension of pb.
// If the extension is not present it returns (nil, nil).
func GetExtension(pb extendableProto, extension *ExtensionDesc) (interface{}, os.Error) {
	if err := checkExtensionTypes(pb, extension); err != nil {
		return nil, err
	}

	b, ok := pb.ExtensionMap()[extension.Field]
	if !ok {
		return nil, nil // not an error
	}

	// Discard wire type and field number varint. It isn't needed.
	_, n := DecodeVarint(b)
	o := NewBuffer(b[n:])

	t := reflect.Typeof(extension.ExtensionType).(*reflect.PtrType)
	props := &Properties{}
	props.Init(t, "irrelevant_name", extension.Tag, 0)

	base := unsafe.New(t)
	var sbase uintptr
	if _, ok := t.Elem().(*reflect.StructType); ok {
		// props.dec will be dec_struct_message, which does not refer to sbase.
		*(*unsafe.Pointer)(base) = unsafe.New(t.Elem())
	} else {
		sbase = uintptr(unsafe.New(t.Elem()))
	}
	if err := props.dec(o, props, uintptr(base), sbase); err != nil {
		return nil, err
	}
	return unsafe.Unreflect(t, base), nil
}

// GetExtensions returns a slice of the extensions present in pb that are also listed in es.
// The returned slice has the same length as es; missing extensions will appear as nil elements.
func GetExtensions(pb interface{}, es []*ExtensionDesc) (extensions []interface{}, err os.Error) {
	epb, ok := pb.(extendableProto)
	if !ok {
		err = os.NewError("not an extendable proto")
		return
	}
	extensions = make([]interface{}, len(es))
	for i, e := range es {
		extensions[i], err = GetExtension(epb, e)
		if err != nil {
			return
		}
	}
	return
}

// TODO: (needed for repeated extensions)
//   - ExtensionSize
//   - AddExtension

// SetExtension sets the specified extension of pb to the specified value.
func SetExtension(pb extendableProto, extension *ExtensionDesc, value interface{}) os.Error {
	if err := checkExtensionTypes(pb, extension); err != nil {
		return err
	}
	if reflect.Typeof(extension.ExtensionType) != reflect.Typeof(value) {
		return os.NewError("bad extension value type")
	}

	props := new(Properties)
	props.Init(reflect.Typeof(extension.ExtensionType), "unknown_name", extension.Tag, 0)

	p := NewBuffer(nil)
	v := reflect.NewValue(value)
	if err := props.enc(p, props, v.UnsafeAddr()); err != nil {
		return err
	}
	pb.ExtensionMap()[extension.Field] = p.buf
	return nil
}
