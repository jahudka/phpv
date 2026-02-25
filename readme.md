# PHPv

Rudimentary PHP version manager for macOS

## What is this?

It's a bit like FNM for PHP, except somewhat simpler. It assumes some things
about your system:
 - You're using either Bash or ZSH
 - You've installed PHP using Homebrew from the `shivammathur/php` tap

## What does it do?

It allows you to switch between PHP versions installed on your system.
Each active shell can have a different PHP version selected. Optionally,
you can also have PHPv hook into your shell to automatically switch PHP
versions when you `cd` into a folder containing a `composer.json` file
with a required PHP version specified.

## Installation

```shell
# change to suit your preference
installDir="$HOME/.local/opt/phpv"

mkdir -p "${installDir}"
cd "${installDir}"
git clone git@github.com:jahudka/phpv.git .
composer install
```

Add PHPv initialization to your `~/.zprofile` or `~/.zshrc`.
Don't forget to change `$HOME/.local/opt/phpv` to the same value
you used for `$installDir` above. The `--with-chdir-hook` argument
enables automatic version switching when you `cd` into a directory
containing a `composer.json` file.

```shell
source "$HOME/.local/opt/phpv/src/setup-env" --with-chdir-hook
```

## Usage

```shell
# Switch to a specific version:
phpv <default|latest|$version>

# 'default' and 'latest' are synonymous
```
