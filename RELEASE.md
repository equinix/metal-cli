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

Push the tag to GitHub and [GitHub Workflows](.github/workflows/release.yml) and [GoReleaser](.goreleaser.yml) will do the rest.

```sh
git push origin --tags 0.0.9
```

This will build and release binaries for several different OS and Architecture combinations.

Any special instructions or notes should be added by editing the release notes that goreleaser publishes. These notes can be found at https://github.com/packethost/packet-cli/releases

