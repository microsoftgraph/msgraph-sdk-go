package models
import (
    "errors"
    "strings"
)
// Product Status of Windows Defender
type WindowsDefenderProductStatus int

const (
    // No status
    NOSTATUS_WINDOWSDEFENDERPRODUCTSTATUS WindowsDefenderProductStatus = iota
    // Service not running
    SERVICENOTRUNNING_WINDOWSDEFENDERPRODUCTSTATUS
    // Service started without any malware protection engine
    SERVICESTARTEDWITHOUTMALWAREPROTECTION_WINDOWSDEFENDERPRODUCTSTATUS
    // Pending full scan due to threat action
    PENDINGFULLSCANDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
    // Pending reboot due to threat action
    PENDINGREBOOTDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
    // Pending manual steps due to threat action 
    PENDINGMANUALSTEPSDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
    // AV signatures out of date
    AVSIGNATURESOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
    // AS signatures out of date
    ASSIGNATURESOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
    // No quick scan has happened for a specified period
    NOQUICKSCANHAPPENEDFORSPECIFIEDPERIOD_WINDOWSDEFENDERPRODUCTSTATUS
    // No full scan has happened for a specified period
    NOFULLSCANHAPPENEDFORSPECIFIEDPERIOD_WINDOWSDEFENDERPRODUCTSTATUS
    // System initiated scan in progress
    SYSTEMINITIATEDSCANINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
    // System initiated clean in progress
    SYSTEMINITIATEDCLEANINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
    // There are samples pending submission
    SAMPLESPENDINGSUBMISSION_WINDOWSDEFENDERPRODUCTSTATUS
    // Product running in evaluation mode
    PRODUCTRUNNINGINEVALUATIONMODE_WINDOWSDEFENDERPRODUCTSTATUS
    // Product running in non-genuine Windows mode
    PRODUCTRUNNINGINNONGENUINEMODE_WINDOWSDEFENDERPRODUCTSTATUS
    // Product expired
    PRODUCTEXPIRED_WINDOWSDEFENDERPRODUCTSTATUS
    // Off-line scan required
    OFFLINESCANREQUIRED_WINDOWSDEFENDERPRODUCTSTATUS
    // Service is shutting down as part of system shutdown
    SERVICESHUTDOWNASPARTOFSYSTEMSHUTDOWN_WINDOWSDEFENDERPRODUCTSTATUS
    // Threat remediation failed critically
    THREATREMEDIATIONFAILEDCRITICALLY_WINDOWSDEFENDERPRODUCTSTATUS
    // Threat remediation failed non-critically
    THREATREMEDIATIONFAILEDNONCRITICALLY_WINDOWSDEFENDERPRODUCTSTATUS
    // No status flags set (well initialized state)
    NOSTATUSFLAGSSET_WINDOWSDEFENDERPRODUCTSTATUS
    // Platform is out of date
    PLATFORMOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
    // Platform update is in progress
    PLATFORMUPDATEINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
    // Platform is about to be outdated
    PLATFORMABOUTTOBEOUTDATED_WINDOWSDEFENDERPRODUCTSTATUS
    // Signature or platform end of life is past or is impending
    SIGNATUREORPLATFORMENDOFLIFEISPASTORISIMPENDING_WINDOWSDEFENDERPRODUCTSTATUS
    // Windows SMode signatures still in use on non-Win10S install
    WINDOWSSMODESIGNATURESINUSEONNONWIN10SINSTALL_WINDOWSDEFENDERPRODUCTSTATUS
)

func (i WindowsDefenderProductStatus) String() string {
    var values []string
    for p := WindowsDefenderProductStatus(1); p <= WINDOWSSMODESIGNATURESINUSEONNONWIN10SINSTALL_WINDOWSDEFENDERPRODUCTSTATUS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"noStatus", "serviceNotRunning", "serviceStartedWithoutMalwareProtection", "pendingFullScanDueToThreatAction", "pendingRebootDueToThreatAction", "pendingManualStepsDueToThreatAction", "avSignaturesOutOfDate", "asSignaturesOutOfDate", "noQuickScanHappenedForSpecifiedPeriod", "noFullScanHappenedForSpecifiedPeriod", "systemInitiatedScanInProgress", "systemInitiatedCleanInProgress", "samplesPendingSubmission", "productRunningInEvaluationMode", "productRunningInNonGenuineMode", "productExpired", "offlineScanRequired", "serviceShutdownAsPartOfSystemShutdown", "threatRemediationFailedCritically", "threatRemediationFailedNonCritically", "noStatusFlagsSet", "platformOutOfDate", "platformUpdateInProgress", "platformAboutToBeOutdated", "signatureOrPlatformEndOfLifeIsPastOrIsImpending", "windowsSModeSignaturesInUseOnNonWin10SInstall"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseWindowsDefenderProductStatus(v string) (any, error) {
    var result WindowsDefenderProductStatus
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "noStatus":
                result |= NOSTATUS_WINDOWSDEFENDERPRODUCTSTATUS
            case "serviceNotRunning":
                result |= SERVICENOTRUNNING_WINDOWSDEFENDERPRODUCTSTATUS
            case "serviceStartedWithoutMalwareProtection":
                result |= SERVICESTARTEDWITHOUTMALWAREPROTECTION_WINDOWSDEFENDERPRODUCTSTATUS
            case "pendingFullScanDueToThreatAction":
                result |= PENDINGFULLSCANDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
            case "pendingRebootDueToThreatAction":
                result |= PENDINGREBOOTDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
            case "pendingManualStepsDueToThreatAction":
                result |= PENDINGMANUALSTEPSDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
            case "avSignaturesOutOfDate":
                result |= AVSIGNATURESOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
            case "asSignaturesOutOfDate":
                result |= ASSIGNATURESOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
            case "noQuickScanHappenedForSpecifiedPeriod":
                result |= NOQUICKSCANHAPPENEDFORSPECIFIEDPERIOD_WINDOWSDEFENDERPRODUCTSTATUS
            case "noFullScanHappenedForSpecifiedPeriod":
                result |= NOFULLSCANHAPPENEDFORSPECIFIEDPERIOD_WINDOWSDEFENDERPRODUCTSTATUS
            case "systemInitiatedScanInProgress":
                result |= SYSTEMINITIATEDSCANINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
            case "systemInitiatedCleanInProgress":
                result |= SYSTEMINITIATEDCLEANINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
            case "samplesPendingSubmission":
                result |= SAMPLESPENDINGSUBMISSION_WINDOWSDEFENDERPRODUCTSTATUS
            case "productRunningInEvaluationMode":
                result |= PRODUCTRUNNINGINEVALUATIONMODE_WINDOWSDEFENDERPRODUCTSTATUS
            case "productRunningInNonGenuineMode":
                result |= PRODUCTRUNNINGINNONGENUINEMODE_WINDOWSDEFENDERPRODUCTSTATUS
            case "productExpired":
                result |= PRODUCTEXPIRED_WINDOWSDEFENDERPRODUCTSTATUS
            case "offlineScanRequired":
                result |= OFFLINESCANREQUIRED_WINDOWSDEFENDERPRODUCTSTATUS
            case "serviceShutdownAsPartOfSystemShutdown":
                result |= SERVICESHUTDOWNASPARTOFSYSTEMSHUTDOWN_WINDOWSDEFENDERPRODUCTSTATUS
            case "threatRemediationFailedCritically":
                result |= THREATREMEDIATIONFAILEDCRITICALLY_WINDOWSDEFENDERPRODUCTSTATUS
            case "threatRemediationFailedNonCritically":
                result |= THREATREMEDIATIONFAILEDNONCRITICALLY_WINDOWSDEFENDERPRODUCTSTATUS
            case "noStatusFlagsSet":
                result |= NOSTATUSFLAGSSET_WINDOWSDEFENDERPRODUCTSTATUS
            case "platformOutOfDate":
                result |= PLATFORMOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
            case "platformUpdateInProgress":
                result |= PLATFORMUPDATEINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
            case "platformAboutToBeOutdated":
                result |= PLATFORMABOUTTOBEOUTDATED_WINDOWSDEFENDERPRODUCTSTATUS
            case "signatureOrPlatformEndOfLifeIsPastOrIsImpending":
                result |= SIGNATUREORPLATFORMENDOFLIFEISPASTORISIMPENDING_WINDOWSDEFENDERPRODUCTSTATUS
            case "windowsSModeSignaturesInUseOnNonWin10SInstall":
                result |= WINDOWSSMODESIGNATURESINUSEONNONWIN10SINSTALL_WINDOWSDEFENDERPRODUCTSTATUS
            default:
                return 0, errors.New("Unknown WindowsDefenderProductStatus value: " + v)
        }
    }
    return &result, nil
}
func SerializeWindowsDefenderProductStatus(values []WindowsDefenderProductStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsDefenderProductStatus) isMultiValue() bool {
    return true
}
