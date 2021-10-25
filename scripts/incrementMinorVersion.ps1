# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License.

<#
.Synopsis
    Increment the minor version string in the gradle.properties if the major,
    minor, or patch version hasn't been manually updated.
.Description
    Assumptions:
        This script assumes it is run from the repo root.
        Minor version is typically auto-incremented.

#>

function Update-TelemetryVersion([string]$telemetryFilePath, [version]$version) {
	$telemetryFileContent = Get-Content -Path $telemetryFilePath -Raw
	$telemetryFileContent = $telemetryFileContent -replace "\d{1,}\.\d{1,}\.\d{1,}", $version.ToString()
	Set-Content -Path $telemetryFilePath $telemetryFileContent
}

function Get-CurrentTelemetryVersion([string]$telemetryFilePath) {
	$telemetryFileContent = Get-Content -Path $telemetryFilePath -Raw
	if($telemetryFileContent -match "(\d{1,}\.\d{1,}\.\d{1,})") {
			return [version]::Parse($Matches[1])
	} else {
			Write-Error "Invalid version number format"
			return $null;
	}
}

function Update-MinorVersionNumber([version]$currentVersion) {
	return [version]::new($currentVersion.Major, $currentVersion.Minor + 1, 0);
}

function Update-MinorVersion() {
	$telemetryFilePath = Join-Path -Path $PWD.ToString() -ChildPath "../graph_request_adapter.go"
	$currentVersion = Get-CurrentTelemetryVersion -telemetryFilePath $telemetryFilePath
	$nextVersion = Update-MinorVersionNumber -currentVersion $currentVersion
	Update-TelemetryVersion -version $nextVersion -telemetryFilePath $telemetryFilePath
}
Update-MinorVersion