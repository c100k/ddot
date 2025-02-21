# ddot

A utility tool to create short-living and disposable `.env` files from your Password Manager.

<p align="center">
    <a href="https://www.youtube.com/watch?v=voQ1STMNT40" rel="noopener noreferrer" target="_blank">
        <img alt="ddot demo with bitwarden and 1Password" src="https://img.youtube.com/vi/voQ1STMNT40/0.jpg" width="600">
    </a>
</p>

> [!IMPORTANT]
> This is a hobby project. It's developed by only one person and should be considered as is. It may keep going for years or stop at any time.
> The source code will probably be published when the structure and documentation is state of the art.

## Why ?

According to the [3rd rule of the twelve-factor app](https://12factor.net/config), all the application config that depends on the environment (`dev`, `staging`, `prod`...) must not be stored as constants in the code. Instead, it must be stored in environment variables.

During local development, developers usually rely on `.env` files, whether they're using [Node.js](https://nodejs.org/dist/latest/docs/api/cli.html#--env-fileconfig), [Ruby on Rails](https://github.com/bkeepers/dotenv), [Go](https://github.com/joho/godotenv), [Python](https://pypi.org/project/python-dotenv), [Rust](https://github.com/allan2/dotenvy) or anything else.

Having this clear text file lying around on the disk can have huge security impacts, especially when it contains sensitive secrets, which are often the perfect candidates for environment-specific config.

Indeed, some attackers target developers in order to scan their disk to exfiltrate the contents of these files. When they contain **AWS credentials with root access** or **crypto wallets keys**, it can become a **big problem**.

## Getting Started

`ddot` runs on your machine and is compatible with the following Password Managers : [bitwarden](https://bitwarden.com) and [1Password](https://1password.com).

First, download the [latest release](https://github.com/c100k/ddot/releases).

Then, get inspiration from the commands below to define the ones that fit your own workflow.

```sh
# create a .env file from the contents of .env.base (fine as long as it does not contain non-sensitive secrets)
ddot loadenv --uri file://.env.base

# create a .env file from the contents of the bitwarden secure note named `myapp-dev`
ddot loadenv --uri bw://myapp-dev

# create a .env file from the contents of the 1Password secure note named `myapp-dev`
ddot loadenv --uri op://myapp-dev

# create a .env file from the contents of all the resources listed
ddot loadenv --uri file://.env.base --uri bw://myapp-dev --uri op://myapp-dev

# create a .env.production file from the contents of all the resources listed
ddot loadenv --out .env.production --uri file://.env.base --uri bw://myapp-prod --uri op://myapp-prod
```

As your can see, `ddot` can combine multiple resources into one `.env` file. For each provider, see the ad-hoc subsection below, with specific setup instructions.

The commands above create the `.env` file and whenever you press <kbd>ctrl</kbd> + <kbd>C</kbd>, it is automatically deleted.

## Providers

Currently, there are 3 providers available.

> [!NOTE]
> Some Password Managers offer features to help securing secrets during local development (e.g. Secrets Manager for bitwarden and the `inject` command for 1Password).
> These are great solutions. The purpose of `ddot` is to offer an alternate solution (choice is always good) and most of all, a provider-agnostic one. By the way, it does not require you to change anything in your codebase.

### File

It works out of the box, nothing specific to install.

### Bitwarden

Under the hood, it calls the `bw` cli. Install it by following the instructions of the [official documentation](https://bitwarden.com/help/cli/#download-and-install).

Check your install.

```sh
bw --version
```

### 1Password

Under the hood, it calls the `op` cli. Install it by following the instructions of the [official documentation](https://developer.1password.com/docs/cli/get-started/#step-1-install-1password-cli).

Make sure to follow the `Turn on the 1Password desktop app integration` section in order to unlock your vault with `Touch ID` when invoking `ddot`.

Check your install.

```sh
op --version
```

## Run from source

```sh
git clone git@github.com:c100k/ddot.git

go fmt && go build

./ddot version
./ddot loadenv --uri file://.env.example.1
./ddot loadenv --uri file://.env.example.1 --uri file://.env.example.2
```
