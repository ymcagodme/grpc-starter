load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Load Python support.
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
load("@web_server//:requirements.bzl", "pip_install")

pip_install()
