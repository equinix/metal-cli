# Release Instructions

These build and release instructions are intended for the maintainers and future maintainers of this project.

## Preparing a new version

There are no preparation steps.

* the version is computed from [Conventional Commit](https://www.conventionalcommits.org/en/v1.0.0/) tags
* There is no changelog; the GitHub release notes are generated based on [Conventional Commit](https://www.conventionalcommits.org/en/v1.0.0/) tags 

## Releasing

Run the GitHub Actions [Release Workflow](.github/workflows/release.yml) on the `main` branch.

The release workflow:
- Uses [Semantic Release](.releaserc.json) to determine the next version number and create the GitHub release
- Uses [GoReleaser](.goreleaser.yml) to build the CLI binaries and attach them to the GitHub release
- Updates the Homebrew tap

This will build and release binaries for several different OS and Architecture combinations.

Any special instructions or notes should be added by editing the release notes that the workflow publishes. These notes can be found at https://github.com/equinix/metal-cli/releases

