load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "ocr_lib",
    srcs = ["ocr.go"],
    importpath = "bdl.com/cgoimport",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_otiai10_gosseract_v2//:gosseract",
    ],
    cgo = True,

)

go_binary(
    name = "ocr",
    embed = [":ocr_lib"],
    visibility = ["//visibility:public"],
)
