load("@gazelle//:def.bzl", "gazelle")
load("@rules_go//go:def.bzl", "go_library")

gazelle(name = "gazelle")

go_library(
    name = "greetings",
    srcs = ["greetings.go"],
    importpath = "bdl.com/greetings",
    visibility = ["//visibility:public"],
)
