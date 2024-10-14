# go-cli-demo
A rudimentary setup for a GO CLI, leveraging spf13/cobra

## Prerequisites
 - Go (https://go.dev/doc/install)
 - Azure CLI (https://learn.microsoft.com/en-us/cli/azure/install-azure-cli)

At last, the executable/binary must be created via build-win.ps1 or build-linux.sh respectively. and can be optionally added to
the PATH environment variable</br>

## Usage
Print a message to the console by executing `./helper hello` in the project's root directory.

## Use login sub-command
Populate the required config file for [configs/dev.yaml](configs/dev.yaml) or [configs/prod.yaml](configs/prod.yaml) respectively.</br>
Load configs via install-configs.ps1 or install-configs.sh.</br>
Execute with `./helper login -e dev` or `./helper login -e prod` respectively
