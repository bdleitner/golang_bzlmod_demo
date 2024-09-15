# BzlMod + Golang Demos: Hello

This project is to test the inclusion of an external but local dependency. This project depends
on the [greetings](../greetings/README.md) project and is geared to mirror [this tutorial]
(https://go.dev/doc/tutorial/call-module-code) but using bazel/bzlmod

After writing the go code per the tutorial, I went to go use the tooling:
* `bazel run @rules_go/go mod init` created a `go.mod` file.
* `bazel run //:gazelle` generated the build targets as expected, including the dependency on 
  `@com_bdl_greetings//:greetings`.

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
