package openapi3

import (
	"encoding/json"
	"github.com/jban332/kin-openapi/jsoninfo"
)

// ExtensionProps provides support for OpenAPI extensions.
// It reads/writes all properties that begin with "x-".
type ExtensionProps struct {
	Extensions map[string]interface{} `json:"-"`
}

// Assert that the type implements the interface
var _ jsoninfo.StrictStruct = &ExtensionProps{}

// EncodeWith will be invoked by package "jsoninfo"
func (props *ExtensionProps) EncodeWith(encoder *jsoninfo.ObjectEncoder, value interface{}) error {
	if m := props.Extensions; m != nil {
		for k, v := range m {
			err := encoder.EncodeExtension(k, v)
			if err != nil {
				return err
			}
		}
	}
	return encoder.EncodeStructFieldsAndExtensions(value)
}

// DecodeWith will be invoked by package "jsoninfo"
func (props *ExtensionProps) DecodeWith(decoder *jsoninfo.ObjectDecoder, value interface{}) error {
	source := decoder.DecodeExtensionMap()
	if source != nil && len(source) > 0 {
		result := make(map[string]interface{}, len(source))
		for k, v := range source {
			result[k] = json.RawMessage(v)
		}
		props.Extensions = result
	}
	return decoder.DecodeStructFieldsAndExtensions(value)
}
