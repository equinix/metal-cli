# Release Instructions

These build and release instructions are intended for the maintainers and future maintainers of this project.

## Preparing a new version

There are no preperation steps.

* the version is computed from git tags
* The changelog is generated from git and lives outside of git

## Tagging

Pull the latest `master` branch and locally `git tag -s 0.0.9`.

When breaking changes are introduced bump the minor or major accordingly, restting the patch version.

## Releasing

Install goreleaser

Run:

```
export GITHUB_TOKEN=...
goreleaser
```

This will build and push binaries for several different OS and Architecture combinations.

Any special instructions or notes can be entered by editing the release notes at https://github.com/packethost/packet-cli/releases

