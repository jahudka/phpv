# PHPv

Rudimentary PHP version manager for macOS

## What is this?

It's a bit like FNM for PHP, except somewhat simpler. It assumes some things
about your system:
 - You're using the default ZSH shell
 - You've installed PHP using Homebrew from the `shivammathur/php` tap

## What does it do?

It allows you to switch between PHP versions installed on your system.
Each active shell can have a different PHP version selected. Optionally,
you can also have PHPv hook into your shell to automatically switch PHP
versions when you `cd` into a folder containing a `composer.json` file
with a required PHP version specified.

## Installation

```shell
# Change to suit your preference
installDir="$HOME/.local/opt/phpv"

mkdir -p "${installDir}"
cd "${installDir}"
git clone git@github.com:jahudka/phpv.git .
composer install

# Install any PHP versions you need:
brew install shivammathur/php/php@7.1 shivammathur/php/php@8.4
```

> Homebrew will warn you that the installed PHP versions weren't
> symlinked to be globally available - this is fine, PHPv doesn't
> need them to be. You _can_ make a specific PHP version available
> globally even outside shell environments by running
> `brew link php@<version>`; PHPv will simply non-destructively
> override the global version when needed.

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
# Switch to a specific PHP version:
phpv 8.4

# Switch to the latest installed PHP version:
phpv latest
```

Running the command without any arguments will print a list of
the available PHP versions.

The automatic version switching when you `cd` into a directory
containing a `composer.json` file will select the lowest available
PHP version matching the requirement specified in the file.
The matching is resolved against the `<major>.<minor>` version
specified when installing PHP, not the full PHP version.

## How does it work?

It simply updates the `$PATH` variable in the current shell environment
to include the selected PHP version.
