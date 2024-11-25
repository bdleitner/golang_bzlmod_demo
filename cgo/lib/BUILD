load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "lib",
    srcs = ["cgo_demo_lib.go"],
    cgo = True,
    importpath = "bdl.com/cgo/lib",
    visibility = ["//visibility:public"],
)

go_test(
    name = "lib_test",
    srcs = ["cgo_demo_lib_test.go"],
    embed = [":lib"],
)
