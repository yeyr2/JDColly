# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: wordCloud.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()

DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0fwordCloud.proto\x12\tmyservice\"0\n\nrpcComment\x12'
                                                          b'\x0f\n\x07\x43ontent\x18\x01 \x03('
                                                          b'\t\x12\x11\n\tproductId\x18\x02 \x01('
                                                          b'\t\"\"\n\x0crpcWordCloud\x12\x12\n\nwordsCloud\x18\x01 '
                                                          b'\x01('
                                                          b'\t2P\n\x07Greeter\x12\x45\n\x11WordCloudAnalysis\x12\x15.myservice.rpcComment\x1a\x17.myservice.rpcWordCloud\"\x00\x42\x13Z\x11goGRPC/WordsCloudb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'wordCloud_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:
    DESCRIPTOR._options = None
    DESCRIPTOR._serialized_options = b'Z\021goGRPC/WordsCloud'
    DESCRIPTOR._RPCCOMMENT._serialized_start = 30
    DESCRIPTOR._RPCCOMMENT._serialized_end = 78
    DESCRIPTOR._RPCWORDCLOUD._serialized_start = 80
    DESCRIPTOR._RPCWORDCLOUD._serialized_end = 114
    DESCRIPTOR._GREETER._serialized_start = 116
    DESCRIPTOR._GREETER._serialized_end = 196
# @@protoc_insertion_point(module_scope)
