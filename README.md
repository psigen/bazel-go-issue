# bazel-go-issue

A minimal example of an issue using gogo/status + gRPC in bazel.

```
$ bazel build //...
INFO: Invocation ID: ebfb834e-0eec-44be-8575-061d749594aa
INFO: Analyzed 4 targets (0 packages loaded, 0 targets configured).
INFO: Found 4 targets...
ERROR: /home/pras/kerfed/tmp/BUILD.bazel:36:10: GoLink bazel-go-issue_/bazel-go-issue failed: (Exit 1): builder failed: error executing command bazel-out/host/bin/external/go_sdk/builder link -sdk external/go_sdk -installsuffix linux_amd64 -arc ... (remaining 257 argument(s) skipped)

Use --sandbox_debug to see verbose messages from the sandbox builder failed: error executing command bazel-out/host/bin/external/go_sdk/builder link -sdk external/go_sdk -installsuffix linux_amd64 -arc ... (remaining 257 argument(s) skipped)

Use --sandbox_debug to see verbose messages from the sandbox
link: package conflict error: github.com/golang/protobuf/ptypes/any: multiple copies of package passed to linker:
	@io_bazel_rules_go//proto/wkt:any_go_proto
	@com_github_golang_protobuf//ptypes/any:any
Set "importmap" to different paths or use 'bazel cquery' to ensure only one
package with this path is linked.
INFO: Elapsed time: 0.383s, Critical Path: 0.21s
INFO: 2 processes: 2 internal.
FAILED: Build did NOT complete successfully
```
