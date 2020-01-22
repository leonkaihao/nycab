// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gogo.proto

package gogoproto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

var E_GoprotoEnumPrefix = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62001,
	Name:          "gogoproto.goproto_enum_prefix",
	Tag:           "varint,62001,opt,name=goproto_enum_prefix",
	Filename:      "gogo.proto",
}

var E_GoprotoEnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62021,
	Name:          "gogoproto.goproto_enum_stringer",
	Tag:           "varint,62021,opt,name=goproto_enum_stringer",
	Filename:      "gogo.proto",
}

var E_EnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62022,
	Name:          "gogoproto.enum_stringer",
	Tag:           "varint,62022,opt,name=enum_stringer",
	Filename:      "gogo.proto",
}

var E_EnumCustomname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         62023,
	Name:          "gogoproto.enum_customname",
	Tag:           "bytes,62023,opt,name=enum_customname",
	Filename:      "gogo.proto",
}

var E_Enumdecl = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62024,
	Name:          "gogoproto.enumdecl",
	Tag:           "varint,62024,opt,name=enumdecl",
	Filename:      "gogo.proto",
}

var E_EnumvalueCustomname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         66001,
	Name:          "gogoproto.enumvalue_customname",
	Tag:           "bytes,66001,opt,name=enumvalue_customname",
	Filename:      "gogo.proto",
}

var E_GoprotoGettersAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63001,
	Name:          "gogoproto.goproto_getters_all",
	Tag:           "varint,63001,opt,name=goproto_getters_all",
	Filename:      "gogo.proto",
}

var E_GoprotoEnumPrefixAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63002,
	Name:          "gogoproto.goproto_enum_prefix_all",
	Tag:           "varint,63002,opt,name=goproto_enum_prefix_all",
	Filename:      "gogo.proto",
}

var E_GoprotoStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63003,
	Name:          "gogoproto.goproto_stringer_all",
	Tag:           "varint,63003,opt,name=goproto_stringer_all",
	Filename:      "gogo.proto",
}

var E_VerboseEqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63004,
	Name:          "gogoproto.verbose_equal_all",
	Tag:           "varint,63004,opt,name=verbose_equal_all",
	Filename:      "gogo.proto",
}

var E_FaceAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63005,
	Name:          "gogoproto.face_all",
	Tag:           "varint,63005,opt,name=face_all",
	Filename:      "gogo.proto",
}

var E_GostringAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63006,
	Name:          "gogoproto.gostring_all",
	Tag:           "varint,63006,opt,name=gostring_all",
	Filename:      "gogo.proto",
}

var E_PopulateAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63007,
	Name:          "gogoproto.populate_all",
	Tag:           "varint,63007,opt,name=populate_all",
	Filename:      "gogo.proto",
}

var E_StringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63008,
	Name:          "gogoproto.stringer_all",
	Tag:           "varint,63008,opt,name=stringer_all",
	Filename:      "gogo.proto",
}

var E_OnlyoneAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63009,
	Name:          "gogoproto.onlyone_all",
	Tag:           "varint,63009,opt,name=onlyone_all",
	Filename:      "gogo.proto",
}

var E_EqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63013,
	Name:          "gogoproto.equal_all",
	Tag:           "varint,63013,opt,name=equal_all",
	Filename:      "gogo.proto",
}

var E_DescriptionAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63014,
	Name:          "gogoproto.description_all",
	Tag:           "varint,63014,opt,name=description_all",
	Filename:      "gogo.proto",
}

var E_TestgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63015,
	Name:          "gogoproto.testgen_all",
	Tag:           "varint,63015,opt,name=testgen_all",
	Filename:      "gogo.proto",
}

var E_BenchgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63016,
	Name:          "gogoproto.benchgen_all",
	Tag:           "varint,63016,opt,name=benchgen_all",
	Filename:      "gogo.proto",
}

var E_MarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63017,
	Name:          "gogoproto.marshaler_all",
	Tag:           "varint,63017,opt,name=marshaler_all",
	Filename:      "gogo.proto",
}

var E_UnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63018,
	Name:          "gogoproto.unmarshaler_all",
	Tag:           "varint,63018,opt,name=unmarshaler_all",
	Filename:      "gogo.proto",
}

var E_StableMarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63019,
	Name:          "gogoproto.stable_marshaler_all",
	Tag:           "varint,63019,opt,name=stable_marshaler_all",
	Filename:      "gogo.proto",
}

var E_SizerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63020,
	Name:          "gogoproto.sizer_all",
	Tag:           "varint,63020,opt,name=sizer_all",
	Filename:      "gogo.proto",
}

var E_GoprotoEnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63021,
	Name:          "gogoproto.goproto_enum_stringer_all",
	Tag:           "varint,63021,opt,name=goproto_enum_stringer_all",
	Filename:      "gogo.proto",
}

var E_EnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63022,
	Name:          "gogoproto.enum_stringer_all",
	Tag:           "varint,63022,opt,name=enum_stringer_all",
	Filename:      "gogo.proto",
}

var E_UnsafeMarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63023,
	Name:          "gogoproto.unsafe_marshaler_all",
	Tag:           "varint,63023,opt,name=unsafe_marshaler_all",
	Filename:      "gogo.proto",
}

var E_UnsafeUnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63024,
	Name:          "gogoproto.unsafe_unmarshaler_all",
	Tag:           "varint,63024,opt,name=unsafe_unmarshaler_all",
	Filename:      "gogo.proto",
}

var E_GoprotoExtensionsMapAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63025,
	Name:          "gogoproto.goproto_extensions_map_all",
	Tag:           "varint,63025,opt,name=goproto_extensions_map_all",
	Filename:      "gogo.proto",
}

var E_GoprotoUnrecognizedAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63026,
	Name:          "gogoproto.goproto_unrecognized_all",
	Tag:           "varint,63026,opt,name=goproto_unrecognized_all",
	Filename:      "gogo.proto",
}

var E_GogoprotoImport = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63027,
	Name:          "gogoproto.gogoproto_import",
	Tag:           "varint,63027,opt,name=gogoproto_import",
	Filename:      "gogo.proto",
}

var E_ProtosizerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63028,
	Name:          "gogoproto.protosizer_all",
	Tag:           "varint,63028,opt,name=protosizer_all",
	Filename:      "gogo.proto",
}

var E_CompareAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63029,
	Name:          "gogoproto.compare_all",
	Tag:           "varint,63029,opt,name=compare_all",
	Filename:      "gogo.proto",
}

var E_TypedeclAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63030,
	Name:          "gogoproto.typedecl_all",
	Tag:           "varint,63030,opt,name=typedecl_all",
	Filename:      "gogo.proto",
}

var E_EnumdeclAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63031,
	Name:          "gogoproto.enumdecl_all",
	Tag:           "varint,63031,opt,name=enumdecl_all",
	Filename:      "gogo.proto",
}

var E_GoprotoRegistration = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63032,
	Name:          "gogoproto.goproto_registration",
	Tag:           "varint,63032,opt,name=goproto_registration",
	Filename:      "gogo.proto",
}

var E_MessagenameAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63033,
	Name:          "gogoproto.messagename_all",
	Tag:           "varint,63033,opt,name=messagename_all",
	Filename:      "gogo.proto",
}

var E_GoprotoSizecacheAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63034,
	Name:          "gogoproto.goproto_sizecache_all",
	Tag:           "varint,63034,opt,name=goproto_sizecache_all",
	Filename:      "gogo.proto",
}

var E_GoprotoUnkeyedAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63035,
	Name:          "gogoproto.goproto_unkeyed_all",
	Tag:           "varint,63035,opt,name=goproto_unkeyed_all",
	Filename:      "gogo.proto",
}

var E_GoprotoGetters = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64001,
	Name:          "gogoproto.goproto_getters",
	Tag:           "varint,64001,opt,name=goproto_getters",
	Filename:      "gogo.proto",
}

var E_GoprotoStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64003,
	Name:          "gogoproto.goproto_stringer",
	Tag:           "varint,64003,opt,name=goproto_stringer",
	Filename:      "gogo.proto",
}

var E_VerboseEqual = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64004,
	Name:          "gogoproto.verbose_equal",
	Tag:           "varint,64004,opt,name=verbose_equal",
	Filename:      "gogo.proto",
}

var E_Face = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64005,
	Name:          "gogoproto.face",
	Tag:           "varint,64005,opt,name=face",
	Filename:      "gogo.proto",
}

var E_Gostring = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64006,
	Name:          "gogoproto.gostring",
	Tag:           "varint,64006,opt,name=gostring",
	Filename:      "gogo.proto",
}

var E_Populate = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64007,
	Name:          "gogoproto.populate",
	Tag:           "varint,64007,opt,name=populate",
	Filename:      "gogo.proto",
}

var E_Stringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         67008,
	Name:          "gogoproto.stringer",
	Tag:           "varint,67008,opt,name=stringer",
	Filename:      "gogo.proto",
}

var E_Onlyone = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64009,
	Name:          "gogoproto.onlyone",
	Tag:           "varint,64009,opt,name=onlyone",
	Filename:      "gogo.proto",
}

var E_Equal = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64013,
	Name:          "gogoproto.equal",
	Tag:           "varint,64013,opt,name=equal",
	Filename:      "gogo.proto",
}

var E_Description = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64014,
	Name:          "gogoproto.description",
	Tag:           "varint,64014,opt,name=description",
	Filename:      "gogo.proto",
}

var E_Testgen = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64015,
	Name:          "gogoproto.testgen",
	Tag:           "varint,64015,opt,name=testgen",
	Filename:      "gogo.proto",
}

var E_Benchgen = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64016,
	Name:          "gogoproto.benchgen",
	Tag:           "varint,64016,opt,name=benchgen",
	Filename:      "gogo.proto",
}

var E_Marshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64017,
	Name:          "gogoproto.marshaler",
	Tag:           "varint,64017,opt,name=marshaler",
	Filename:      "gogo.proto",
}

var E_Unmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64018,
	Name:          "gogoproto.unmarshaler",
	Tag:           "varint,64018,opt,name=unmarshaler",
	Filename:      "gogo.proto",
}

var E_StableMarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64019,
	Name:          "gogoproto.stable_marshaler",
	Tag:           "varint,64019,opt,name=stable_marshaler",
	Filename:      "gogo.proto",
}

var E_Sizer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64020,
	Name:          "gogoproto.sizer",
	Tag:           "varint,64020,opt,name=sizer",
	Filename:      "gogo.proto",
}

var E_UnsafeMarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64023,
	Name:          "gogoproto.unsafe_marshaler",
	Tag:           "varint,64023,opt,name=unsafe_marshaler",
	Filename:      "gogo.proto",
}

var E_UnsafeUnmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64024,
	Name:          "gogoproto.unsafe_unmarshaler",
	Tag:           "varint,64024,opt,name=unsafe_unmarshaler",
	Filename:      "gogo.proto",
}

var E_GoprotoExtensionsMap = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64025,
	Name:          "gogoproto.goproto_extensions_map",
	Tag:           "varint,64025,opt,name=goproto_extensions_map",
	Filename:      "gogo.proto",
}

var E_GoprotoUnrecognized = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64026,
	Name:          "gogoproto.goproto_unrecognized",
	Tag:           "varint,64026,opt,name=goproto_unrecognized",
	Filename:      "gogo.proto",
}

var E_Protosizer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64028,
	Name:          "gogoproto.protosizer",
	Tag:           "varint,64028,opt,name=protosizer",
	Filename:      "gogo.proto",
}

var E_Compare = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64029,
	Name:          "gogoproto.compare",
	Tag:           "varint,64029,opt,name=compare",
	Filename:      "gogo.proto",
}

var E_Typedecl = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64030,
	Name:          "gogoproto.typedecl",
	Tag:           "varint,64030,opt,name=typedecl",
	Filename:      "gogo.proto",
}

var E_Messagename = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64033,
	Name:          "gogoproto.messagename",
	Tag:           "varint,64033,opt,name=messagename",
	Filename:      "gogo.proto",
}

var E_GoprotoSizecache = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64034,
	Name:          "gogoproto.goproto_sizecache",
	Tag:           "varint,64034,opt,name=goproto_sizecache",
	Filename:      "gogo.proto",
}

var E_GoprotoUnkeyed = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64035,
	Name:          "gogoproto.goproto_unkeyed",
	Tag:           "varint,64035,opt,name=goproto_unkeyed",
	Filename:      "gogo.proto",
}

var E_Nullable = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65001,
	Name:          "gogoproto.nullable",
	Tag:           "varint,65001,opt,name=nullable",
	Filename:      "gogo.proto",
}

var E_Embed = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65002,
	Name:          "gogoproto.embed",
	Tag:           "varint,65002,opt,name=embed",
	Filename:      "gogo.proto",
}

var E_Customtype = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65003,
	Name:          "gogoproto.customtype",
	Tag:           "bytes,65003,opt,name=customtype",
	Filename:      "gogo.proto",
}

var E_Customname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65004,
	Name:          "gogoproto.customname",
	Tag:           "bytes,65004,opt,name=customname",
	Filename:      "gogo.proto",
}

var E_Jsontag = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65005,
	Name:          "gogoproto.jsontag",
	Tag:           "bytes,65005,opt,name=jsontag",
	Filename:      "gogo.proto",
}

var E_Moretags = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65006,
	Name:          "gogoproto.moretags",
	Tag:           "bytes,65006,opt,name=moretags",
	Filename:      "gogo.proto",
}

var E_Casttype = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65007,
	Name:          "gogoproto.casttype",
	Tag:           "bytes,65007,opt,name=casttype",
	Filename:      "gogo.proto",
}

var E_Castkey = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65008,
	Name:          "gogoproto.castkey",
	Tag:           "bytes,65008,opt,name=castkey",
	Filename:      "gogo.proto",
}

var E_Castvalue = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65009,
	Name:          "gogoproto.castvalue",
	Tag:           "bytes,65009,opt,name=castvalue",
	Filename:      "gogo.proto",
}

var E_Stdtime = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65010,
	Name:          "gogoproto.stdtime",
	Tag:           "varint,65010,opt,name=stdtime",
	Filename:      "gogo.proto",
}

var E_Stdduration = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65011,
	Name:          "gogoproto.stdduration",
	Tag:           "varint,65011,opt,name=stdduration",
	Filename:      "gogo.proto",
}

var E_Wktpointer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65012,
	Name:          "gogoproto.wktpointer",
	Tag:           "varint,65012,opt,name=wktpointer",
	Filename:      "gogo.proto",
}

func init() {
	proto.RegisterExtension(E_GoprotoEnumPrefix)
	proto.RegisterExtension(E_GoprotoEnumStringer)
	proto.RegisterExtension(E_EnumStringer)
	proto.RegisterExtension(E_EnumCustomname)
	proto.RegisterExtension(E_Enumdecl)
	proto.RegisterExtension(E_EnumvalueCustomname)
	proto.RegisterExtension(E_GoprotoGettersAll)
	proto.RegisterExtension(E_GoprotoEnumPrefixAll)
	proto.RegisterExtension(E_GoprotoStringerAll)
	proto.RegisterExtension(E_VerboseEqualAll)
	proto.RegisterExtension(E_FaceAll)
	proto.RegisterExtension(E_GostringAll)
	proto.RegisterExtension(E_PopulateAll)
	proto.RegisterExtension(E_StringerAll)
	proto.RegisterExtension(E_OnlyoneAll)
	proto.RegisterExtension(E_EqualAll)
	proto.RegisterExtension(E_DescriptionAll)
	proto.RegisterExtension(E_TestgenAll)
	proto.RegisterExtension(E_BenchgenAll)
	proto.RegisterExtension(E_MarshalerAll)
	proto.RegisterExtension(E_UnmarshalerAll)
	proto.RegisterExtension(E_StableMarshalerAll)
	proto.RegisterExtension(E_SizerAll)
	proto.RegisterExtension(E_GoprotoEnumStringerAll)
	proto.RegisterExtension(E_EnumStringerAll)
	proto.RegisterExtension(E_UnsafeMarshalerAll)
	proto.RegisterExtension(E_UnsafeUnmarshalerAll)
	proto.RegisterExtension(E_GoprotoExtensionsMapAll)
	proto.RegisterExtension(E_GoprotoUnrecognizedAll)
	proto.RegisterExtension(E_GogoprotoImport)
	proto.RegisterExtension(E_ProtosizerAll)
	proto.RegisterExtension(E_CompareAll)
	proto.RegisterExtension(E_TypedeclAll)
	proto.RegisterExtension(E_EnumdeclAll)
	proto.RegisterExtension(E_GoprotoRegistration)
	proto.RegisterExtension(E_MessagenameAll)
	proto.RegisterExtension(E_GoprotoSizecacheAll)
	proto.RegisterExtension(E_GoprotoUnkeyedAll)
	proto.RegisterExtension(E_GoprotoGetters)
	proto.RegisterExtension(E_GoprotoStringer)
	proto.RegisterExtension(E_VerboseEqual)
	proto.RegisterExtension(E_Face)
	proto.RegisterExtension(E_Gostring)
	proto.RegisterExtension(E_Populate)
	proto.RegisterExtension(E_Stringer)
	proto.RegisterExtension(E_Onlyone)
	proto.RegisterExtension(E_Equal)
	proto.RegisterExtension(E_Description)
	proto.RegisterExtension(E_Testgen)
	proto.RegisterExtension(E_Benchgen)
	proto.RegisterExtension(E_Marshaler)
	proto.RegisterExtension(E_Unmarshaler)
	proto.RegisterExtension(E_StableMarshaler)
	proto.RegisterExtension(E_Sizer)
	proto.RegisterExtension(E_UnsafeMarshaler)
	proto.RegisterExtension(E_UnsafeUnmarshaler)
	proto.RegisterExtension(E_GoprotoExtensionsMap)
	proto.RegisterExtension(E_GoprotoUnrecognized)
	proto.RegisterExtension(E_Protosizer)
	proto.RegisterExtension(E_Compare)
	proto.RegisterExtension(E_Typedecl)
	proto.RegisterExtension(E_Messagename)
	proto.RegisterExtension(E_GoprotoSizecache)
	proto.RegisterExtension(E_GoprotoUnkeyed)
	proto.RegisterExtension(E_Nullable)
	proto.RegisterExtension(E_Embed)
	proto.RegisterExtension(E_Customtype)
	proto.RegisterExtension(E_Customname)
	proto.RegisterExtension(E_Jsontag)
	proto.RegisterExtension(E_Moretags)
	proto.RegisterExtension(E_Casttype)
	proto.RegisterExtension(E_Castkey)
	proto.RegisterExtension(E_Castvalue)
	proto.RegisterExtension(E_Stdtime)
	proto.RegisterExtension(E_Stdduration)
	proto.RegisterExtension(E_Wktpointer)
}

func init() { proto.RegisterFile("gogo.proto", fileDescriptor_592445b5231bc2b9) }

var fileDescriptor_592445b5231bc2b9 = []byte{
	// 1361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x98, 0x4b, 0x6c, 0x15, 0x55,
	0x18, 0x80, 0x69, 0x84, 0xd0, 0x9e, 0xb6, 0x94, 0xb6, 0x88, 0x48, 0xf4, 0x8a, 0xae, 0x5c, 0x95,
	0x15, 0x31, 0x9c, 0x86, 0x90, 0xd2, 0x94, 0x06, 0x23, 0x58, 0x0b, 0xc5, 0x57, 0xcc, 0xcd, 0xdc,
	0xb9, 0xa7, 0xd3, 0x91, 0x99, 0x39, 0xe3, 0xcc, 0x19, 0xa4, 0xec, 0x0c, 0x3e, 0x62, 0x8c, 0x8a,
	0x8f, 0x44, 0x41, 0x40, 0xd1, 0xf8, 0x7e, 0xe2, 0xfb, 0xb1, 0x71, 0xa3, 0xb2, 0xc4, 0x9d, 0x4b,
	0x03, 0x6e, 0xd4, 0xfa, 0xee, 0xae, 0x1b, 0xf3, 0xcf, 0xfc, 0xff, 0xdc, 0x73, 0xa7, 0x37, 0x39,
	0xe7, 0xee, 0xa6, 0xb7, 0xe7, 0xfb, 0xee, 0x99, 0xff, 0x9f, 0xf3, 0xff, 0xff, 0x1d, 0xc6, 0x3c,
	0xe9, 0xc9, 0xb1, 0x38, 0x91, 0x4a, 0x8e, 0xf4, 0xc1, 0x75, 0x7e, 0xb9, 0x75, 0x9b, 0x27, 0xa5,
	0x17, 0x88, 0xed, 0xf9, 0x5f, 0x8d, 0x6c, 0x7e, 0x7b, 0x53, 0xa4, 0x6e, 0xe2, 0xc7, 0x4a, 0x26,
	0xc5, 0x62, 0x7e, 0x80, 0x8d, 0xe2, 0xe2, 0xba, 0x88, 0xb2, 0xb0, 0x1e, 0x27, 0x62, 0xde, 0x3f,
	0x36, 0x72, 0xdd, 0x58, 0x41, 0x8e, 0x11, 0x39, 0x36, 0x15, 0x65, 0xe1, 0xed, 0xb1, 0xf2, 0x65,
	0x94, 0x6e, 0xb9, 0xf0, 0xcb, 0x55, 0xdb, 0x7a, 0x6e, 0xee, 0x9d, 0x1d, 0x46, 0x14, 0xfe, 0x37,
	0x93, 0x83, 0x7c, 0x96, 0x5d, 0xdd, 0xe6, 0x4b, 0x55, 0xe2, 0x47, 0x9e, 0x48, 0x0c, 0xc6, 0xef,
	0xd0, 0x38, 0xaa, 0x19, 0x0f, 0x22, 0xca, 0x27, 0xd9, 0x60, 0x37, 0xae, 0xef, 0xd1, 0x35, 0x20,
	0x74, 0xc9, 0x34, 0x1b, 0xca, 0x25, 0x6e, 0x96, 0x2a, 0x19, 0x46, 0x4e, 0x28, 0x0c, 0x9a, 0x1f,
	0x72, 0x4d, 0xdf, 0xec, 0x06, 0xc0, 0x26, 0x4b, 0x8a, 0x73, 0xd6, 0x0b, 0x9f, 0x34, 0x85, 0x1b,
	0x18, 0x0c, 0x17, 0x71, 0x23, 0xe5, 0x7a, 0x7e, 0x98, 0x6d, 0x82, 0xeb, 0xa3, 0x4e, 0x90, 0x09,
	0x7d, 0x27, 0x37, 0x76, 0xf4, 0x1c, 0x86, 0x65, 0x24, 0xfb, 0xf1, 0xc4, 0xda, 0x7c, 0x3b, 0xa3,
	0xa5, 0x40, 0xdb, 0x93, 0x96, 0x45, 0x4f, 0x28, 0x25, 0x92, 0xb4, 0xee, 0x04, 0x9d, 0xb6, 0xb7,
	0xd7, 0x0f, 0x4a, 0xe3, 0xa9, 0xa5, 0xf6, 0x2c, 0x4e, 0x17, 0xe4, 0x44, 0x10, 0xf0, 0x39, 0x76,
	0x4d, 0x87, 0xa7, 0xc2, 0xc2, 0x79, 0x1a, 0x9d, 0x9b, 0x56, 0x3d, 0x19, 0xa0, 0x9d, 0x61, 0xf4,
	0x79, 0x99, 0x4b, 0x0b, 0xe7, 0x4b, 0xe8, 0x1c, 0x41, 0x96, 0x52, 0x0a, 0xc6, 0x5b, 0xd9, 0xf0,
	0x51, 0x91, 0x34, 0x64, 0x2a, 0xea, 0xe2, 0x81, 0xcc, 0x09, 0x2c, 0x74, 0x67, 0x50, 0x37, 0x84,
	0xe0, 0x14, 0x70, 0xe0, 0xda, 0xc9, 0x7a, 0xe7, 0x1d, 0x57, 0x58, 0x28, 0xce, 0xa2, 0x62, 0x3d,
	0xac, 0x07, 0x74, 0x82, 0x0d, 0x78, 0xb2, 0xb8, 0x25, 0x0b, 0xfc, 0x1c, 0xe2, 0xfd, 0xc4, 0xa0,
	0x22, 0x96, 0x71, 0x16, 0x38, 0xca, 0x66, 0x07, 0x2f, 0x93, 0x82, 0x18, 0x54, 0x74, 0x11, 0xd6,
	0x57, 0x48, 0x91, 0x6a, 0xf1, 0xdc, 0xcd, 0xfa, 0x65, 0x14, 0x2c, 0xca, 0xc8, 0x66, 0x13, 0xe7,
	0xd1, 0xc0, 0x10, 0x01, 0xc1, 0x38, 0xeb, 0xb3, 0x4d, 0xc4, 0xeb, 0x4b, 0x74, 0x3c, 0x28, 0x03,
	0xd3, 0x6c, 0x88, 0x0a, 0x94, 0x2f, 0x23, 0x0b, 0xc5, 0x1b, 0xa8, 0xd8, 0xa0, 0x61, 0x78, 0x1b,
	0x4a, 0xa4, 0xca, 0x13, 0x36, 0x92, 0x37, 0xe9, 0x36, 0x10, 0xc1, 0x50, 0x36, 0x44, 0xe4, 0x2e,
	0xd8, 0x19, 0xde, 0xa2, 0x50, 0x12, 0x03, 0x8a, 0x49, 0x36, 0x18, 0x3a, 0x49, 0xba, 0xe0, 0x04,
	0x56, 0xe9, 0x78, 0x1b, 0x1d, 0x03, 0x25, 0x84, 0x11, 0xc9, 0xa2, 0x6e, 0x34, 0xef, 0x50, 0x44,
	0x34, 0x0c, 0x8f, 0x5e, 0xaa, 0x9c, 0x46, 0x20, 0xea, 0xdd, 0xd8, 0xde, 0xa5, 0xa3, 0x57, 0xb0,
	0xfb, 0x75, 0xe3, 0x38, 0xeb, 0x4b, 0xfd, 0xe3, 0x56, 0x9a, 0xf7, 0x28, 0xd3, 0x39, 0x00, 0xf0,
	0xdd, 0xec, 0xda, 0x8e, 0x6d, 0xc2, 0x42, 0xf6, 0x3e, 0xca, 0x36, 0x77, 0x68, 0x15, 0x58, 0x12,
	0xba, 0x55, 0x7e, 0x40, 0x25, 0x41, 0x54, 0x5c, 0x33, 0x6c, 0x53, 0x16, 0xa5, 0xce, 0x7c, 0x77,
	0x51, 0xfb, 0x90, 0xa2, 0x56, 0xb0, 0x6d, 0x51, 0x3b, 0xc4, 0x36, 0xa3, 0xb1, 0xbb, 0xbc, 0x7e,
	0x44, 0x85, 0xb5, 0xa0, 0xe7, 0xda, 0xb3, 0x7b, 0x2f, 0xdb, 0x5a, 0x86, 0xf3, 0x98, 0x12, 0x51,
	0x0a, 0x4c, 0x3d, 0x74, 0x62, 0x0b, 0xf3, 0x05, 0x34, 0x53, 0xc5, 0x9f, 0x2a, 0x05, 0xfb, 0x9d,
	0x18, 0xe4, 0x77, 0xb1, 0x2d, 0x24, 0xcf, 0xa2, 0x44, 0xb8, 0xd2, 0x8b, 0xfc, 0xe3, 0xa2, 0x69,
	0xa1, 0xfe, 0xb8, 0x92, 0xaa, 0x39, 0x0d, 0x07, 0xf3, 0x3e, 0xb6, 0xb1, 0x9c, 0x55, 0xea, 0x7e,
	0x18, 0xcb, 0x44, 0x19, 0x8c, 0x9f, 0x50, 0xa6, 0x4a, 0x6e, 0x5f, 0x8e, 0xf1, 0x29, 0xb6, 0x21,
	0xff, 0xd3, 0xf6, 0x91, 0xfc, 0x14, 0x45, 0x83, 0x2d, 0x0a, 0x0b, 0x87, 0x2b, 0xc3, 0xd8, 0x49,
	0x6c, 0xea, 0xdf, 0x67, 0x54, 0x38, 0x10, 0xc1, 0xc2, 0xa1, 0x16, 0x63, 0x01, 0xdd, 0xde, 0xc2,
	0xf0, 0x39, 0x15, 0x0e, 0x62, 0x50, 0x41, 0x03, 0x83, 0x85, 0xe2, 0x0b, 0x52, 0x10, 0x03, 0x8a,
	0x3b, 0x5a, 0x8d, 0x36, 0x11, 0x9e, 0x9f, 0xaa, 0xc4, 0x81, 0xd5, 0x06, 0xd5, 0x97, 0x4b, 0xed,
	0x43, 0xd8, 0xac, 0x86, 0x42, 0x25, 0x0a, 0x45, 0x9a, 0x3a, 0x9e, 0x80, 0x89, 0xc3, 0x62, 0x63,
	0x5f, 0x51, 0x25, 0xd2, 0x30, 0xd8, 0x9b, 0x36, 0x21, 0x42, 0xd8, 0x5d, 0xc7, 0x5d, 0xb0, 0xd1,
	0x7d, 0x5d, 0xd9, 0xdc, 0x41, 0x62, 0xc1, 0xa9, 0xcd, 0x3f, 0x59, 0x74, 0x44, 0x2c, 0x5a, 0x3d,
	0x9d, 0xdf, 0x54, 0xe6, 0x9f, 0xb9, 0x82, 0x2c, 0x6a, 0xc8, 0x50, 0x65, 0x9e, 0x1a, 0xb9, 0x61,
	0x95, 0x6b, 0x7f, 0x71, 0x5f, 0xa4, 0x7b, 0x68, 0x19, 0xef, 0xb7, 0x7d, 0x9c, 0xe2, 0xb7, 0xc1,
	0x43, 0xde, 0x3e, 0xf4, 0x98, 0x65, 0x27, 0x96, 0xcb, 0xe7, 0xbc, 0x6d, 0xe6, 0xe1, 0x7b, 0xd9,
	0x60, 0xdb, 0xc0, 0x63, 0x56, 0x3d, 0x8c, 0xaa, 0x01, 0x7d, 0xde, 0xe1, 0x3b, 0xd8, 0x5a, 0x18,
	0x5e, 0xcc, 0xf8, 0x23, 0x88, 0xe7, 0xcb, 0xf9, 0x2e, 0xd6, 0x4b, 0x43, 0x8b, 0x19, 0x7d, 0x14,
	0xd1, 0x12, 0x01, 0x9c, 0x06, 0x16, 0x33, 0xfe, 0x18, 0xe1, 0x84, 0x00, 0x6e, 0x1f, 0xc2, 0x6f,
	0x9f, 0x58, 0x8b, 0x4d, 0x87, 0x62, 0x37, 0xce, 0xd6, 0xe3, 0xa4, 0x62, 0xa6, 0x1f, 0xc7, 0x2f,
	0x27, 0x82, 0xdf, 0xc2, 0xd6, 0x59, 0x06, 0xfc, 0x49, 0x44, 0x8b, 0xf5, 0x7c, 0x92, 0xf5, 0x6b,
	0xd3, 0x89, 0x19, 0x7f, 0x0a, 0x71, 0x9d, 0x82, 0xad, 0xe3, 0x74, 0x62, 0x16, 0x3c, 0x4d, 0x5b,
	0x47, 0x02, 0xc2, 0x46, 0x83, 0x89, 0x99, 0x3e, 0x49, 0x51, 0x27, 0x84, 0xef, 0x66, 0x7d, 0x65,
	0xb3, 0x31, 0xf3, 0xcf, 0x20, 0xdf, 0x62, 0x20, 0x02, 0x5a, 0xb3, 0x33, 0x2b, 0x9e, 0xa5, 0x08,
	0x68, 0x14, 0x1c, 0xa3, 0xea, 0x00, 0x63, 0x36, 0x3d, 0x47, 0xc7, 0xa8, 0x32, 0xbf, 0x40, 0x36,
	0xf3, 0x9a, 0x6f, 0x56, 0x3c, 0x4f, 0xd9, 0xcc, 0xd7, 0xc3, 0x36, 0xaa, 0x13, 0x81, 0xd9, 0xf1,
	0x02, 0x6d, 0xa3, 0x32, 0x10, 0xf0, 0x19, 0x36, 0xb2, 0x7a, 0x1a, 0x30, 0xfb, 0x5e, 0x44, 0xdf,
	0xf0, 0xaa, 0x61, 0x80, 0xdf, 0xc9, 0x36, 0x77, 0x9e, 0x04, 0xcc, 0xd6, 0x53, 0xcb, 0x95, 0xdf,
	0x6e, 0xfa, 0x20, 0xc0, 0x0f, 0xb5, 0x5a, 0x8a, 0x3e, 0x05, 0x98, 0xb5, 0xa7, 0x97, 0xdb, 0x0b,
	0xb7, 0x3e, 0x04, 0xf0, 0x09, 0xc6, 0x5a, 0x0d, 0xd8, 0xec, 0x3a, 0x83, 0x2e, 0x0d, 0x82, 0xa3,
	0x81, 0xfd, 0xd7, 0xcc, 0x9f, 0xa5, 0xa3, 0x81, 0x04, 0x1c, 0x0d, 0x6a, 0xbd, 0x66, 0xfa, 0x1c,
	0x1d, 0x0d, 0x42, 0xe0, 0xc9, 0xd6, 0xba, 0x9b, 0xd9, 0x70, 0x9e, 0x9e, 0x6c, 0x8d, 0xe2, 0x07,
	0xd8, 0xf0, 0xaa, 0x86, 0x68, 0x56, 0xbd, 0x8a, 0xaa, 0x8d, 0xd5, 0x7e, 0xa8, 0x37, 0x2f, 0x6c,
	0x86, 0x66, 0xdb, 0x6b, 0x95, 0xe6, 0x85, 0xbd, 0x90, 0x8f, 0xb3, 0xde, 0x28, 0x0b, 0x02, 0x38,
	0x3c, 0x23, 0xd7, 0x77, 0xe8, 0xa6, 0x22, 0x68, 0x92, 0xe2, 0xd7, 0x15, 0x8c, 0x0e, 0x01, 0x7c,
	0x07, 0x5b, 0x27, 0xc2, 0x86, 0x68, 0x9a, 0xc8, 0xdf, 0x56, 0xa8, 0x60, 0xc2, 0x6a, 0xbe, 0x9b,
	0xb1, 0xe2, 0xd5, 0x08, 0x84, 0xd9, 0xc4, 0xfe, 0xbe, 0x52, 0xbc, 0xa5, 0xd1, 0x90, 0x96, 0x20,
	0x4f, 0x8a, 0x41, 0xb0, 0xd4, 0x2e, 0xc8, 0x33, 0xb2, 0x93, 0xad, 0xbf, 0x3f, 0x95, 0x91, 0x72,
	0x3c, 0x13, 0xfd, 0x07, 0xd2, 0xb4, 0x1e, 0x02, 0x16, 0xca, 0x44, 0x28, 0xc7, 0x4b, 0x4d, 0xec,
	0x9f, 0xc8, 0x96, 0x00, 0xc0, 0xae, 0x93, 0x2a, 0x9b, 0xfb, 0xfe, 0x8b, 0x60, 0x02, 0x60, 0xd3,
	0x70, 0x7d, 0x44, 0x2c, 0x9a, 0xd8, 0xbf, 0x69, 0xd3, 0xb8, 0x9e, 0xef, 0x62, 0x7d, 0x70, 0x99,
	0xbf, 0x55, 0x32, 0xc1, 0xff, 0x20, 0xdc, 0x22, 0xe0, 0x9b, 0x53, 0xd5, 0x54, 0xbe, 0x39, 0xd8,
	0xff, 0x62, 0xa6, 0x69, 0x3d, 0x9f, 0x60, 0xfd, 0xa9, 0x6a, 0x36, 0x33, 0x9c, 0x4f, 0x0d, 0xf8,
	0x7f, 0x2b, 0xe5, 0x2b, 0x8b, 0x92, 0x81, 0x6c, 0x3f, 0x78, 0x44, 0xc5, 0xd2, 0x8f, 0x94, 0x48,
	0x4c, 0x86, 0x65, 0x34, 0x68, 0xc8, 0x9e, 0xfb, 0x2e, 0x5e, 0xae, 0xf5, 0x5c, 0xba, 0x5c, 0xeb,
	0xf9, 0xf9, 0x72, 0xad, 0xe7, 0xe4, 0x95, 0xda, 0x9a, 0x4b, 0x57, 0x6a, 0x6b, 0x7e, 0xba, 0x52,
	0x5b, 0xc3, 0x46, 0x5d, 0x19, 0x56, 0x7d, 0x7b, 0xd8, 0xb4, 0x9c, 0x96, 0x33, 0x79, 0xfd, 0xb9,
	0xe7, 0x26, 0xcf, 0x57, 0x0b, 0x59, 0x63, 0xcc, 0x95, 0xe1, 0x76, 0xf8, 0x45, 0xd2, 0x7a, 0xd1,
	0x5a, 0xfe, 0x3e, 0xf9, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x47, 0x65, 0x6e, 0xf9, 0x9b, 0x15, 0x00,
	0x00,
}
