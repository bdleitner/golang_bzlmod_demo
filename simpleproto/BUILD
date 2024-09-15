#gazelle:prefix bdl.com/demos/simpleproto
load("@gazelle//:def.bzl", "gazelle")
load("@rules_go//go:def.bzl", "go_library")
load("@rules_go//proto:def.bzl", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")

gazelle(name = "gazelle")

proto_library(
    name = "expb_proto",
    srcs = ["example.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "expb_go_proto",
    importpath = "bdl.com/demos/simpleproto/expb",
    proto = ":expb_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "simpleproto",
    srcs = ["example.go"],
    importpath = "bdl.com/demos/simpleproto",
    visibility = ["//visibility:public"],
    deps = [":expb_go_proto"],
)
