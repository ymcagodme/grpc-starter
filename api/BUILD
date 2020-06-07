load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "grpc_server",
    srcs = ["main.go"],
    deps = [
        "//api/core:add_page",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
