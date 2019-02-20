# summon-aws-secrets
[Summon](https://github.com/cyberark/summon) provider for AWS Secrets Manager

## Install

Use the auto-install script. This will install the latest version of summon-aws-secrets.
The script requires sudo to place summon-aws-secrets in `/usr/local/lib/summon`.

```
curl -sSL https://raw.githubusercontent.com/cyberark/summon-aws-secrets/master/install.sh | bash
```

Otherwise, download the [latest release](https://github.com/cyberark/summon-aws-secrets/releases) and extract it to the directory `/usr/local/lib/summon`.

## Usage in isolation

Give summon-aws-secrets a variable name and it will fetch it for you and print the value to stdout.

```sh-session
$ # Configure in similar fashion to AWS CLI see https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html
$ summon-aws-secrets prod/aws/iam/user/robot/access_key_id
8h9psadf89sdahfp98
```

AWS Secrets Manager also supports multiple key value pairs stored as json blob. If you use this approach, it is possible to select the value of any of the key/value pairs by specifying a column followed by the key name:

Example Secret JSON:
```json
{
  "user-1": "password-1",
  "user-2": "password-2",
  "user-3": "password-3"
}
```

```bash
$ summon-aws-secrets prod/aws/iam/user/robot/access_key_id:user-2
{ "user-1": "password-1", "user-2": "password-2", "user-3": "password-3"}

$ summon-aws-secrets prod/aws/iam/user/robot/access_key_id:user-2
password-2
```

### Flags

`summon-aws-secrets` supports a single flag.

* `-v, --version` Output version number and quit

## Usage as a provider for Summon

[Summon](https://github.com/cyberark/summon/) is a command-line tool that reads a file in secrets.yml format and injects secrets as environment variables into any process. Once the process exits, the secrets are gone.

*Example*

As an example let's use the `env` command: 

Following installation, define your keys in a `secrets.yml` file

```yml
AWS_ACCESS_KEY_ID: !var aws/iam/user/robot/access_key_id
AWS_SECRET_ACCESS_KEY: !var aws/iam/user/robot/secret_access_key
```

By default, summon will look for `secrets.yml` in the directory it is called from and export the secret values to the environment of the command it wraps.

Wrap the `env` in summon:

```sh
$ # Configure in similar fashion to AWS CLI see https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html
$ summon --provider summon-aws-secrets env
...
AWS_ACCESS_KEY_ID=AKIAJS34242K1123J3K43
AWS_SECRET_ACCESS_KEY=A23MSKSKSJASHDIWM
...
```

`summon` resolves the entries in secrets.yml with the AWS Secrets Manager provider and makes the secret values available to the environment of the command `env`.

## Configuration

This provider uses the same configuration pattern as the [AWS CLI
](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html) to connect to AWS.
