# BzlMod + Golang Demo

I've been trying to migrate my existing messy Golang code library from plain bazel to use bzlmod.

I found the [rules_go](https://github.com/bazelbuild/rules_go/blob/master/docs/go/core/bzlmod.md) bzlmod reference and was trying to use that as a tutorial. I don' thave much experience with modules using `go.mod`, so at first I was avoiding it and just using my local regsitry to point at the different workspaces.  This worked mostly\* ok until I finally had some external dependencies.

* "mostly" - Goland would report on an import from a different workspace that build constraints eliminated all go files from the foreign package, even though I'm not using build constraints anywhere... at least I never set any up.

So next I decided to take the plunge and use the recommended tooling to work with go modules. But that didn't work as I wasn't able to get things to build correctly: `bazel build` would fail saying that the dependent module didn't have the right package.

That led me to some searches on how to use modules more generally, and so I found [this tutorial](https://go.dev/doc/tutorial/create-module). So I figured I should be able to use that but modify it to use bzlmod.

So I created the `greetings` package and module, starting with
```
--MODULE.bazel---
module(
    name = "com_bdl_greetings",
    compatibility_level = 0,
    version = "head",
)

bazel_dep(name = "rules_go", version = "0.50.1")
bazel_dep(name = "gazelle", version = "0.38.0")
```

and went through the process of generating `go.mod`.
A run of `bazel run //:gazelle` generated the build targets as expected.

Then I went and created the `hello` package and the tooling all worked as advertised. Per the `rules_go` instructions I used `bazel run @rules_go//go` instead of invoking the `go` command directly.
A run of `bazel run //:gazelle` generated build targets as expected, including the dependency on `@com_bdl_greetings//:greetings`, so everything was looking fine.

At first I got the error:
```
Not imported, but reported as direct dependencies by the extension (may cause the build to fail):
    com_bdl_greetings
```

But that was expected as I hadn't added
```
go_deps = use_extension("@gazelle//:extensions.bzl", "go_deps")
go_deps.from_file(go_mod = "//:go.mod")
use_repo(go_deps, "com_bdl_greetings")
```

to `MODULE.bazel` yet. So I did that and then all the `gazelle` and `go mod tidy` runs succeed and everything looks good.

But then `bazel build :hello` gives me:

```
INFO: Repository gazelle~~go_deps~com_bdl_greetings instantiated at:
  <builtin>: in <toplevel>
Repository rule go_repository defined at:
  D:/_bazel_out/5y5g3qvt/external/gazelle~/internal/go_repository.bzl:363:32: in <toplevel>
ERROR: An error occurred during the fetch of repository 'gazelle~~go_deps~com_bdl_greetings':                                                                                                             
   Traceback (most recent call last):                                                                                                                                                                     
        File "D:/_bazel_out/5y5g3qvt/external/gazelle~/internal/go_repository.bzl", line 282, column 13, in _go_repository_impl                                                                           
                fail("%s: %s" % (ctx.name, result.stderr))
Error in fail: gazelle~~go_deps~com_bdl_greetings: fetch_repo: read c:\projects\go\demos\greetings\bazel-bin: Incorrect function.
ERROR: no such package '@@gazelle~~go_deps~com_bdl_greetings//': gazelle~~go_deps~com_bdl_greetings: fetch_repo: read c:\projects\go\demos\greetings\bazel-bin: Incorrect function.
ERROR: C:/projects/go/demos/hello/BUILD:6:11: //:hello_lib depends on @@gazelle~~go_deps~com_bdl_greetings//:greetings in repository @@gazelle~~go_deps~com_bdl_greetings which failed to fetch. no such package '@@gazelle~~go_deps~com_bdl_greetings//': gazelle~~go_deps~com_bdl_greetings: fetch_repo: read c:\projects\go\demos\greetings\bazel-bin: Incorrect function.                                       
ERROR: Analysis of target '//:hello' failed; build aborted: Analysis failed                                                                                                                               
INFO: Elapsed time: 1.485s, Critical Path: 0.00s
INFO: 1 process: 1 internal.                                                                                                                                                                              
ERROR: Build did NOT complete successfully 
```

I don't understand why it says "no such package" and `c:\projects\go\demos\greetings\bazel-bin: Incorrect function.` seems very weird.

This was supposed to be my start-from-scratch very simple tutorial and I can't even make that work.
What am I missing?
