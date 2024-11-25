# Bzlmod + Golang Demo: Import that uses CGO.

This very small module highlights an issue with importing a library that uses cgo
into a workspace using bzlmod.

I wanted to try using OCR for a project and found the gosseract library.
Attempting to include it using bzlmod proved difficult, but eventually I was able to get to a point
where the import was recognized, `go.mod` had the needed require element
and the `MODULE.bazel` file included the `use_repo` for this.

At this point, `bazel run @rules_go/go -- mod tidy` runs without error and makes no changes.
A partial sync using the IJ bazel plugin also works and allows
my code to recognize the call into the gosseract library.

However, attempting a `blaze build //ocr` gives the error:

```
INFO: Analyzed target //ocr:ocr (0 packages loaded, 0 targets configured).
ERROR: C:/projects/go/demos/cgoimport/ocr/BUILD:15:10: GoCompilePkg ocr/ocr.a failed: (Exit 1): builder.exe failed: error executing GoCompilePkg command (from target //ocr:ocr) bazel-out\x64_windows-opt-exec-ST-d57f47055a04\bin\external\rules_go~~go_sdk~go_default_sdk\builder_reset\builder.exe compilepkg -sdk external/rules_go~~go_sdk~go_default_sdk -goroot ... (remaining 27 arguments skipped)
ocr\ocr.go:10:15: undefined: gs.NewClient
compilepkg: error running subcommand external\rules_go~~go_sdk~go_default_sdk\pkg\tool\windows_amd64\compile.exe: exit status 2
Target //ocr:ocr failed to build
Use --verbose_failures to see the command lines of failed build steps.
INFO: Elapsed time: 0.677s, Critical Path: 0.05s
INFO: 2 processes: 2 internal.
ERROR: Build did NOT complete successfully
```

Even though IJ recognizes `gs.NewClient`, the command line `bazel build` does not.

It gets even more interesting:
`bazel query @com_github_otiai10_gosseract_v2//...`
Gives the output:
```
@com_github_otiai10_gosseract_v2//:go_default_library                                                                                                                                                                               
@com_github_otiai10_gosseract_v2//:gosseract
@com_github_otiai10_gosseract_v2//:gosseract_test
```

The first two of these compile fine with `bazel build`, but

`bazel build @com_github_otiai10_gosseract_v2//:gosseract_test`

Gives the error:
``` 
ERROR: no such package '@@[unknown repo 'com_github_otiai10_mint' requested from @@gazelle~~go_deps~com_github_otiai10_gosseract_v2]//': The repository '@@[unknown repo 'com_github_otiai10_mint' requested from @@gazelle~~go_deps~com_github_otiai10_gosseract_v2]' could not be resolved: No repository visible as '@com_github_otiai10_mint' from repository '@@gazelle~~go_deps~com_github_otiai10_gosseract_v2'
ERROR: D:/_bazel_out/uni3i7pb/external/gazelle~~go_deps~com_github_otiai10_gosseract_v2/BUILD.bazel:39:8: no such package '@@[unknown repo 'com_github_otiai10_mint' requested from @@gazelle~~go_deps~com_github_otiai10_gosseract_v2]//': The repository '@@[unknown repo 'com_github_otiai10_mint' requested from @@gazelle~~go_deps~com_github_otiai10_gosseract_v2]' could not be resolved: No repository visible as '@com_github_otiai10_mint' from repository '@@gazelle~~go_deps~com_github_otiai10_gosseract_v2' and referenced by '@@gazelle~~go_deps~com_github_otiai10_gosseract_v2//:gosseract_test'
ERROR: Analysis of target '@@gazelle~~go_deps~com_github_otiai10_gosseract_v2//:gosseract_test' failed; build aborted: Analysis failed
INFO: Elapsed time: 0.814s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
ERROR: Build did NOT complete successfully
```

Which I take to mean that there's a dependency needed.  If I add this to `go.mod` and then
run `go mod tidy`, it gets removed (so that's an issue).

I don't remember how I got it to stick in my non-minimal example setup (maybe it only disappears 
now because of a version change in bazel?), but when I get the `go.mod` and `go.sum` files in 
sync (I think) and then try again, I get:

```
INFO: Analyzed target @@gazelle~~go_deps~com_github_otiai10_gosseract_v2//:gosseract_test (0 packages loaded, 0 targets configured).
ERROR: D:/_bazel_out/uni3i7pb/external/gazelle~~go_deps~com_github_otiai10_gosseract_v2/BUILD.bazel:39:8: GoCompilePkg external/gazelle~~go_deps~com_github_otiai10_gosseract_v2/gosseract_test.internal.a failed: (Exit 1): builder.exe failed: error executing GoCompilePkg command (from target @@gazelle~~go_deps~com_github_otiai10_gosseract_v2//:gosseract_test) bazel-out\x64_windows-opt-exec-ST-d57f47055a04\bin\external\rules_go~~go_sdk~go_default_sdk\builder_reset\builder.exe compilepkg -sdk external/rules_go~~go_sdk~go_default_sdk -goroot ... (remaining 47 arguments skipped)
external\gazelle~~go_deps~com_github_otiai10_gosseract_v2\all_test.go:23:23: undefined: Version
external\gazelle~~go_deps~com_github_otiai10_gosseract_v2\all_test.go:33:13: undefined: Version
external\gazelle~~go_deps~com_github_otiai10_gosseract_v2\all_test.go:38:12: undefined: NewClient
external\gazelle~~go_deps~com_github_otiai10_gosseract_v2\all_test.go:41:2: undefined: ClearPersistentCache
external\gazelle~~go_deps~com_github_otiai10_gosseract_v2\all_test.go:45:12: undefined: NewClient
external\gazelle~~go_deps~com_github_otiai10_gosseract_v2\all_test.go:52:12: undefined: NewClient
external\gazelle~~go_deps~com_github_otiai10_gosseract_v2\all_test.go:63:36: undefined: getDataPath
external\gazelle~~go_deps~com_github_otiai10_gosseract_v2\all_test.go:85:12: undefined: NewClient
external\gazelle~~go_deps~com_github_otiai10_gosseract_v2\all_test.go:92:12: undefined: NewClient
external\gazelle~~go_deps~com_github_otiai10_gosseract_v2\all_test.go:120:14: undefined: Client
external\gazelle~~go_deps~com_github_otiai10_gosseract_v2\all_test.go:120:14: too many errors
compilepkg: error running subcommand external\rules_go~~go_sdk~go_default_sdk\pkg\tool\windows_amd64\compile.exe: exit status 2
Target @@gazelle~~go_deps~com_github_otiai10_gosseract_v2//:gosseract_test failed to build                                                                                                                                          
Use --verbose_failures to see the command lines of failed build steps.                                                                                                                                                              
INFO: Elapsed time: 0.369s, Critical Path: 0.05s                                                                                                                                                                                    
INFO: 2 processes: 2 internal.                                                                                                                                                                                                      
ERROR: Build did NOT complete successfully
```

So, the test file cannot find symbols in the same package. I suspected this might have something
to do with cgo, so I created a [minimal example](../cgo/README.md) of trying
to use cgo locally and I was able to reproduce the behavior with a local pkg instead of an import.
