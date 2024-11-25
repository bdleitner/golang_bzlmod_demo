load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "main_lib",
    srcs = ["cgo_demo.go"],
    importpath = "bdl.com/cgo/main",
    visibility = ["//visibility:private"],
    deps = ["//lib"],
)

go_binary(
    name = "main",
    embed = [":main_lib"],
    visibility = ["//visibility:public"],
)
