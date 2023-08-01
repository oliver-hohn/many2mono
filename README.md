# many2mono

`many2mono` is a CLI tool for merging many repos into an existing repo, while still keeping the commit history.

It merges the passed in repos into the current working directory as sub-folders, e.g.
```sh
cd repo_a/
ls # => app/

many2mono run repo_b/ repo_c/
ls # => app/ repo_b/ repo_c/
```

It is based on the work from:
- https://jeffkreeftmeijer.com/git-combine/
- https://alexharv074.github.io/puppet/2017/10/04/merge-a-git-repository-and-its-history-into-a-subdirectory-of-a-second-git-repository.html

## Installation
1. Download the latest release: https://github.com/oliver-hohn/woodhouse/releases
    - Note: `darwin` is for Mac OS (Intel and Apple Silicon)
1. Place it in your `PATH`:
   ```sh
   # sudo may be needed
   mv ~/Downloads/many2mono_x.x.x_x/many2mono /usr/local/bin/
   ```
1. Check the installation:
   ```
   many2mono --help
   ```
   _On Mac OS, you may see an error: "cannot be opened because the developer cannot be verified". To authorize the package go to "System Settings" > "Privacy & Security" > "Security" > "Allow Anyway". Re-run the `--help` command to confirm the installation is correct_.

## Usage
1. `cd` into the repo you want to merge the other repos into:
   ```sh
   cd repo_a/
   ```
1. (_Optional_): Check out a new branch:
   ```
   git checkout -b merging_repos
   ```
1. Merge the other repos using their SSH URL (_this avoids having to pass in credentials for `many2mono` to fetch the other repos_):
   ```sh
   many2mono run git@github.com:USERNAME/REPO_B.git git@github.com:USERNAME/REPO_C.git
   ```
   - The `--dry-run` option can be used to inspect what operations `many2mono` will do, without actually doing them.
   - By default, `many2mono` merges the `main` branch from the repos passed in. If another branch is needed it can be passed in using: `-b master`.
     - _Note: This branch will be used for all repos being merged in. If you only want to use it for one of the repos, then merge that repo in separately, e.g._
       ```sh
       many2mono run git@github.com:USERNAME/REPO_B.git git@github.com:USERNAME/REPO_C.git
       many2mono run -b master git@github.com:USERNAME/REPO_D.git
       ```

## Considerations
- When merging repos, the commit history is kept, however, when viewing the files from the merged in repos, these will not have their previous commit history. However, the file will still have their old blame history.
- `many2mono` only supports SSH URLs for the repos being merged in. Support for other paths may be added in the future.