# BzlMod + Golang Demo

This repo contains projects that demo the difficulties/errors I'm seeing when trying to get started with bzlmod and golang.
Each subfolder here is its own project (in my local system, I have separate Goland projects for each one).

* [`greetings`](greetings/README.md) - a simple module intended for use inside `hello`
* [`hello`](hello/README.md) - a module that has an external, local depdency on `greetings`
* [`simpleproto`](simpleproto/README.md) - A simple module that uses `go_proto_library` but doesn't depend on the golang proto package and its external dependency.
* [`protoexp`](protoexp/README.md) - An example that does depend on the external proto package.
I've been trying to migrate my existing messy Golang code library from plain bazel to use bzlmod.

I found the [rules_go](https://github.com/bazelbuild/rules_go/blob/master/docs/go/core/bzlmod.md) bzlmod reference and was trying to use that as a tutorial. I don' thave much experience with modules using `go.mod`, so at first I was avoiding it and just using my local regsitry to point at the different workspaces.  This worked mostly\* ok until I finally had some external dependencies.

* "mostly" - Goland would report on an import from a different workspace that build constraints eliminated all go files from the foreign package, even though I'm not using build constraints anywhere... at least I never set any up.

So next I decided to take the plunge and use the recommended tooling to work with go modules. But that didn't work as I wasn't able to get things to build correctly: `bazel build` would fail saying that the dependent module didn't have the right package.

That led me to some searches on how to use modules more generally, and so I found [this tutorial](https://go.dev/doc/tutorial/create-module). So I figured I should be able to use that but modify it to use bzlmod.

So I created the `greetings` and `hello` modules to test drive that. See the [`hello README`](hello/README.md) for the results there.

Then, later, I decided that maybe if I couldn't get "local" external dependencies to work, maybe if I could get normal ones to work I could just have one monolithic golang project. It's not what I'd like, but it could work. I ran into trouble with the first time I needed an external depdency.
