load("@web_server//:requirements.bzl", "requirement")
load("@rules_python//python:defs.bzl", "py_binary")

py_binary(
    name = "app",
    srcs = ["app.py"],
    deps = [
        "//protos:shortn_py_pb",
        "//protos:shortn_py_pb_grpc",
        requirement("protobuf"),
        requirement("flask"),
        requirement("grpcio"),
        requirement("grpcio-tools"),
        requirement("jinja2"),
        requirement("markupsafe"),
        requirement("werkzeug"),
        requirement("itsdangerous"),
        requirement("click"),
        requirement("six"),
    ],
)
