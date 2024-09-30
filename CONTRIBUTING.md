# Contributing

For general contribution and community guidelines, please see the [community repo](https://github.com/cyberark/community).

## Contributing

1. [Fork the project](https://help.github.com/en/github/getting-started-with-github/fork-a-repo)
2. [Clone your fork](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/cloning-a-repository)
3. Make local changes to your fork by editing files
4. [Commit your changes](https://help.github.com/en/github/managing-files-in-a-repository/adding-a-file-to-a-repository-using-the-command-line)
5. [Push your local changes to the remote server](https://help.github.com/en/github/using-git/pushing-commits-to-a-remote-repository)
6. [Create new Pull Request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request-from-a-fork)

From here your pull request will be reviewed and once you've responded to all
feedback it will be merged into the project. Congratulations, you're a
contributor!

## Releasing

Releases should be created by maintainers only. To create and promote a
release, follow the instructions in this section.

### Update the version and changelog

The following checklist should be followed when creating a release:

- [ ] Open a PR with the following changes for `main` branch and wait for it to be merged:
  - [ ] Bump version in [`CHANGELOG.md`](CHANGELOG.md).
  - [ ] Update NOTICES.txt

    ```sh-session
    go install github.com/google/go-licenses@latest
    # Verify that dependencies fit into supported licenses types. If there is new dependency having unsupported license,
    # that license should be included to notices.tpl file in order to get generated in NOTICES.txt.
    $(go env GOPATH)/bin/go-licenses check ./... --allowed_licenses="MIT,ISC,Apache-2.0,BSD-3-Clause,BSD-2-Clause"
    # If no errors occur, proceed to generate updated NOTICES.txt
    $(go env GOPATH)/bin/go-licenses report ./... --template notices.tpl > NOTICES.txt --ignore github.com/cyberark/summon-aws-secrets
   ```

### Release and Promote

1. Merging into the main branch will automatically trigger a release.
   If successful, this release can be promoted at a later time.

2. Jenkins build parameters can be utilized to promote a successful release
   or manually trigger aditional releases as needed.

3. Reference the
   [internal automated release doc](https://github.com/conjurinc/docs/blob/master/reference/infrastructure/automated_releases.md#release-and-promotion-process)
   for releasing and promoting.
