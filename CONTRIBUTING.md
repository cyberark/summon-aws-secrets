# Contributing

For general contribution and community guidelines, please see the [community repo](https://github.com/cyberark/community).

## Contributing

1. [Fork the project](https://help.github.com/en/github/getting-started-with-github/fork-a-repo)
2. [Clone your fork](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/cloning-a-repository)
3. Make local changes to your fork by editing files
3. [Commit your changes](https://help.github.com/en/github/managing-files-in-a-repository/adding-a-file-to-a-repository-using-the-command-line)
4. [Push your local changes to the remote server](https://help.github.com/en/github/using-git/pushing-commits-to-a-remote-repository)
5. [Create new Pull Request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request-from-a-fork)

From here your pull request will be reviewed and once you've responded to all
feedback it will be merged into the project. Congratulations, you're a
contributor!

## Releasing

### Update the version and changelog

1. Examine the changelog and decide on the version bump rank (major, minor, patch).
2. Change the title of _Unreleased_ section of the changelog to the target version.
   - Be sure to add the date (ISO 8601 format) to the section header.
3. Add a new, empty _Unreleased_ section to the changelog.
   - Remember to update the references at the bottom of the document.
4. Update `VERSION` and `version.go` files to reflect the version change.
5. Commit these changes. `Bump version to x.y.z` is an acceptable commit message.
6. Push your changes to a branch, and get the PR reviewed and merged.

### Tag the version

1. Tag the version on the master branch using `git tag -s vx.y.z -m vx.y.z`.
   Note this requires you to be able to sign releases. Consult the
   [GitHub documentation on signing commits](https://docs.github.com/en/authentication/managing-commit-signature-verification)
   to set this up.
2. Push the tag: `git push vx.y.z` (or `git push origin vx.y.z` if you are
   working from your local machine).

### Add a new GitHub release

1. Create a new release from the tag in the GitHub UI.
2. Add the changelog for the current version to the release description.
3. Retrieve the following artifacts from the Jenkins build on the tagged branch,
   and attach them to the release:
   - `SHA256SUMS.txt`
   - `summon-aws-secrets-darwin-amd64.tar.gz`
   - `summon-aws-secrets-darwin-arm64.tar.gz`
   - `summon-aws-secrets-freebsd-amd64.tar.gz`
   - `summon-aws-secrets-linux-amd64.tar.gz`
   - `summon-aws-secrets-netbsd-amd64.tar.gz`
   - `summon-aws-secrets-openbsd-amd64.tar.gz`
   - `summon-aws-secrets-solaris-amd64.tar.gz`
   - `summon-aws-secrets-windows-amd64.tar.gz`

### Update Homebrew Tools

1. Create a PR in [`cyberark/homebrew-tools`](https://github.com/cyberark/homebrew-tools)
   to update [`summon-aws-secrets.rb` formula](https://github.com/cyberark/homebrew-tools/blob/master/summon-aws-secrets.rb)
   using the file `dist/summon-aws-secrets.rb` retrieved from Jenkins artifacts.
   - Make sure the SHA hashes for the artifacts match the values in the
     `SHA256SUMS.txt` file attached to the release.
