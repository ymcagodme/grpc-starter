load("@rules_proto//proto:defs.bzl", "proto_library")
load("@com_github_grpc_grpc//bazel:python_rules.bzl", "py_grpc_library", "py_proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "shortn_proto",
    srcs = ["shortn.proto"],
)

py_proto_library(
    # Note: the output file name would always have {{filename}}_pb2.py.
    # https://github.com/grpc/grpc/issues/15444
    name = "shortn_py_pb",
    visibility = ["//web_server:__subpackages__"],
    deps = [":shortn_proto"],
)

py_grpc_library(
    name = "shortn_py_pb_grpc",
    srcs = [":shortn_proto"],
    visibility = ["//web_server:__subpackages__"],
    deps = [":shortn_py_pb"],
)

go_proto_library(
    name = "shortn_go_pb",
    compiler = "@io_bazel_rules_go//proto:go_grpc",
    importpath = "github.com/ymcagodme/shortn/protos",
    proto = ":shortn_proto",
    visibility = ["//api:__subpackages__"],
)
