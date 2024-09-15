# BzlMod + Golang Proto External Dependency Example

Trying to follow the guidelines from
the [rules_go](https://github.com/bazelbuild/rules_go/blob/master/docs/go/core/bzlmod.md) bzlmod
reference.

This project is a copy of [simpleproto](../simpleproto/README.md), but in addition,
there is a dependency on the github golang-proto module.

In this project, I added a dependency on `github.com/golang/protobuf/proto`.

Let's see what happens with the tooling:

1. `bazel run //:gazelle` runs ok and adds the `"@com_github_golang_protobuf//proto"` dependency
   to the example `go_library` target, as expected.
2. At this point, `bazel build` fails with

    ```
    PS C:\projects\go\demos\protoexp> bazel build ...
    ERROR: no such package '@@[unknown repo 'com_github_golang_protobuf' requested from @@] //proto': The repository '@@[unknown repo 'com_github_golang_protobuf' requested from @@]' 
    could not be resolved: No repository visible as '@com_github_golang_protobuf' from main repository
    ERROR: C:/projects/go/demos/protoexp/BUILD:22:11: no such package '@@[unknown repo  'com_github_golang_protobuf' requested from @@]//proto':
        The repository '@@[unknown repo 'com_github_golang_protobuf' requested from @@]' could not be resolved: No repository visible 
        as '@com_github_golang_protobuf' from main repository and referenced by '//:protoexp'    

    ERROR: Analysis of target '//:protoexp' failed; build aborted: Analysis failed
    INFO: Elapsed time: 1.396s, Critical Path: 0.00s                                                

    INFO: 1 process: 1 internal.                                                                    

    ERROR: Build did NOT complete successfully         
    ```
   But this is expected, as the dependency has not been added yet.

3. `bazel run @rules_go/go -- mod tidy -v` gives the output:

    ```
    INFO: Running command line: bazel-bin/external/rules_go~/go/tools/go_bin_runner/bin/go.exe mod tidy -v
    go: finding module for package github.com/golang/protobuf/proto
    go: finding module for package bdl.com/demos/protoexp/expb
    go: found github.com/golang/protobuf/proto in github.com/golang/protobuf v1.5.4
    go: finding module for package bdl.com/demos/protoexp/expb
    go: bdl.com/demos/proto imports
    bdl.com/demos/protoexp/expb: cannot find module providing package bdl.com/demos/protoexp/expb: 
    unrecognized import path "bdl.com/demos/protoexp/expb": reading https://bdl.com/demos/protoexp/expb?go-get=1: 404 Not Found
    ```
   
    So the same issue with the unrecognized import as in [simpleproto](../simpleproto/README.md).
   It does find module for protobuf, but it makes no changes to the `go.mod` file.
4. I ran ` bazel run @rules_go//go -- get github.com/golang/protobuf@v1.5.4` which added
    ```
    require (
    	github.com/golang/protobuf v1.5.4 // indirect
    )
    ```
   To the `go.mod` file.  Note that I got a deprecation warning and a suggestion to use `google.
   golang.org/protobuf` instead, but we'll look at that in a minute. The [rules_go BzlMod docs 
   for external dependencies](https://github.com/bazelbuild/rules_go/blob/master/docs/go/core/bzlmod.md#external-dependencies) says that 
   in Bazel 7.1.1, the `use_repo` directive should be automatically updated. I'm on bazel 7.3.1,
   but this does not happen.
5. Manually add to the `MODULE.bazel` file:
    ```
    go_deps = use_extension("@gazelle//:extensions.bzl", "go_deps")
    go_deps.from_file(go_mod = "//:go.mod")
    use_repo(go_deps, "com_github_golang_protobuf")
   ```
   Now `bazel run @rules_go/go -- mod tidy -v` gives the message:  

   ```
   WARNING: C:/projects/go/demos/protoexp/MODULE.bazel:10:24: The module extension go_deps defined 
   in @gazelle//:extensions.bzl reported incorrect imports of repositories via use_repo():

   Imported, but reported as indirect dependencies by the extension:
   com_github_golang_protobuf

   Fix the use_repo calls by running 'bazel mod tidy'.
   ```
   The raw `bazel mod tidy` gives the same error... telling me to call `bazel mod tody`.
6. And, finally, `bazel build ...` gives the error:
    ```
    PS C:\projects\go\demos\protoexp> bazel build ...                                               
    ERROR: no such package '@@[unknown repo 'com_github_golang_protobuf' requested from @@] //proto': The repository '@@[unknown repo 'com_github_golang_protobuf' requested from @@]' could not be resolved: No repository visible as '@com_github_golang_protobuf' from main repository
    ERROR: C:/projects/go/demos/protoexp/BUILD:22:11: no such package '@@[unknown repo 'com_github_golang_protobuf' requested from @@]//proto': The repository '@@[unknown repo 'com_github_golang_protobuf' requested from @@]' could not be resolved: No repository visible as '@com_github_golang_protobuf' from main repository and referenced by '//:protoexp'
    ERROR: Analysis of target '//:protoexp' failed; build aborted: Analysis failed
    INFO: Elapsed time: 0.902s, Critical Path: 0.00s                                                 
    INFO: 1 process: 1 internal.                                                                     
    ERROR: Build did NOT complete successfully   
    ```
   
So I believe I have followed the instructions from the [`rules_go` docs](https://github.com/bazelbuild/rules_go/blob/master/docs/go/core/bzlmod.md)
but I can't get any external dependencies to resolve.  What am I missing?