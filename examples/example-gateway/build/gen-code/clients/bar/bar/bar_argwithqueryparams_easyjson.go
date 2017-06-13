// Code generated by zanzibar
// @generated
// Checksum : jtPtilDuw3gWvZltH/NjOw==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonCec46174DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithQueryParams(in *jlexer.Lexer, out *Bar_ArgWithQueryParams_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(BarResponse)
				}
				(*out.Success).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCec46174EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithQueryParams(out *jwriter.Writer, in Bar_ArgWithQueryParams_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"success\":")
		if in.Success == nil {
			out.RawString("null")
		} else {
			(*in.Success).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_ArgWithQueryParams_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCec46174EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithQueryParams(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_ArgWithQueryParams_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCec46174EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithQueryParams(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_ArgWithQueryParams_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCec46174DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithQueryParams(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_ArgWithQueryParams_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCec46174DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithQueryParams(l, v)
}
func easyjsonCec46174DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithQueryParams1(in *jlexer.Lexer, out *Bar_ArgWithQueryParams_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var NameSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
			NameSet = true
		case "userUUID":
			if in.IsNull() {
				in.Skip()
				out.UserUUID = nil
			} else {
				if out.UserUUID == nil {
					out.UserUUID = new(string)
				}
				*out.UserUUID = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !NameSet {
		in.AddError(fmt.Errorf("key 'name' is required"))
	}
}
func easyjsonCec46174EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithQueryParams1(out *jwriter.Writer, in Bar_ArgWithQueryParams_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if in.UserUUID != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"userUUID\":")
		if in.UserUUID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.UserUUID))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_ArgWithQueryParams_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCec46174EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithQueryParams1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_ArgWithQueryParams_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCec46174EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithQueryParams1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_ArgWithQueryParams_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCec46174DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithQueryParams1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_ArgWithQueryParams_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCec46174DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgWithQueryParams1(l, v)
}
