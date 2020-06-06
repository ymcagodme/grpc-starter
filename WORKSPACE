load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# -------------------
# Load Python support
# -------------------
http_archive(
    name = "rules_python",
    sha256 = "b5668cde8bb6e3515057ef465a35ad712214962f0b3a314e551204266c7be90c",
    strip_prefix = "rules_python-0.0.2",
    url = "https://github.com/bazelbuild/rules_python/releases/download/0.0.2/rules_python-0.0.2.tar.gz",
)

load("@rules_python//python:repositories.bzl", "py_repositories")

py_repositories()

# Only needed if using the packaging rules.
load("@rules_python//python:pip.bzl", "pip_import", "pip_repositories")

pip_repositories()

# Create a central repo that knows about the dependencies needed for
# requirements.txt.
pip_import(
    # or pip3_import
    name = "web_server",
    requirements = "//web_server:requirements.txt",
)

# Load the central repo's install function from its `//:requirements.bzl` file,
# and call it to install all deps in requirements.
load("@web_server//:requirements.bzl", webserver_pip_install = "pip_install")

webserver_pip_install()

# ---------------------
# Load protobuf support
# ---------------------
http_archive(
    name = "rules_proto",
    strip_prefix = "rules_proto-master",
    url = "https://github.com/bazelbuild/rules_proto/archive/master.zip",
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

# -----------------
# Load gRPC support
# -----------------
http_archive(
    name = "com_github_grpc_grpc",
    sha256 = "3eb0b15a107aae2f61f49574e7131b2da368231e53d7acca59c25469a58c5ebe",
    strip_prefix = "grpc-a04f0db4dba6294035ba0d48270621d6357fba17",
    urls = [
        "https://github.com/grpc/grpc/archive/a04f0db4dba6294035ba0d48270621d6357fba17.tar.gz",
    ],
)

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

load("@com_github_grpc_grpc//bazel:grpc_extra_deps.bzl", "grpc_extra_deps")

grpc_extra_deps()
