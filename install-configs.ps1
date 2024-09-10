$destinationFolder = "~/.helper"
if (!(Test-Path -path $destinationFolder)) {New-Item $destinationFolder -Type Directory}
Write-Host "rm -Recurse $destinationFolder"
rm -Recurse $destinationFolder
Write-Host "cp -Recurse configs $destinationFolder"
cp -Recurse configs $destinationFolder