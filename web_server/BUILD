load("@rules_python//python:defs.bzl", "py_binary")
load("@web_server//:requirements.bzl", "requirement")

py_binary(
    name = "app",
    srcs = ["app.py"],
    deps = [
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
