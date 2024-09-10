# go-cli-demo
A rudamentory setup for a GO CLI, leveraging spf13/cobra

Create executable/binary via build-win.ps1 or build-linux.sh respectively.
Execute with ./helper hello

## Use login sub-command
Populate required config file configs/dev.yaml or configs/prod.yaml respecitvely
Load configs for via install-configs.ps1 or install-configs.sh
Execute with ./helper login -e dev or ./helper login -e prod respecitvely