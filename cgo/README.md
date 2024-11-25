# Bzlmod + Golang Demo: CGO Usage

This very small module highlights an issue with using CGO and bazel/bzlmod.

When the code using "C" inside `cgo_demo_lib.go`'s `Example` function is commented out
(along with `import "C"`), everything compiles.

But when this code is uncommented, so it actually uses the C code (example taken from https://go.dev/wiki/cgo)
Then: `bazel build //lib:lib` succeeds, but:

* `bazel build //lib:lib_test`
* `bazel build //main:main_lib`
* `bazel build //main:main`


All fail with:

```
INFO: Analyzed target //lib:lib_test (3 packages loaded, 26 targets configured).
ERROR: C:/projects/go/demos/cgo/lib/BUILD:11:8: GoCompilePkg lib/lib_test.internal.a failed: (Exit 1): builder.exe failed: error executing GoCompilePkg command (from target //lib:lib_test) bazel-out\x64_windows-opt-exec-ST-d57f47055a04\bin\external\rules_go~~go_sdk~go_default_sdk\builder_reset\builder.exe compilepkg -sdk external/rules_go~~go_sdk~go_default_sdk -goroot ... (remaining 29 arguments skipped)
lib\cgo_demo_lib_test.go:8:2: undefined: Example
compilepkg: error running subcommand external\rules_go~~go_sdk~go_default_sdk\pkg\tool\windows_amd64\compile.exe: exit status 2
Target //lib:lib_test failed to build                                                                                                                                                                                               
Use --verbose_failures to see the command lines of failed build steps.                                                                                                                                                              
INFO: Elapsed time: 1.331s, Critical Path: 0.06s
INFO: 3 processes: 2 internal, 1 local.                                                                                                                                                                                             
ERROR: Build did NOT complete successfully 
```

The `Example` function is not found and gives an `undefined: Example` error.

I first discovered this while attempting to import a cgo-utilizing package.
Minimal example for that case is [here](../cgoimport/README.md).