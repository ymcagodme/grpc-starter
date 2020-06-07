load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "grpc_server",
    srcs = ["main.go"],
    deps = [
        "//api/core:add_page",
        "//protos:shortn_go_pb",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//reflection:go_default_library",
        "@org_golang_x_net//context:go_default_library",
    ],
)
