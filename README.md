# GitHub as a file server 

Abuse GitHub unpublicized [attachment](https://docs.github.com/en/get-started/writing-on-github/working-with-advanced-formatting/attaching-files) API to serve a file.

Especially useful for hosting image files that can be referenced in markdown files.

## Usage

1. Install
    ```shell
    go install github.com/barn2plugins/github-s3/cmd/github-s3@latest
    ```
2. Login to your GitHub account, and obtain cookie named `user_session` from GitHub web browser session.
3. Run
    ```shell
    export GITHUB_SESSION=<github-user-session>   
    github-s3 [-repo owner/repo] [-pat ghp_1234567890abcdefghijklmnopqrstuvwyz0] <path-to-file>
    ```

### Parameters
The command supports two parameters:
 * `-repo` (optional): This parameter specifies the repository where you want to upload your assets, in the form `<owner>/<respository>` (if you skip this parameter, the `cli/cli` repository will be used)
 * `-pat` (optional): This parameter provides the Personal Access Token of the user uploading the assets. This parameter is only required if the repository specified with the `-repo` parameter is private. The Personal Access Token must be created with the `repo` and the `write:discussion` permissions

### Auto cookie detection

If you don't want to obtain the cookie manually, you can use [github-s3-auto](./cmd/github-s3-auto) to automatically find the cookie from your web browser session.

```shell
go install github.com/barn2plugins/github-s3/cmd/github-s3-auto@latest

github-s3-auto [-repo owner/repo] [-pat ghp_1234567890abcdefghijklmnopqrstuvwyz0] <path-to-file>
```

## Disclaimer

Please note that this project relies on an unpublicized API of GitHub, and its usage may be subject to changes in GitHub's policies or API. Use it responsibly and ensure compliance with GitHub's terms of service.
