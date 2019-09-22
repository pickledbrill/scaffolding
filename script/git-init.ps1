param (
    [Parameter(Mandatory=$true)][string]
    $giturl,
    [Parameter(Mandatory=$true)][string]
    $workdir
)

Write-Host $giturl
Write-Host $workdir

git init -C $workdir
git -C $workdir remote add origin $giturl