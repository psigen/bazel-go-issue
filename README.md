# bazel-go-issue
A minimal example of an issue using gogo/status in bazel.

```
$ bazel build //...
INFO: Invocation ID: b2095db6-98ae-41f3-b505-8597aa444779
INFO: Analyzed 2 targets (0 packages loaded, 0 targets configured).
INFO: Found 2 targets...
ERROR: /home/pras/.cache/bazel/_bazel_pras/7ff1d612b81fb750bcd6f1a5fd699714/external/com_github_gogo_status/BUILD.bazel:3:11: GoCompilePkg external/com_github_gogo_status/status.a failed: (Exit 1): builder failed: error executing command bazel-out/host/bin/external/go_sdk/builder compilepkg -sdk external/go_sdk -installsuffix linux_amd64 -src external/com_github_gogo_status/status.go -arc ... (remaining 27 argument(s) skipped)

Use --sandbox_debug to see verbose messages from the sandbox builder failed: error executing command bazel-out/host/bin/external/go_sdk/builder compilepkg -sdk external/go_sdk -installsuffix linux_amd64 -src external/com_github_gogo_status/status.go -arc ... (remaining 27 argument(s) skipped)

Use --sandbox_debug to see verbose messages from the sandbox
external/com_github_gogo_status/status.go:158:22: cannot use &types.Any literal (type *types.Any) as type *any.Any in append
external/com_github_gogo_status/status.go:186:21: cannot use any (type *types.Any) as type *any.Any in append
external/com_github_gogo_status/status.go:200:31: cannot use any (type *any.Any) as type *types.Any in argument to types.UnmarshalAny
compilepkg: error running subcommand external/go_sdk/pkg/tool/linux_amd64/compile: exit status 2
INFO: Elapsed time: 0.394s, Critical Path: 0.27s
INFO: 2 processes: 2 internal.
FAILED: Build did NOT complete successfully
```
