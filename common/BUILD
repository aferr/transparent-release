#
# Copyright 2022 The Project Oak Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//:__subpackages__"])

go_library(
    name = "common",
    srcs = ["common.go"],
    deps = [
        "@com_github_pelletier_toml//:go_default_library",
        "//slsa:slsa",
    ],
    importpath = "github.com/project-oak/transparent-release/common",
)

go_test(
    name = "common_test",
    size = "small",
    srcs = ["common_test.go"],
    data = [
        "//testdata:build.toml",
        "//schema/amber-slsa-buildtype/v1:provenance.json",
        "//schema/amber-slsa-buildtype/v1:example.json"
    ],
    embed = [":common"],
    deps = [
        "@com_github_google_go_cmp//cmp:go_default_library",
        "//slsa:slsa",
    ],
)
