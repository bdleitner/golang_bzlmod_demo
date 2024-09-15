# BzlMod + Golang Simple Proto Example

Trying to follow the guidelines from
the [rules_go](https://github.com/bazelbuild/rules_go/blob/master/docs/go/core/bzlmod.md) bzlmod
reference.

In this project, I just have a simple proto use. `bael run //:gazelle` works
and `bazel build` is able to build the target. Since I'm not depending on the golang `proto` 
package, that's all there is.

But of note: `bazel run @rules_go/go -- mod tidy -v`

Gives the following output:
```
PS C:\projects\go\demos\simpleproto> bazel run @rules_go//go -- mod tidy -v
INFO: Analyzed target @@rules_go~//go:go (2 packages loaded, 12 targets configured).
INFO: Found 1 target...
Target @@rules_go~//go/tools/go_bin_runner:go_bin_runner up-to-date:                                                                                                                                      
  bazel-bin/external/rules_go~/go/tools/go_bin_runner/bin/go.exe                                                                                                                                          
INFO: Elapsed time: 1.624s, Critical Path: 0.03s                                                                                                                                                          
INFO: 1 process: 1 internal.                                                                                                                                                                              
INFO: Build completed successfully, 1 total action                                                                                                                                                        
INFO: Running command line: bazel-bin/external/rules_go~/go/tools/go_bin_runner/bin/go.exe mod tidy -v
go: finding module for package bdl.com/demos/simpleproto/expb
go: bdl.com/demos/proto imports
        bdl.com/demos/simpleproto/expb: cannot find module providing package bdl.com/demos/simpleproto/expb: unrecognized import path "bdl.com/demos/simpleproto/expb": reading https://bdl.com/demos/simpleproto/expb?go-get=1: 404 Not Found
```

So even though the `go_proto_library` says that's its importpath, the `go mod tidy` run cannot 
find it.
But a Bazel sync (using the Goland plugin) does work and everything seems ok.

Then let's see what happens when we try to depend on the golang proto library:

See [the protoexp README](../protoexp/README.md).