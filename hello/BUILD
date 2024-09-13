load("@gazelle//:def.bzl", "gazelle")
load("@rules_go//go:def.bzl", "go_binary", "go_library")

gazelle(name = "gazelle")

go_library(
    name = "hello_lib",
    srcs = ["hello.go"],
    importpath = "bdl.com/hello",
    visibility = ["//visibility:private"],
    deps = ["@com_bdl_greetings//:greetings"],
)

go_binary(
    name = "hello",
    embed = [":hello_lib"],
    visibility = ["//visibility:public"],
)
