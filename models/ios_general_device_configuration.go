package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosGeneralDeviceConfiguration 
type IosGeneralDeviceConfiguration struct {
    DeviceConfiguration
    // Indicates whether or not to allow account modification when the device is in supervised mode.
    accountBlockModification *bool
    // Indicates whether or not to allow activation lock when the device is in the supervised mode.
    activationLockAllowWhenSupervised *bool
    // Indicates whether or not to allow AirDrop when the device is in supervised mode.
    airDropBlocked *bool
    // Indicates whether or not to cause AirDrop to be considered an unmanaged drop target (iOS 9.0 and later).
    airDropForceUnmanagedDropTarget *bool
    // Indicates whether or not to enforce all devices receiving AirPlay requests from this device to use a pairing password.
    airPlayForcePairingPasswordForOutgoingRequests *bool
    // Indicates whether or not to block the user from using News when the device is in supervised mode (iOS 9.0 and later).
    appleNewsBlocked *bool
    // Indicates whether or not to allow Apple Watch pairing when the device is in supervised mode (iOS 9.0 and later).
    appleWatchBlockPairing *bool
    // Indicates whether or not to force a paired Apple Watch to use Wrist Detection (iOS 8.2 and later).
    appleWatchForceWristDetection *bool
    // Gets or sets the list of iOS apps allowed to autonomously enter Single App Mode. Supervised only. iOS 7.0 and later. This collection can contain a maximum of 500 elements.
    appsSingleAppModeList []AppListItemable
    // Indicates whether or not to block the automatic downloading of apps purchased on other devices when the device is in supervised mode (iOS 9.0 and later).
    appStoreBlockAutomaticDownloads *bool
    // Indicates whether or not to block the user from using the App Store. Requires a supervised device for iOS 13 and later.
    appStoreBlocked *bool
    // Indicates whether or not to block the user from making in app purchases.
    appStoreBlockInAppPurchases *bool
    // Indicates whether or not to block the App Store app, not restricting installation through Host apps. Applies to supervised mode only (iOS 9.0 and later).
    appStoreBlockUIAppInstallation *bool
    // Indicates whether or not to require a password when using the app store.
    appStoreRequirePassword *bool
    // List of apps in the visibility list (either visible/launchable apps list or hidden/unlaunchable apps list, controlled by AppsVisibilityListType) (iOS 9.3 and later). This collection can contain a maximum of 10000 elements.
    appsVisibilityList []AppListItemable
    // Type of list that is in the AppsVisibilityList. Possible values are: none, appsInListCompliant, appsNotInListCompliant.
    appsVisibilityListType *AppListType
    // Indicates whether or not to allow modification of Bluetooth settings when the device is in supervised mode (iOS 10.0 and later).
    bluetoothBlockModification *bool
    // Indicates whether or not to block the user from accessing the camera of the device. Requires a supervised device for iOS 13 and later.
    cameraBlocked *bool
    // Indicates whether or not to block data roaming.
    cellularBlockDataRoaming *bool
    // Indicates whether or not to block global background fetch while roaming.
    cellularBlockGlobalBackgroundFetchWhileRoaming *bool
    // Indicates whether or not to allow changes to cellular app data usage settings when the device is in supervised mode.
    cellularBlockPerAppDataModification *bool
    // Indicates whether or not to block Personal Hotspot.
    cellularBlockPersonalHotspot *bool
    // Indicates whether or not to block voice roaming.
    cellularBlockVoiceRoaming *bool
    // Indicates whether or not to block untrusted TLS certificates.
    certificatesBlockUntrustedTlsCertificates *bool
    // Indicates whether or not to allow remote screen observation by Classroom app when the device is in supervised mode (iOS 9.3 and later).
    classroomAppBlockRemoteScreenObservation *bool
    // Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to view a student's screen without prompting when the device is in supervised mode.
    classroomAppForceUnpromptedScreenObservation *bool
    // List that is in the AppComplianceList. Possible values are: none, appsInListCompliant, appsNotInListCompliant.
    compliantAppListType *AppListType
    // List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
    compliantAppsList []AppListItemable
    // Indicates whether or not to block the user from installing configuration profiles and certificates interactively when the device is in supervised mode.
    configurationProfileBlockChanges *bool
    // Indicates whether or not to block definition lookup when the device is in supervised mode (iOS 8.1.3 and later ).
    definitionLookupBlocked *bool
    // Indicates whether or not to allow the user to enables restrictions in the device settings when the device is in supervised mode.
    deviceBlockEnableRestrictions *bool
    // Indicates whether or not to allow the use of the 'Erase all content and settings' option on the device when the device is in supervised mode.
    deviceBlockEraseContentAndSettings *bool
    // Indicates whether or not to allow device name modification when the device is in supervised mode (iOS 9.0 and later).
    deviceBlockNameModification *bool
    // Indicates whether or not to block diagnostic data submission.
    diagnosticDataBlockSubmission *bool
    // Indicates whether or not to allow diagnostics submission settings modification when the device is in supervised mode (iOS 9.3.2 and later).
    diagnosticDataBlockSubmissionModification *bool
    // Indicates whether or not to block the user from viewing managed documents in unmanaged apps.
    documentsBlockManagedDocumentsInUnmanagedApps *bool
    // Indicates whether or not to block the user from viewing unmanaged documents in managed apps.
    documentsBlockUnmanagedDocumentsInManagedApps *bool
    // An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
    emailInDomainSuffixes []string
    // Indicates whether or not to block the user from trusting an enterprise app.
    enterpriseAppBlockTrust *bool
    // [Deprecated] Configuring this setting and setting the value to 'true' has no effect on the device.
    enterpriseAppBlockTrustModification *bool
    // Indicates whether or not to block the user from using FaceTime. Requires a supervised device for iOS 13 and later.
    faceTimeBlocked *bool
    // Indicates whether or not to block changes to Find My Friends when the device is in supervised mode.
    findMyFriendsBlocked *bool
    // Indicates whether or not to block the user from using Game Center when the device is in supervised mode.
    gameCenterBlocked *bool
    // Indicates whether or not to block the user from having friends in Game Center. Requires a supervised device for iOS 13 and later.
    gamingBlockGameCenterFriends *bool
    // Indicates whether or not to block the user from using multiplayer gaming. Requires a supervised device for iOS 13 and later.
    gamingBlockMultiplayer *bool
    // indicates whether or not to allow host pairing to control the devices an iOS device can pair with when the iOS device is in supervised mode.
    hostPairingBlocked *bool
    // Indicates whether or not to block the user from using the iBooks Store when the device is in supervised mode.
    iBooksStoreBlocked *bool
    // Indicates whether or not to block the user from downloading media from the iBookstore that has been tagged as erotica.
    iBooksStoreBlockErotica *bool
    // Indicates whether or not to block the user from continuing work they started on iOS device to another iOS or macOS device.
    iCloudBlockActivityContinuation *bool
    // Indicates whether or not to block iCloud backup. Requires a supervised device for iOS 13 and later.
    iCloudBlockBackup *bool
    // Indicates whether or not to block iCloud document sync. Requires a supervised device for iOS 13 and later.
    iCloudBlockDocumentSync *bool
    // Indicates whether or not to block Managed Apps Cloud Sync.
    iCloudBlockManagedAppsSync *bool
    // Indicates whether or not to block iCloud Photo Library.
    iCloudBlockPhotoLibrary *bool
    // Indicates whether or not to block iCloud Photo Stream Sync.
    iCloudBlockPhotoStreamSync *bool
    // Indicates whether or not to block Shared Photo Stream.
    iCloudBlockSharedPhotoStream *bool
    // Indicates whether or not to require backups to iCloud be encrypted.
    iCloudRequireEncryptedBackup *bool
    // Indicates whether or not to block the user from accessing explicit content in iTunes and the App Store. Requires a supervised device for iOS 13 and later.
    iTunesBlockExplicitContent *bool
    // Indicates whether or not to block Music service and revert Music app to classic mode when the device is in supervised mode (iOS 9.3 and later and macOS 10.12 and later).
    iTunesBlockMusicService *bool
    // Indicates whether or not to block the user from using iTunes Radio when the device is in supervised mode (iOS 9.3 and later).
    iTunesBlockRadio *bool
    // Indicates whether or not to block keyboard auto-correction when the device is in supervised mode (iOS 8.1.3 and later).
    keyboardBlockAutoCorrect *bool
    // Indicates whether or not to block the user from using dictation input when the device is in supervised mode.
    keyboardBlockDictation *bool
    // Indicates whether or not to block predictive keyboards when device is in supervised mode (iOS 8.1.3 and later).
    keyboardBlockPredictive *bool
    // Indicates whether or not to block keyboard shortcuts when the device is in supervised mode (iOS 9.0 and later).
    keyboardBlockShortcuts *bool
    // Indicates whether or not to block keyboard spell-checking when the device is in supervised mode (iOS 8.1.3 and later).
    keyboardBlockSpellCheck *bool
    // Indicates whether or not to allow assistive speak while in kiosk mode.
    kioskModeAllowAssistiveSpeak *bool
    // Indicates whether or not to allow access to the Assistive Touch Settings while in kiosk mode.
    kioskModeAllowAssistiveTouchSettings *bool
    // Indicates whether or not to allow device auto lock while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockAutoLock instead.
    kioskModeAllowAutoLock *bool
    // Indicates whether or not to allow access to the Color Inversion Settings while in kiosk mode.
    kioskModeAllowColorInversionSettings *bool
    // Indicates whether or not to allow use of the ringer switch while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockRingerSwitch instead.
    kioskModeAllowRingerSwitch *bool
    // Indicates whether or not to allow screen rotation while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockScreenRotation instead.
    kioskModeAllowScreenRotation *bool
    // Indicates whether or not to allow use of the sleep button while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockSleepButton instead.
    kioskModeAllowSleepButton *bool
    // Indicates whether or not to allow use of the touchscreen while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockTouchscreen instead.
    kioskModeAllowTouchscreen *bool
    // Indicates whether or not to allow access to the voice over settings while in kiosk mode.
    kioskModeAllowVoiceOverSettings *bool
    // Indicates whether or not to allow use of the volume buttons while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockVolumeButtons instead.
    kioskModeAllowVolumeButtons *bool
    // Indicates whether or not to allow access to the zoom settings while in kiosk mode.
    kioskModeAllowZoomSettings *bool
    // URL in the app store to the app to use for kiosk mode. Use if KioskModeManagedAppId is not known.
    kioskModeAppStoreUrl *string
    // ID for built-in apps to use for kiosk mode. Used when KioskModeManagedAppId and KioskModeAppStoreUrl are not set.
    kioskModeBuiltInAppId *string
    // Managed app id of the app to use for kiosk mode. If KioskModeManagedAppId is specified then KioskModeAppStoreUrl will be ignored.
    kioskModeManagedAppId *string
    // Indicates whether or not to require assistive touch while in kiosk mode.
    kioskModeRequireAssistiveTouch *bool
    // Indicates whether or not to require color inversion while in kiosk mode.
    kioskModeRequireColorInversion *bool
    // Indicates whether or not to require mono audio while in kiosk mode.
    kioskModeRequireMonoAudio *bool
    // Indicates whether or not to require voice over while in kiosk mode.
    kioskModeRequireVoiceOver *bool
    // Indicates whether or not to require zoom while in kiosk mode.
    kioskModeRequireZoom *bool
    // Indicates whether or not to block the user from using control center on the lock screen.
    lockScreenBlockControlCenter *bool
    // Indicates whether or not to block the user from using the notification view on the lock screen.
    lockScreenBlockNotificationView *bool
    // Indicates whether or not to block the user from using passbook when the device is locked.
    lockScreenBlockPassbook *bool
    // Indicates whether or not to block the user from using the Today View on the lock screen.
    lockScreenBlockTodayView *bool
    // Media content rating settings for Apps. Possible values are: allAllowed, allBlocked, agesAbove4, agesAbove9, agesAbove12, agesAbove17.
    mediaContentRatingApps *RatingAppsType
    // Media content rating settings for Australia
    mediaContentRatingAustralia MediaContentRatingAustraliaable
    // Media content rating settings for Canada
    mediaContentRatingCanada MediaContentRatingCanadaable
    // Media content rating settings for France
    mediaContentRatingFrance MediaContentRatingFranceable
    // Media content rating settings for Germany
    mediaContentRatingGermany MediaContentRatingGermanyable
    // Media content rating settings for Ireland
    mediaContentRatingIreland MediaContentRatingIrelandable
    // Media content rating settings for Japan
    mediaContentRatingJapan MediaContentRatingJapanable
    // Media content rating settings for New Zealand
    mediaContentRatingNewZealand MediaContentRatingNewZealandable
    // Media content rating settings for United Kingdom
    mediaContentRatingUnitedKingdom MediaContentRatingUnitedKingdomable
    // Media content rating settings for United States
    mediaContentRatingUnitedStates MediaContentRatingUnitedStatesable
    // Indicates whether or not to block the user from using the Messages app on the supervised device.
    messagesBlocked *bool
    // List of managed apps and the network rules that applies to them. This collection can contain a maximum of 1000 elements.
    networkUsageRules []IosNetworkUsageRuleable
    // Indicates whether or not to allow notifications settings modification (iOS 9.3 and later).
    notificationsBlockSettingsModification *bool
    // Block modification of registered Touch ID fingerprints when in supervised mode.
    passcodeBlockFingerprintModification *bool
    // Indicates whether or not to block fingerprint unlock.
    passcodeBlockFingerprintUnlock *bool
    // Indicates whether or not to allow passcode modification on the supervised device (iOS 9.0 and later).
    passcodeBlockModification *bool
    // Indicates whether or not to block simple passcodes.
    passcodeBlockSimple *bool
    // Number of days before the passcode expires. Valid values 1 to 65535
    passcodeExpirationDays *int32
    // Number of character sets a passcode must contain. Valid values 0 to 4
    passcodeMinimumCharacterSetCount *int32
    // Minimum length of passcode. Valid values 4 to 14
    passcodeMinimumLength *int32
    // Minutes of inactivity before a passcode is required.
    passcodeMinutesOfInactivityBeforeLock *int32
    // Minutes of inactivity before the screen times out.
    passcodeMinutesOfInactivityBeforeScreenTimeout *int32
    // Number of previous passcodes to block. Valid values 1 to 24
    passcodePreviousPasscodeBlockCount *int32
    // Indicates whether or not to require a passcode.
    passcodeRequired *bool
    // Type of passcode that is required. Possible values are: deviceDefault, alphanumeric, numeric.
    passcodeRequiredType *RequiredPasswordType
    // Number of sign in failures allowed before wiping the device. Valid values 2 to 11
    passcodeSignInFailureCountBeforeWipe *int32
    // Indicates whether or not to block the user from using podcasts on the supervised device (iOS 8.0 and later).
    podcastsBlocked *bool
    // Indicates whether or not to block the user from using Auto fill in Safari. Requires a supervised device for iOS 13 and later.
    safariBlockAutofill *bool
    // Indicates whether or not to block the user from using Safari. Requires a supervised device for iOS 13 and later.
    safariBlocked *bool
    // Indicates whether or not to block JavaScript in Safari.
    safariBlockJavaScript *bool
    // Indicates whether or not to block popups in Safari.
    safariBlockPopups *bool
    // Cookie settings for Safari. Possible values are: browserDefault, blockAlways, allowCurrentWebSite, allowFromWebsitesVisited, allowAlways.
    safariCookieSettings *WebBrowserCookieSettings
    // URLs matching the patterns listed here will be considered managed.
    safariManagedDomains []string
    // Users can save passwords in Safari only from URLs matching the patterns listed here. Applies to devices in supervised mode (iOS 9.3 and later).
    safariPasswordAutoFillDomains []string
    // Indicates whether or not to require fraud warning in Safari.
    safariRequireFraudWarning *bool
    // Indicates whether or not to block the user from taking Screenshots.
    screenCaptureBlocked *bool
    // Indicates whether or not to block the user from using Siri.
    siriBlocked *bool
    // Indicates whether or not to block the user from using Siri when locked.
    siriBlockedWhenLocked *bool
    // Indicates whether or not to block Siri from querying user-generated content when used on a supervised device.
    siriBlockUserGeneratedContent *bool
    // Indicates whether or not to prevent Siri from dictating, or speaking profane language on supervised device.
    siriRequireProfanityFilter *bool
    // Indicates whether or not to block Spotlight search from returning internet results on supervised device.
    spotlightBlockInternetResults *bool
    // Indicates whether or not to block voice dialing.
    voiceDialingBlocked *bool
    // Indicates whether or not to allow wallpaper modification on supervised device (iOS 9.0 and later) .
    wallpaperBlockModification *bool
    // Indicates whether or not to force the device to use only Wi-Fi networks from configuration profiles when the device is in supervised mode. Available for devices running iOS and iPadOS versions 14.4 and earlier. Devices running 14.5+ should use the setting, 'WiFiConnectToAllowedNetworksOnlyForced.
    wiFiConnectOnlyToConfiguredNetworks *bool
}
// NewIosGeneralDeviceConfiguration instantiates a new IosGeneralDeviceConfiguration and sets the default values.
func NewIosGeneralDeviceConfiguration()(*IosGeneralDeviceConfiguration) {
    m := &IosGeneralDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    return m
}
// CreateIosGeneralDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosGeneralDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosGeneralDeviceConfiguration(), nil
}
// GetAccountBlockModification gets the accountBlockModification property value. Indicates whether or not to allow account modification when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetAccountBlockModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accountBlockModification
    }
}
// GetActivationLockAllowWhenSupervised gets the activationLockAllowWhenSupervised property value. Indicates whether or not to allow activation lock when the device is in the supervised mode.
func (m *IosGeneralDeviceConfiguration) GetActivationLockAllowWhenSupervised()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.activationLockAllowWhenSupervised
    }
}
// GetAirDropBlocked gets the airDropBlocked property value. Indicates whether or not to allow AirDrop when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetAirDropBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.airDropBlocked
    }
}
// GetAirDropForceUnmanagedDropTarget gets the airDropForceUnmanagedDropTarget property value. Indicates whether or not to cause AirDrop to be considered an unmanaged drop target (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAirDropForceUnmanagedDropTarget()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.airDropForceUnmanagedDropTarget
    }
}
// GetAirPlayForcePairingPasswordForOutgoingRequests gets the airPlayForcePairingPasswordForOutgoingRequests property value. Indicates whether or not to enforce all devices receiving AirPlay requests from this device to use a pairing password.
func (m *IosGeneralDeviceConfiguration) GetAirPlayForcePairingPasswordForOutgoingRequests()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.airPlayForcePairingPasswordForOutgoingRequests
    }
}
// GetAppleNewsBlocked gets the appleNewsBlocked property value. Indicates whether or not to block the user from using News when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAppleNewsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appleNewsBlocked
    }
}
// GetAppleWatchBlockPairing gets the appleWatchBlockPairing property value. Indicates whether or not to allow Apple Watch pairing when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAppleWatchBlockPairing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appleWatchBlockPairing
    }
}
// GetAppleWatchForceWristDetection gets the appleWatchForceWristDetection property value. Indicates whether or not to force a paired Apple Watch to use Wrist Detection (iOS 8.2 and later).
func (m *IosGeneralDeviceConfiguration) GetAppleWatchForceWristDetection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appleWatchForceWristDetection
    }
}
// GetAppsSingleAppModeList gets the appsSingleAppModeList property value. Gets or sets the list of iOS apps allowed to autonomously enter Single App Mode. Supervised only. iOS 7.0 and later. This collection can contain a maximum of 500 elements.
func (m *IosGeneralDeviceConfiguration) GetAppsSingleAppModeList()([]AppListItemable) {
    if m == nil {
        return nil
    } else {
        return m.appsSingleAppModeList
    }
}
// GetAppStoreBlockAutomaticDownloads gets the appStoreBlockAutomaticDownloads property value. Indicates whether or not to block the automatic downloading of apps purchased on other devices when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAppStoreBlockAutomaticDownloads()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appStoreBlockAutomaticDownloads
    }
}
// GetAppStoreBlocked gets the appStoreBlocked property value. Indicates whether or not to block the user from using the App Store. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetAppStoreBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appStoreBlocked
    }
}
// GetAppStoreBlockInAppPurchases gets the appStoreBlockInAppPurchases property value. Indicates whether or not to block the user from making in app purchases.
func (m *IosGeneralDeviceConfiguration) GetAppStoreBlockInAppPurchases()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appStoreBlockInAppPurchases
    }
}
// GetAppStoreBlockUIAppInstallation gets the appStoreBlockUIAppInstallation property value. Indicates whether or not to block the App Store app, not restricting installation through Host apps. Applies to supervised mode only (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAppStoreBlockUIAppInstallation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appStoreBlockUIAppInstallation
    }
}
// GetAppStoreRequirePassword gets the appStoreRequirePassword property value. Indicates whether or not to require a password when using the app store.
func (m *IosGeneralDeviceConfiguration) GetAppStoreRequirePassword()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appStoreRequirePassword
    }
}
// GetAppsVisibilityList gets the appsVisibilityList property value. List of apps in the visibility list (either visible/launchable apps list or hidden/unlaunchable apps list, controlled by AppsVisibilityListType) (iOS 9.3 and later). This collection can contain a maximum of 10000 elements.
func (m *IosGeneralDeviceConfiguration) GetAppsVisibilityList()([]AppListItemable) {
    if m == nil {
        return nil
    } else {
        return m.appsVisibilityList
    }
}
// GetAppsVisibilityListType gets the appsVisibilityListType property value. Type of list that is in the AppsVisibilityList. Possible values are: none, appsInListCompliant, appsNotInListCompliant.
func (m *IosGeneralDeviceConfiguration) GetAppsVisibilityListType()(*AppListType) {
    if m == nil {
        return nil
    } else {
        return m.appsVisibilityListType
    }
}
// GetBluetoothBlockModification gets the bluetoothBlockModification property value. Indicates whether or not to allow modification of Bluetooth settings when the device is in supervised mode (iOS 10.0 and later).
func (m *IosGeneralDeviceConfiguration) GetBluetoothBlockModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.bluetoothBlockModification
    }
}
// GetCameraBlocked gets the cameraBlocked property value. Indicates whether or not to block the user from accessing the camera of the device. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetCameraBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.cameraBlocked
    }
}
// GetCellularBlockDataRoaming gets the cellularBlockDataRoaming property value. Indicates whether or not to block data roaming.
func (m *IosGeneralDeviceConfiguration) GetCellularBlockDataRoaming()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.cellularBlockDataRoaming
    }
}
// GetCellularBlockGlobalBackgroundFetchWhileRoaming gets the cellularBlockGlobalBackgroundFetchWhileRoaming property value. Indicates whether or not to block global background fetch while roaming.
func (m *IosGeneralDeviceConfiguration) GetCellularBlockGlobalBackgroundFetchWhileRoaming()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.cellularBlockGlobalBackgroundFetchWhileRoaming
    }
}
// GetCellularBlockPerAppDataModification gets the cellularBlockPerAppDataModification property value. Indicates whether or not to allow changes to cellular app data usage settings when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetCellularBlockPerAppDataModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.cellularBlockPerAppDataModification
    }
}
// GetCellularBlockPersonalHotspot gets the cellularBlockPersonalHotspot property value. Indicates whether or not to block Personal Hotspot.
func (m *IosGeneralDeviceConfiguration) GetCellularBlockPersonalHotspot()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.cellularBlockPersonalHotspot
    }
}
// GetCellularBlockVoiceRoaming gets the cellularBlockVoiceRoaming property value. Indicates whether or not to block voice roaming.
func (m *IosGeneralDeviceConfiguration) GetCellularBlockVoiceRoaming()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.cellularBlockVoiceRoaming
    }
}
// GetCertificatesBlockUntrustedTlsCertificates gets the certificatesBlockUntrustedTlsCertificates property value. Indicates whether or not to block untrusted TLS certificates.
func (m *IosGeneralDeviceConfiguration) GetCertificatesBlockUntrustedTlsCertificates()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.certificatesBlockUntrustedTlsCertificates
    }
}
// GetClassroomAppBlockRemoteScreenObservation gets the classroomAppBlockRemoteScreenObservation property value. Indicates whether or not to allow remote screen observation by Classroom app when the device is in supervised mode (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) GetClassroomAppBlockRemoteScreenObservation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.classroomAppBlockRemoteScreenObservation
    }
}
// GetClassroomAppForceUnpromptedScreenObservation gets the classroomAppForceUnpromptedScreenObservation property value. Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to view a student's screen without prompting when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetClassroomAppForceUnpromptedScreenObservation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.classroomAppForceUnpromptedScreenObservation
    }
}
// GetCompliantAppListType gets the compliantAppListType property value. List that is in the AppComplianceList. Possible values are: none, appsInListCompliant, appsNotInListCompliant.
func (m *IosGeneralDeviceConfiguration) GetCompliantAppListType()(*AppListType) {
    if m == nil {
        return nil
    } else {
        return m.compliantAppListType
    }
}
// GetCompliantAppsList gets the compliantAppsList property value. List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
func (m *IosGeneralDeviceConfiguration) GetCompliantAppsList()([]AppListItemable) {
    if m == nil {
        return nil
    } else {
        return m.compliantAppsList
    }
}
// GetConfigurationProfileBlockChanges gets the configurationProfileBlockChanges property value. Indicates whether or not to block the user from installing configuration profiles and certificates interactively when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetConfigurationProfileBlockChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.configurationProfileBlockChanges
    }
}
// GetDefinitionLookupBlocked gets the definitionLookupBlocked property value. Indicates whether or not to block definition lookup when the device is in supervised mode (iOS 8.1.3 and later ).
func (m *IosGeneralDeviceConfiguration) GetDefinitionLookupBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.definitionLookupBlocked
    }
}
// GetDeviceBlockEnableRestrictions gets the deviceBlockEnableRestrictions property value. Indicates whether or not to allow the user to enables restrictions in the device settings when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetDeviceBlockEnableRestrictions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceBlockEnableRestrictions
    }
}
// GetDeviceBlockEraseContentAndSettings gets the deviceBlockEraseContentAndSettings property value. Indicates whether or not to allow the use of the 'Erase all content and settings' option on the device when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetDeviceBlockEraseContentAndSettings()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceBlockEraseContentAndSettings
    }
}
// GetDeviceBlockNameModification gets the deviceBlockNameModification property value. Indicates whether or not to allow device name modification when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetDeviceBlockNameModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceBlockNameModification
    }
}
// GetDiagnosticDataBlockSubmission gets the diagnosticDataBlockSubmission property value. Indicates whether or not to block diagnostic data submission.
func (m *IosGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmission()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.diagnosticDataBlockSubmission
    }
}
// GetDiagnosticDataBlockSubmissionModification gets the diagnosticDataBlockSubmissionModification property value. Indicates whether or not to allow diagnostics submission settings modification when the device is in supervised mode (iOS 9.3.2 and later).
func (m *IosGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmissionModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.diagnosticDataBlockSubmissionModification
    }
}
// GetDocumentsBlockManagedDocumentsInUnmanagedApps gets the documentsBlockManagedDocumentsInUnmanagedApps property value. Indicates whether or not to block the user from viewing managed documents in unmanaged apps.
func (m *IosGeneralDeviceConfiguration) GetDocumentsBlockManagedDocumentsInUnmanagedApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.documentsBlockManagedDocumentsInUnmanagedApps
    }
}
// GetDocumentsBlockUnmanagedDocumentsInManagedApps gets the documentsBlockUnmanagedDocumentsInManagedApps property value. Indicates whether or not to block the user from viewing unmanaged documents in managed apps.
func (m *IosGeneralDeviceConfiguration) GetDocumentsBlockUnmanagedDocumentsInManagedApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.documentsBlockUnmanagedDocumentsInManagedApps
    }
}
// GetEmailInDomainSuffixes gets the emailInDomainSuffixes property value. An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
func (m *IosGeneralDeviceConfiguration) GetEmailInDomainSuffixes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.emailInDomainSuffixes
    }
}
// GetEnterpriseAppBlockTrust gets the enterpriseAppBlockTrust property value. Indicates whether or not to block the user from trusting an enterprise app.
func (m *IosGeneralDeviceConfiguration) GetEnterpriseAppBlockTrust()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseAppBlockTrust
    }
}
// GetEnterpriseAppBlockTrustModification gets the enterpriseAppBlockTrustModification property value. [Deprecated] Configuring this setting and setting the value to 'true' has no effect on the device.
func (m *IosGeneralDeviceConfiguration) GetEnterpriseAppBlockTrustModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseAppBlockTrustModification
    }
}
// GetFaceTimeBlocked gets the faceTimeBlocked property value. Indicates whether or not to block the user from using FaceTime. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetFaceTimeBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.faceTimeBlocked
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosGeneralDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["accountBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountBlockModification(val)
        }
        return nil
    }
    res["activationLockAllowWhenSupervised"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationLockAllowWhenSupervised(val)
        }
        return nil
    }
    res["airDropBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAirDropBlocked(val)
        }
        return nil
    }
    res["airDropForceUnmanagedDropTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAirDropForceUnmanagedDropTarget(val)
        }
        return nil
    }
    res["airPlayForcePairingPasswordForOutgoingRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAirPlayForcePairingPasswordForOutgoingRequests(val)
        }
        return nil
    }
    res["appleNewsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleNewsBlocked(val)
        }
        return nil
    }
    res["appleWatchBlockPairing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleWatchBlockPairing(val)
        }
        return nil
    }
    res["appleWatchForceWristDetection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleWatchForceWristDetection(val)
        }
        return nil
    }
    res["appsSingleAppModeList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetAppsSingleAppModeList(res)
        }
        return nil
    }
    res["appStoreBlockAutomaticDownloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreBlockAutomaticDownloads(val)
        }
        return nil
    }
    res["appStoreBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreBlocked(val)
        }
        return nil
    }
    res["appStoreBlockInAppPurchases"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreBlockInAppPurchases(val)
        }
        return nil
    }
    res["appStoreBlockUIAppInstallation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreBlockUIAppInstallation(val)
        }
        return nil
    }
    res["appStoreRequirePassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreRequirePassword(val)
        }
        return nil
    }
    res["appsVisibilityList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetAppsVisibilityList(res)
        }
        return nil
    }
    res["appsVisibilityListType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppListType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsVisibilityListType(val.(*AppListType))
        }
        return nil
    }
    res["bluetoothBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetoothBlockModification(val)
        }
        return nil
    }
    res["cameraBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCameraBlocked(val)
        }
        return nil
    }
    res["cellularBlockDataRoaming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockDataRoaming(val)
        }
        return nil
    }
    res["cellularBlockGlobalBackgroundFetchWhileRoaming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockGlobalBackgroundFetchWhileRoaming(val)
        }
        return nil
    }
    res["cellularBlockPerAppDataModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockPerAppDataModification(val)
        }
        return nil
    }
    res["cellularBlockPersonalHotspot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockPersonalHotspot(val)
        }
        return nil
    }
    res["cellularBlockVoiceRoaming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockVoiceRoaming(val)
        }
        return nil
    }
    res["certificatesBlockUntrustedTlsCertificates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificatesBlockUntrustedTlsCertificates(val)
        }
        return nil
    }
    res["classroomAppBlockRemoteScreenObservation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassroomAppBlockRemoteScreenObservation(val)
        }
        return nil
    }
    res["classroomAppForceUnpromptedScreenObservation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassroomAppForceUnpromptedScreenObservation(val)
        }
        return nil
    }
    res["compliantAppListType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppListType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantAppListType(val.(*AppListType))
        }
        return nil
    }
    res["compliantAppsList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetCompliantAppsList(res)
        }
        return nil
    }
    res["configurationProfileBlockChanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationProfileBlockChanges(val)
        }
        return nil
    }
    res["definitionLookupBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionLookupBlocked(val)
        }
        return nil
    }
    res["deviceBlockEnableRestrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceBlockEnableRestrictions(val)
        }
        return nil
    }
    res["deviceBlockEraseContentAndSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceBlockEraseContentAndSettings(val)
        }
        return nil
    }
    res["deviceBlockNameModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceBlockNameModification(val)
        }
        return nil
    }
    res["diagnosticDataBlockSubmission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiagnosticDataBlockSubmission(val)
        }
        return nil
    }
    res["diagnosticDataBlockSubmissionModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiagnosticDataBlockSubmissionModification(val)
        }
        return nil
    }
    res["documentsBlockManagedDocumentsInUnmanagedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentsBlockManagedDocumentsInUnmanagedApps(val)
        }
        return nil
    }
    res["documentsBlockUnmanagedDocumentsInManagedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentsBlockUnmanagedDocumentsInManagedApps(val)
        }
        return nil
    }
    res["emailInDomainSuffixes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEmailInDomainSuffixes(res)
        }
        return nil
    }
    res["enterpriseAppBlockTrust"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseAppBlockTrust(val)
        }
        return nil
    }
    res["enterpriseAppBlockTrustModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseAppBlockTrustModification(val)
        }
        return nil
    }
    res["faceTimeBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaceTimeBlocked(val)
        }
        return nil
    }
    res["findMyFriendsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFindMyFriendsBlocked(val)
        }
        return nil
    }
    res["gameCenterBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGameCenterBlocked(val)
        }
        return nil
    }
    res["gamingBlockGameCenterFriends"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGamingBlockGameCenterFriends(val)
        }
        return nil
    }
    res["gamingBlockMultiplayer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGamingBlockMultiplayer(val)
        }
        return nil
    }
    res["hostPairingBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostPairingBlocked(val)
        }
        return nil
    }
    res["iBooksStoreBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIBooksStoreBlocked(val)
        }
        return nil
    }
    res["iBooksStoreBlockErotica"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIBooksStoreBlockErotica(val)
        }
        return nil
    }
    res["iCloudBlockActivityContinuation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockActivityContinuation(val)
        }
        return nil
    }
    res["iCloudBlockBackup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockBackup(val)
        }
        return nil
    }
    res["iCloudBlockDocumentSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockDocumentSync(val)
        }
        return nil
    }
    res["iCloudBlockManagedAppsSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockManagedAppsSync(val)
        }
        return nil
    }
    res["iCloudBlockPhotoLibrary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockPhotoLibrary(val)
        }
        return nil
    }
    res["iCloudBlockPhotoStreamSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockPhotoStreamSync(val)
        }
        return nil
    }
    res["iCloudBlockSharedPhotoStream"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockSharedPhotoStream(val)
        }
        return nil
    }
    res["iCloudRequireEncryptedBackup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudRequireEncryptedBackup(val)
        }
        return nil
    }
    res["iTunesBlockExplicitContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetITunesBlockExplicitContent(val)
        }
        return nil
    }
    res["iTunesBlockMusicService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetITunesBlockMusicService(val)
        }
        return nil
    }
    res["iTunesBlockRadio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetITunesBlockRadio(val)
        }
        return nil
    }
    res["keyboardBlockAutoCorrect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardBlockAutoCorrect(val)
        }
        return nil
    }
    res["keyboardBlockDictation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardBlockDictation(val)
        }
        return nil
    }
    res["keyboardBlockPredictive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardBlockPredictive(val)
        }
        return nil
    }
    res["keyboardBlockShortcuts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardBlockShortcuts(val)
        }
        return nil
    }
    res["keyboardBlockSpellCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardBlockSpellCheck(val)
        }
        return nil
    }
    res["kioskModeAllowAssistiveSpeak"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowAssistiveSpeak(val)
        }
        return nil
    }
    res["kioskModeAllowAssistiveTouchSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowAssistiveTouchSettings(val)
        }
        return nil
    }
    res["kioskModeAllowAutoLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowAutoLock(val)
        }
        return nil
    }
    res["kioskModeAllowColorInversionSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowColorInversionSettings(val)
        }
        return nil
    }
    res["kioskModeAllowRingerSwitch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowRingerSwitch(val)
        }
        return nil
    }
    res["kioskModeAllowScreenRotation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowScreenRotation(val)
        }
        return nil
    }
    res["kioskModeAllowSleepButton"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowSleepButton(val)
        }
        return nil
    }
    res["kioskModeAllowTouchscreen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowTouchscreen(val)
        }
        return nil
    }
    res["kioskModeAllowVoiceOverSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowVoiceOverSettings(val)
        }
        return nil
    }
    res["kioskModeAllowVolumeButtons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowVolumeButtons(val)
        }
        return nil
    }
    res["kioskModeAllowZoomSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowZoomSettings(val)
        }
        return nil
    }
    res["kioskModeAppStoreUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAppStoreUrl(val)
        }
        return nil
    }
    res["kioskModeBuiltInAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeBuiltInAppId(val)
        }
        return nil
    }
    res["kioskModeManagedAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedAppId(val)
        }
        return nil
    }
    res["kioskModeRequireAssistiveTouch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeRequireAssistiveTouch(val)
        }
        return nil
    }
    res["kioskModeRequireColorInversion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeRequireColorInversion(val)
        }
        return nil
    }
    res["kioskModeRequireMonoAudio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeRequireMonoAudio(val)
        }
        return nil
    }
    res["kioskModeRequireVoiceOver"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeRequireVoiceOver(val)
        }
        return nil
    }
    res["kioskModeRequireZoom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeRequireZoom(val)
        }
        return nil
    }
    res["lockScreenBlockControlCenter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenBlockControlCenter(val)
        }
        return nil
    }
    res["lockScreenBlockNotificationView"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenBlockNotificationView(val)
        }
        return nil
    }
    res["lockScreenBlockPassbook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenBlockPassbook(val)
        }
        return nil
    }
    res["lockScreenBlockTodayView"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenBlockTodayView(val)
        }
        return nil
    }
    res["mediaContentRatingApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRatingAppsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingApps(val.(*RatingAppsType))
        }
        return nil
    }
    res["mediaContentRatingAustralia"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingAustraliaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingAustralia(val.(MediaContentRatingAustraliaable))
        }
        return nil
    }
    res["mediaContentRatingCanada"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingCanadaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingCanada(val.(MediaContentRatingCanadaable))
        }
        return nil
    }
    res["mediaContentRatingFrance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingFranceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingFrance(val.(MediaContentRatingFranceable))
        }
        return nil
    }
    res["mediaContentRatingGermany"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingGermanyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingGermany(val.(MediaContentRatingGermanyable))
        }
        return nil
    }
    res["mediaContentRatingIreland"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingIrelandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingIreland(val.(MediaContentRatingIrelandable))
        }
        return nil
    }
    res["mediaContentRatingJapan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingJapanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingJapan(val.(MediaContentRatingJapanable))
        }
        return nil
    }
    res["mediaContentRatingNewZealand"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingNewZealandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingNewZealand(val.(MediaContentRatingNewZealandable))
        }
        return nil
    }
    res["mediaContentRatingUnitedKingdom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingUnitedKingdomFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingUnitedKingdom(val.(MediaContentRatingUnitedKingdomable))
        }
        return nil
    }
    res["mediaContentRatingUnitedStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingUnitedStatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingUnitedStates(val.(MediaContentRatingUnitedStatesable))
        }
        return nil
    }
    res["messagesBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessagesBlocked(val)
        }
        return nil
    }
    res["networkUsageRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosNetworkUsageRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosNetworkUsageRuleable, len(val))
            for i, v := range val {
                res[i] = v.(IosNetworkUsageRuleable)
            }
            m.SetNetworkUsageRules(res)
        }
        return nil
    }
    res["notificationsBlockSettingsModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationsBlockSettingsModification(val)
        }
        return nil
    }
    res["passcodeBlockFingerprintModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeBlockFingerprintModification(val)
        }
        return nil
    }
    res["passcodeBlockFingerprintUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeBlockFingerprintUnlock(val)
        }
        return nil
    }
    res["passcodeBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeBlockModification(val)
        }
        return nil
    }
    res["passcodeBlockSimple"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeBlockSimple(val)
        }
        return nil
    }
    res["passcodeExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeExpirationDays(val)
        }
        return nil
    }
    res["passcodeMinimumCharacterSetCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinimumCharacterSetCount(val)
        }
        return nil
    }
    res["passcodeMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinimumLength(val)
        }
        return nil
    }
    res["passcodeMinutesOfInactivityBeforeLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinutesOfInactivityBeforeLock(val)
        }
        return nil
    }
    res["passcodeMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinutesOfInactivityBeforeScreenTimeout(val)
        }
        return nil
    }
    res["passcodePreviousPasscodeBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodePreviousPasscodeBlockCount(val)
        }
        return nil
    }
    res["passcodeRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeRequired(val)
        }
        return nil
    }
    res["passcodeRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeRequiredType(val.(*RequiredPasswordType))
        }
        return nil
    }
    res["passcodeSignInFailureCountBeforeWipe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeSignInFailureCountBeforeWipe(val)
        }
        return nil
    }
    res["podcastsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPodcastsBlocked(val)
        }
        return nil
    }
    res["safariBlockAutofill"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafariBlockAutofill(val)
        }
        return nil
    }
    res["safariBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafariBlocked(val)
        }
        return nil
    }
    res["safariBlockJavaScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafariBlockJavaScript(val)
        }
        return nil
    }
    res["safariBlockPopups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafariBlockPopups(val)
        }
        return nil
    }
    res["safariCookieSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWebBrowserCookieSettings)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafariCookieSettings(val.(*WebBrowserCookieSettings))
        }
        return nil
    }
    res["safariManagedDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSafariManagedDomains(res)
        }
        return nil
    }
    res["safariPasswordAutoFillDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSafariPasswordAutoFillDomains(res)
        }
        return nil
    }
    res["safariRequireFraudWarning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafariRequireFraudWarning(val)
        }
        return nil
    }
    res["screenCaptureBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreenCaptureBlocked(val)
        }
        return nil
    }
    res["siriBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiriBlocked(val)
        }
        return nil
    }
    res["siriBlockedWhenLocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiriBlockedWhenLocked(val)
        }
        return nil
    }
    res["siriBlockUserGeneratedContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiriBlockUserGeneratedContent(val)
        }
        return nil
    }
    res["siriRequireProfanityFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiriRequireProfanityFilter(val)
        }
        return nil
    }
    res["spotlightBlockInternetResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpotlightBlockInternetResults(val)
        }
        return nil
    }
    res["voiceDialingBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVoiceDialingBlocked(val)
        }
        return nil
    }
    res["wallpaperBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWallpaperBlockModification(val)
        }
        return nil
    }
    res["wiFiConnectOnlyToConfiguredNetworks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiConnectOnlyToConfiguredNetworks(val)
        }
        return nil
    }
    return res
}
// GetFindMyFriendsBlocked gets the findMyFriendsBlocked property value. Indicates whether or not to block changes to Find My Friends when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetFindMyFriendsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.findMyFriendsBlocked
    }
}
// GetGameCenterBlocked gets the gameCenterBlocked property value. Indicates whether or not to block the user from using Game Center when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetGameCenterBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.gameCenterBlocked
    }
}
// GetGamingBlockGameCenterFriends gets the gamingBlockGameCenterFriends property value. Indicates whether or not to block the user from having friends in Game Center. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetGamingBlockGameCenterFriends()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.gamingBlockGameCenterFriends
    }
}
// GetGamingBlockMultiplayer gets the gamingBlockMultiplayer property value. Indicates whether or not to block the user from using multiplayer gaming. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetGamingBlockMultiplayer()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.gamingBlockMultiplayer
    }
}
// GetHostPairingBlocked gets the hostPairingBlocked property value. indicates whether or not to allow host pairing to control the devices an iOS device can pair with when the iOS device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetHostPairingBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hostPairingBlocked
    }
}
// GetIBooksStoreBlocked gets the iBooksStoreBlocked property value. Indicates whether or not to block the user from using the iBooks Store when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetIBooksStoreBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iBooksStoreBlocked
    }
}
// GetIBooksStoreBlockErotica gets the iBooksStoreBlockErotica property value. Indicates whether or not to block the user from downloading media from the iBookstore that has been tagged as erotica.
func (m *IosGeneralDeviceConfiguration) GetIBooksStoreBlockErotica()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iBooksStoreBlockErotica
    }
}
// GetICloudBlockActivityContinuation gets the iCloudBlockActivityContinuation property value. Indicates whether or not to block the user from continuing work they started on iOS device to another iOS or macOS device.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockActivityContinuation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockActivityContinuation
    }
}
// GetICloudBlockBackup gets the iCloudBlockBackup property value. Indicates whether or not to block iCloud backup. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockBackup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockBackup
    }
}
// GetICloudBlockDocumentSync gets the iCloudBlockDocumentSync property value. Indicates whether or not to block iCloud document sync. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockDocumentSync()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockDocumentSync
    }
}
// GetICloudBlockManagedAppsSync gets the iCloudBlockManagedAppsSync property value. Indicates whether or not to block Managed Apps Cloud Sync.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockManagedAppsSync()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockManagedAppsSync
    }
}
// GetICloudBlockPhotoLibrary gets the iCloudBlockPhotoLibrary property value. Indicates whether or not to block iCloud Photo Library.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockPhotoLibrary()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockPhotoLibrary
    }
}
// GetICloudBlockPhotoStreamSync gets the iCloudBlockPhotoStreamSync property value. Indicates whether or not to block iCloud Photo Stream Sync.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockPhotoStreamSync()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockPhotoStreamSync
    }
}
// GetICloudBlockSharedPhotoStream gets the iCloudBlockSharedPhotoStream property value. Indicates whether or not to block Shared Photo Stream.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockSharedPhotoStream()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockSharedPhotoStream
    }
}
// GetICloudRequireEncryptedBackup gets the iCloudRequireEncryptedBackup property value. Indicates whether or not to require backups to iCloud be encrypted.
func (m *IosGeneralDeviceConfiguration) GetICloudRequireEncryptedBackup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudRequireEncryptedBackup
    }
}
// GetITunesBlockExplicitContent gets the iTunesBlockExplicitContent property value. Indicates whether or not to block the user from accessing explicit content in iTunes and the App Store. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetITunesBlockExplicitContent()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iTunesBlockExplicitContent
    }
}
// GetITunesBlockMusicService gets the iTunesBlockMusicService property value. Indicates whether or not to block Music service and revert Music app to classic mode when the device is in supervised mode (iOS 9.3 and later and macOS 10.12 and later).
func (m *IosGeneralDeviceConfiguration) GetITunesBlockMusicService()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iTunesBlockMusicService
    }
}
// GetITunesBlockRadio gets the iTunesBlockRadio property value. Indicates whether or not to block the user from using iTunes Radio when the device is in supervised mode (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) GetITunesBlockRadio()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iTunesBlockRadio
    }
}
// GetKeyboardBlockAutoCorrect gets the keyboardBlockAutoCorrect property value. Indicates whether or not to block keyboard auto-correction when the device is in supervised mode (iOS 8.1.3 and later).
func (m *IosGeneralDeviceConfiguration) GetKeyboardBlockAutoCorrect()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keyboardBlockAutoCorrect
    }
}
// GetKeyboardBlockDictation gets the keyboardBlockDictation property value. Indicates whether or not to block the user from using dictation input when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetKeyboardBlockDictation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keyboardBlockDictation
    }
}
// GetKeyboardBlockPredictive gets the keyboardBlockPredictive property value. Indicates whether or not to block predictive keyboards when device is in supervised mode (iOS 8.1.3 and later).
func (m *IosGeneralDeviceConfiguration) GetKeyboardBlockPredictive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keyboardBlockPredictive
    }
}
// GetKeyboardBlockShortcuts gets the keyboardBlockShortcuts property value. Indicates whether or not to block keyboard shortcuts when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetKeyboardBlockShortcuts()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keyboardBlockShortcuts
    }
}
// GetKeyboardBlockSpellCheck gets the keyboardBlockSpellCheck property value. Indicates whether or not to block keyboard spell-checking when the device is in supervised mode (iOS 8.1.3 and later).
func (m *IosGeneralDeviceConfiguration) GetKeyboardBlockSpellCheck()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keyboardBlockSpellCheck
    }
}
// GetKioskModeAllowAssistiveSpeak gets the kioskModeAllowAssistiveSpeak property value. Indicates whether or not to allow assistive speak while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowAssistiveSpeak()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeAllowAssistiveSpeak
    }
}
// GetKioskModeAllowAssistiveTouchSettings gets the kioskModeAllowAssistiveTouchSettings property value. Indicates whether or not to allow access to the Assistive Touch Settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowAssistiveTouchSettings()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeAllowAssistiveTouchSettings
    }
}
// GetKioskModeAllowAutoLock gets the kioskModeAllowAutoLock property value. Indicates whether or not to allow device auto lock while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockAutoLock instead.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowAutoLock()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeAllowAutoLock
    }
}
// GetKioskModeAllowColorInversionSettings gets the kioskModeAllowColorInversionSettings property value. Indicates whether or not to allow access to the Color Inversion Settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowColorInversionSettings()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeAllowColorInversionSettings
    }
}
// GetKioskModeAllowRingerSwitch gets the kioskModeAllowRingerSwitch property value. Indicates whether or not to allow use of the ringer switch while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockRingerSwitch instead.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowRingerSwitch()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeAllowRingerSwitch
    }
}
// GetKioskModeAllowScreenRotation gets the kioskModeAllowScreenRotation property value. Indicates whether or not to allow screen rotation while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockScreenRotation instead.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowScreenRotation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeAllowScreenRotation
    }
}
// GetKioskModeAllowSleepButton gets the kioskModeAllowSleepButton property value. Indicates whether or not to allow use of the sleep button while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockSleepButton instead.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowSleepButton()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeAllowSleepButton
    }
}
// GetKioskModeAllowTouchscreen gets the kioskModeAllowTouchscreen property value. Indicates whether or not to allow use of the touchscreen while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockTouchscreen instead.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowTouchscreen()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeAllowTouchscreen
    }
}
// GetKioskModeAllowVoiceOverSettings gets the kioskModeAllowVoiceOverSettings property value. Indicates whether or not to allow access to the voice over settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowVoiceOverSettings()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeAllowVoiceOverSettings
    }
}
// GetKioskModeAllowVolumeButtons gets the kioskModeAllowVolumeButtons property value. Indicates whether or not to allow use of the volume buttons while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockVolumeButtons instead.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowVolumeButtons()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeAllowVolumeButtons
    }
}
// GetKioskModeAllowZoomSettings gets the kioskModeAllowZoomSettings property value. Indicates whether or not to allow access to the zoom settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowZoomSettings()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeAllowZoomSettings
    }
}
// GetKioskModeAppStoreUrl gets the kioskModeAppStoreUrl property value. URL in the app store to the app to use for kiosk mode. Use if KioskModeManagedAppId is not known.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAppStoreUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeAppStoreUrl
    }
}
// GetKioskModeBuiltInAppId gets the kioskModeBuiltInAppId property value. ID for built-in apps to use for kiosk mode. Used when KioskModeManagedAppId and KioskModeAppStoreUrl are not set.
func (m *IosGeneralDeviceConfiguration) GetKioskModeBuiltInAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeBuiltInAppId
    }
}
// GetKioskModeManagedAppId gets the kioskModeManagedAppId property value. Managed app id of the app to use for kiosk mode. If KioskModeManagedAppId is specified then KioskModeAppStoreUrl will be ignored.
func (m *IosGeneralDeviceConfiguration) GetKioskModeManagedAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeManagedAppId
    }
}
// GetKioskModeRequireAssistiveTouch gets the kioskModeRequireAssistiveTouch property value. Indicates whether or not to require assistive touch while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeRequireAssistiveTouch()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeRequireAssistiveTouch
    }
}
// GetKioskModeRequireColorInversion gets the kioskModeRequireColorInversion property value. Indicates whether or not to require color inversion while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeRequireColorInversion()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeRequireColorInversion
    }
}
// GetKioskModeRequireMonoAudio gets the kioskModeRequireMonoAudio property value. Indicates whether or not to require mono audio while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeRequireMonoAudio()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeRequireMonoAudio
    }
}
// GetKioskModeRequireVoiceOver gets the kioskModeRequireVoiceOver property value. Indicates whether or not to require voice over while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeRequireVoiceOver()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeRequireVoiceOver
    }
}
// GetKioskModeRequireZoom gets the kioskModeRequireZoom property value. Indicates whether or not to require zoom while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeRequireZoom()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.kioskModeRequireZoom
    }
}
// GetLockScreenBlockControlCenter gets the lockScreenBlockControlCenter property value. Indicates whether or not to block the user from using control center on the lock screen.
func (m *IosGeneralDeviceConfiguration) GetLockScreenBlockControlCenter()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.lockScreenBlockControlCenter
    }
}
// GetLockScreenBlockNotificationView gets the lockScreenBlockNotificationView property value. Indicates whether or not to block the user from using the notification view on the lock screen.
func (m *IosGeneralDeviceConfiguration) GetLockScreenBlockNotificationView()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.lockScreenBlockNotificationView
    }
}
// GetLockScreenBlockPassbook gets the lockScreenBlockPassbook property value. Indicates whether or not to block the user from using passbook when the device is locked.
func (m *IosGeneralDeviceConfiguration) GetLockScreenBlockPassbook()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.lockScreenBlockPassbook
    }
}
// GetLockScreenBlockTodayView gets the lockScreenBlockTodayView property value. Indicates whether or not to block the user from using the Today View on the lock screen.
func (m *IosGeneralDeviceConfiguration) GetLockScreenBlockTodayView()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.lockScreenBlockTodayView
    }
}
// GetMediaContentRatingApps gets the mediaContentRatingApps property value. Media content rating settings for Apps. Possible values are: allAllowed, allBlocked, agesAbove4, agesAbove9, agesAbove12, agesAbove17.
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingApps()(*RatingAppsType) {
    if m == nil {
        return nil
    } else {
        return m.mediaContentRatingApps
    }
}
// GetMediaContentRatingAustralia gets the mediaContentRatingAustralia property value. Media content rating settings for Australia
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingAustralia()(MediaContentRatingAustraliaable) {
    if m == nil {
        return nil
    } else {
        return m.mediaContentRatingAustralia
    }
}
// GetMediaContentRatingCanada gets the mediaContentRatingCanada property value. Media content rating settings for Canada
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingCanada()(MediaContentRatingCanadaable) {
    if m == nil {
        return nil
    } else {
        return m.mediaContentRatingCanada
    }
}
// GetMediaContentRatingFrance gets the mediaContentRatingFrance property value. Media content rating settings for France
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingFrance()(MediaContentRatingFranceable) {
    if m == nil {
        return nil
    } else {
        return m.mediaContentRatingFrance
    }
}
// GetMediaContentRatingGermany gets the mediaContentRatingGermany property value. Media content rating settings for Germany
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingGermany()(MediaContentRatingGermanyable) {
    if m == nil {
        return nil
    } else {
        return m.mediaContentRatingGermany
    }
}
// GetMediaContentRatingIreland gets the mediaContentRatingIreland property value. Media content rating settings for Ireland
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingIreland()(MediaContentRatingIrelandable) {
    if m == nil {
        return nil
    } else {
        return m.mediaContentRatingIreland
    }
}
// GetMediaContentRatingJapan gets the mediaContentRatingJapan property value. Media content rating settings for Japan
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingJapan()(MediaContentRatingJapanable) {
    if m == nil {
        return nil
    } else {
        return m.mediaContentRatingJapan
    }
}
// GetMediaContentRatingNewZealand gets the mediaContentRatingNewZealand property value. Media content rating settings for New Zealand
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingNewZealand()(MediaContentRatingNewZealandable) {
    if m == nil {
        return nil
    } else {
        return m.mediaContentRatingNewZealand
    }
}
// GetMediaContentRatingUnitedKingdom gets the mediaContentRatingUnitedKingdom property value. Media content rating settings for United Kingdom
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingUnitedKingdom()(MediaContentRatingUnitedKingdomable) {
    if m == nil {
        return nil
    } else {
        return m.mediaContentRatingUnitedKingdom
    }
}
// GetMediaContentRatingUnitedStates gets the mediaContentRatingUnitedStates property value. Media content rating settings for United States
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingUnitedStates()(MediaContentRatingUnitedStatesable) {
    if m == nil {
        return nil
    } else {
        return m.mediaContentRatingUnitedStates
    }
}
// GetMessagesBlocked gets the messagesBlocked property value. Indicates whether or not to block the user from using the Messages app on the supervised device.
func (m *IosGeneralDeviceConfiguration) GetMessagesBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.messagesBlocked
    }
}
// GetNetworkUsageRules gets the networkUsageRules property value. List of managed apps and the network rules that applies to them. This collection can contain a maximum of 1000 elements.
func (m *IosGeneralDeviceConfiguration) GetNetworkUsageRules()([]IosNetworkUsageRuleable) {
    if m == nil {
        return nil
    } else {
        return m.networkUsageRules
    }
}
// GetNotificationsBlockSettingsModification gets the notificationsBlockSettingsModification property value. Indicates whether or not to allow notifications settings modification (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) GetNotificationsBlockSettingsModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notificationsBlockSettingsModification
    }
}
// GetPasscodeBlockFingerprintModification gets the passcodeBlockFingerprintModification property value. Block modification of registered Touch ID fingerprints when in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetPasscodeBlockFingerprintModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passcodeBlockFingerprintModification
    }
}
// GetPasscodeBlockFingerprintUnlock gets the passcodeBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *IosGeneralDeviceConfiguration) GetPasscodeBlockFingerprintUnlock()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passcodeBlockFingerprintUnlock
    }
}
// GetPasscodeBlockModification gets the passcodeBlockModification property value. Indicates whether or not to allow passcode modification on the supervised device (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetPasscodeBlockModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passcodeBlockModification
    }
}
// GetPasscodeBlockSimple gets the passcodeBlockSimple property value. Indicates whether or not to block simple passcodes.
func (m *IosGeneralDeviceConfiguration) GetPasscodeBlockSimple()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passcodeBlockSimple
    }
}
// GetPasscodeExpirationDays gets the passcodeExpirationDays property value. Number of days before the passcode expires. Valid values 1 to 65535
func (m *IosGeneralDeviceConfiguration) GetPasscodeExpirationDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodeExpirationDays
    }
}
// GetPasscodeMinimumCharacterSetCount gets the passcodeMinimumCharacterSetCount property value. Number of character sets a passcode must contain. Valid values 0 to 4
func (m *IosGeneralDeviceConfiguration) GetPasscodeMinimumCharacterSetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodeMinimumCharacterSetCount
    }
}
// GetPasscodeMinimumLength gets the passcodeMinimumLength property value. Minimum length of passcode. Valid values 4 to 14
func (m *IosGeneralDeviceConfiguration) GetPasscodeMinimumLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodeMinimumLength
    }
}
// GetPasscodeMinutesOfInactivityBeforeLock gets the passcodeMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a passcode is required.
func (m *IosGeneralDeviceConfiguration) GetPasscodeMinutesOfInactivityBeforeLock()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodeMinutesOfInactivityBeforeLock
    }
}
// GetPasscodeMinutesOfInactivityBeforeScreenTimeout gets the passcodeMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *IosGeneralDeviceConfiguration) GetPasscodeMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodeMinutesOfInactivityBeforeScreenTimeout
    }
}
// GetPasscodePreviousPasscodeBlockCount gets the passcodePreviousPasscodeBlockCount property value. Number of previous passcodes to block. Valid values 1 to 24
func (m *IosGeneralDeviceConfiguration) GetPasscodePreviousPasscodeBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodePreviousPasscodeBlockCount
    }
}
// GetPasscodeRequired gets the passcodeRequired property value. Indicates whether or not to require a passcode.
func (m *IosGeneralDeviceConfiguration) GetPasscodeRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passcodeRequired
    }
}
// GetPasscodeRequiredType gets the passcodeRequiredType property value. Type of passcode that is required. Possible values are: deviceDefault, alphanumeric, numeric.
func (m *IosGeneralDeviceConfiguration) GetPasscodeRequiredType()(*RequiredPasswordType) {
    if m == nil {
        return nil
    } else {
        return m.passcodeRequiredType
    }
}
// GetPasscodeSignInFailureCountBeforeWipe gets the passcodeSignInFailureCountBeforeWipe property value. Number of sign in failures allowed before wiping the device. Valid values 2 to 11
func (m *IosGeneralDeviceConfiguration) GetPasscodeSignInFailureCountBeforeWipe()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodeSignInFailureCountBeforeWipe
    }
}
// GetPodcastsBlocked gets the podcastsBlocked property value. Indicates whether or not to block the user from using podcasts on the supervised device (iOS 8.0 and later).
func (m *IosGeneralDeviceConfiguration) GetPodcastsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.podcastsBlocked
    }
}
// GetSafariBlockAutofill gets the safariBlockAutofill property value. Indicates whether or not to block the user from using Auto fill in Safari. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetSafariBlockAutofill()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.safariBlockAutofill
    }
}
// GetSafariBlocked gets the safariBlocked property value. Indicates whether or not to block the user from using Safari. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetSafariBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.safariBlocked
    }
}
// GetSafariBlockJavaScript gets the safariBlockJavaScript property value. Indicates whether or not to block JavaScript in Safari.
func (m *IosGeneralDeviceConfiguration) GetSafariBlockJavaScript()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.safariBlockJavaScript
    }
}
// GetSafariBlockPopups gets the safariBlockPopups property value. Indicates whether or not to block popups in Safari.
func (m *IosGeneralDeviceConfiguration) GetSafariBlockPopups()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.safariBlockPopups
    }
}
// GetSafariCookieSettings gets the safariCookieSettings property value. Cookie settings for Safari. Possible values are: browserDefault, blockAlways, allowCurrentWebSite, allowFromWebsitesVisited, allowAlways.
func (m *IosGeneralDeviceConfiguration) GetSafariCookieSettings()(*WebBrowserCookieSettings) {
    if m == nil {
        return nil
    } else {
        return m.safariCookieSettings
    }
}
// GetSafariManagedDomains gets the safariManagedDomains property value. URLs matching the patterns listed here will be considered managed.
func (m *IosGeneralDeviceConfiguration) GetSafariManagedDomains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.safariManagedDomains
    }
}
// GetSafariPasswordAutoFillDomains gets the safariPasswordAutoFillDomains property value. Users can save passwords in Safari only from URLs matching the patterns listed here. Applies to devices in supervised mode (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) GetSafariPasswordAutoFillDomains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.safariPasswordAutoFillDomains
    }
}
// GetSafariRequireFraudWarning gets the safariRequireFraudWarning property value. Indicates whether or not to require fraud warning in Safari.
func (m *IosGeneralDeviceConfiguration) GetSafariRequireFraudWarning()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.safariRequireFraudWarning
    }
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether or not to block the user from taking Screenshots.
func (m *IosGeneralDeviceConfiguration) GetScreenCaptureBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.screenCaptureBlocked
    }
}
// GetSiriBlocked gets the siriBlocked property value. Indicates whether or not to block the user from using Siri.
func (m *IosGeneralDeviceConfiguration) GetSiriBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.siriBlocked
    }
}
// GetSiriBlockedWhenLocked gets the siriBlockedWhenLocked property value. Indicates whether or not to block the user from using Siri when locked.
func (m *IosGeneralDeviceConfiguration) GetSiriBlockedWhenLocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.siriBlockedWhenLocked
    }
}
// GetSiriBlockUserGeneratedContent gets the siriBlockUserGeneratedContent property value. Indicates whether or not to block Siri from querying user-generated content when used on a supervised device.
func (m *IosGeneralDeviceConfiguration) GetSiriBlockUserGeneratedContent()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.siriBlockUserGeneratedContent
    }
}
// GetSiriRequireProfanityFilter gets the siriRequireProfanityFilter property value. Indicates whether or not to prevent Siri from dictating, or speaking profane language on supervised device.
func (m *IosGeneralDeviceConfiguration) GetSiriRequireProfanityFilter()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.siriRequireProfanityFilter
    }
}
// GetSpotlightBlockInternetResults gets the spotlightBlockInternetResults property value. Indicates whether or not to block Spotlight search from returning internet results on supervised device.
func (m *IosGeneralDeviceConfiguration) GetSpotlightBlockInternetResults()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.spotlightBlockInternetResults
    }
}
// GetVoiceDialingBlocked gets the voiceDialingBlocked property value. Indicates whether or not to block voice dialing.
func (m *IosGeneralDeviceConfiguration) GetVoiceDialingBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.voiceDialingBlocked
    }
}
// GetWallpaperBlockModification gets the wallpaperBlockModification property value. Indicates whether or not to allow wallpaper modification on supervised device (iOS 9.0 and later) .
func (m *IosGeneralDeviceConfiguration) GetWallpaperBlockModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.wallpaperBlockModification
    }
}
// GetWiFiConnectOnlyToConfiguredNetworks gets the wiFiConnectOnlyToConfiguredNetworks property value. Indicates whether or not to force the device to use only Wi-Fi networks from configuration profiles when the device is in supervised mode. Available for devices running iOS and iPadOS versions 14.4 and earlier. Devices running 14.5+ should use the setting, 'WiFiConnectToAllowedNetworksOnlyForced.
func (m *IosGeneralDeviceConfiguration) GetWiFiConnectOnlyToConfiguredNetworks()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.wiFiConnectOnlyToConfiguredNetworks
    }
}
// Serialize serializes information the current object
func (m *IosGeneralDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountBlockModification", m.GetAccountBlockModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("activationLockAllowWhenSupervised", m.GetActivationLockAllowWhenSupervised())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("airDropBlocked", m.GetAirDropBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("airDropForceUnmanagedDropTarget", m.GetAirDropForceUnmanagedDropTarget())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("airPlayForcePairingPasswordForOutgoingRequests", m.GetAirPlayForcePairingPasswordForOutgoingRequests())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appleNewsBlocked", m.GetAppleNewsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appleWatchBlockPairing", m.GetAppleWatchBlockPairing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appleWatchForceWristDetection", m.GetAppleWatchForceWristDetection())
        if err != nil {
            return err
        }
    }
    if m.GetAppsSingleAppModeList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppsSingleAppModeList()))
        for i, v := range m.GetAppsSingleAppModeList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appsSingleAppModeList", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appStoreBlockAutomaticDownloads", m.GetAppStoreBlockAutomaticDownloads())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appStoreBlocked", m.GetAppStoreBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appStoreBlockInAppPurchases", m.GetAppStoreBlockInAppPurchases())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appStoreBlockUIAppInstallation", m.GetAppStoreBlockUIAppInstallation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appStoreRequirePassword", m.GetAppStoreRequirePassword())
        if err != nil {
            return err
        }
    }
    if m.GetAppsVisibilityList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppsVisibilityList()))
        for i, v := range m.GetAppsVisibilityList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appsVisibilityList", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppsVisibilityListType() != nil {
        cast := (*m.GetAppsVisibilityListType()).String()
        err = writer.WriteStringValue("appsVisibilityListType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bluetoothBlockModification", m.GetBluetoothBlockModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cameraBlocked", m.GetCameraBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockDataRoaming", m.GetCellularBlockDataRoaming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockGlobalBackgroundFetchWhileRoaming", m.GetCellularBlockGlobalBackgroundFetchWhileRoaming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockPerAppDataModification", m.GetCellularBlockPerAppDataModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockPersonalHotspot", m.GetCellularBlockPersonalHotspot())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockVoiceRoaming", m.GetCellularBlockVoiceRoaming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("certificatesBlockUntrustedTlsCertificates", m.GetCertificatesBlockUntrustedTlsCertificates())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("classroomAppBlockRemoteScreenObservation", m.GetClassroomAppBlockRemoteScreenObservation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("classroomAppForceUnpromptedScreenObservation", m.GetClassroomAppForceUnpromptedScreenObservation())
        if err != nil {
            return err
        }
    }
    if m.GetCompliantAppListType() != nil {
        cast := (*m.GetCompliantAppListType()).String()
        err = writer.WriteStringValue("compliantAppListType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCompliantAppsList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompliantAppsList()))
        for i, v := range m.GetCompliantAppsList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("compliantAppsList", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("configurationProfileBlockChanges", m.GetConfigurationProfileBlockChanges())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("definitionLookupBlocked", m.GetDefinitionLookupBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceBlockEnableRestrictions", m.GetDeviceBlockEnableRestrictions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceBlockEraseContentAndSettings", m.GetDeviceBlockEraseContentAndSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceBlockNameModification", m.GetDeviceBlockNameModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("diagnosticDataBlockSubmission", m.GetDiagnosticDataBlockSubmission())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("diagnosticDataBlockSubmissionModification", m.GetDiagnosticDataBlockSubmissionModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("documentsBlockManagedDocumentsInUnmanagedApps", m.GetDocumentsBlockManagedDocumentsInUnmanagedApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("documentsBlockUnmanagedDocumentsInManagedApps", m.GetDocumentsBlockUnmanagedDocumentsInManagedApps())
        if err != nil {
            return err
        }
    }
    if m.GetEmailInDomainSuffixes() != nil {
        err = writer.WriteCollectionOfStringValues("emailInDomainSuffixes", m.GetEmailInDomainSuffixes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enterpriseAppBlockTrust", m.GetEnterpriseAppBlockTrust())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enterpriseAppBlockTrustModification", m.GetEnterpriseAppBlockTrustModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("faceTimeBlocked", m.GetFaceTimeBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("findMyFriendsBlocked", m.GetFindMyFriendsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("gameCenterBlocked", m.GetGameCenterBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("gamingBlockGameCenterFriends", m.GetGamingBlockGameCenterFriends())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("gamingBlockMultiplayer", m.GetGamingBlockMultiplayer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hostPairingBlocked", m.GetHostPairingBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iBooksStoreBlocked", m.GetIBooksStoreBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iBooksStoreBlockErotica", m.GetIBooksStoreBlockErotica())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockActivityContinuation", m.GetICloudBlockActivityContinuation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockBackup", m.GetICloudBlockBackup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockDocumentSync", m.GetICloudBlockDocumentSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockManagedAppsSync", m.GetICloudBlockManagedAppsSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockPhotoLibrary", m.GetICloudBlockPhotoLibrary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockPhotoStreamSync", m.GetICloudBlockPhotoStreamSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockSharedPhotoStream", m.GetICloudBlockSharedPhotoStream())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudRequireEncryptedBackup", m.GetICloudRequireEncryptedBackup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iTunesBlockExplicitContent", m.GetITunesBlockExplicitContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iTunesBlockMusicService", m.GetITunesBlockMusicService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iTunesBlockRadio", m.GetITunesBlockRadio())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keyboardBlockAutoCorrect", m.GetKeyboardBlockAutoCorrect())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keyboardBlockDictation", m.GetKeyboardBlockDictation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keyboardBlockPredictive", m.GetKeyboardBlockPredictive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keyboardBlockShortcuts", m.GetKeyboardBlockShortcuts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keyboardBlockSpellCheck", m.GetKeyboardBlockSpellCheck())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowAssistiveSpeak", m.GetKioskModeAllowAssistiveSpeak())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowAssistiveTouchSettings", m.GetKioskModeAllowAssistiveTouchSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowAutoLock", m.GetKioskModeAllowAutoLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowColorInversionSettings", m.GetKioskModeAllowColorInversionSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowRingerSwitch", m.GetKioskModeAllowRingerSwitch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowScreenRotation", m.GetKioskModeAllowScreenRotation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowSleepButton", m.GetKioskModeAllowSleepButton())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowTouchscreen", m.GetKioskModeAllowTouchscreen())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowVoiceOverSettings", m.GetKioskModeAllowVoiceOverSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowVolumeButtons", m.GetKioskModeAllowVolumeButtons())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowZoomSettings", m.GetKioskModeAllowZoomSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeAppStoreUrl", m.GetKioskModeAppStoreUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeBuiltInAppId", m.GetKioskModeBuiltInAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeManagedAppId", m.GetKioskModeManagedAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeRequireAssistiveTouch", m.GetKioskModeRequireAssistiveTouch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeRequireColorInversion", m.GetKioskModeRequireColorInversion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeRequireMonoAudio", m.GetKioskModeRequireMonoAudio())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeRequireVoiceOver", m.GetKioskModeRequireVoiceOver())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeRequireZoom", m.GetKioskModeRequireZoom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lockScreenBlockControlCenter", m.GetLockScreenBlockControlCenter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lockScreenBlockNotificationView", m.GetLockScreenBlockNotificationView())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lockScreenBlockPassbook", m.GetLockScreenBlockPassbook())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lockScreenBlockTodayView", m.GetLockScreenBlockTodayView())
        if err != nil {
            return err
        }
    }
    if m.GetMediaContentRatingApps() != nil {
        cast := (*m.GetMediaContentRatingApps()).String()
        err = writer.WriteStringValue("mediaContentRatingApps", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingAustralia", m.GetMediaContentRatingAustralia())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingCanada", m.GetMediaContentRatingCanada())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingFrance", m.GetMediaContentRatingFrance())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingGermany", m.GetMediaContentRatingGermany())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingIreland", m.GetMediaContentRatingIreland())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingJapan", m.GetMediaContentRatingJapan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingNewZealand", m.GetMediaContentRatingNewZealand())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingUnitedKingdom", m.GetMediaContentRatingUnitedKingdom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingUnitedStates", m.GetMediaContentRatingUnitedStates())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("messagesBlocked", m.GetMessagesBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkUsageRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkUsageRules()))
        for i, v := range m.GetNetworkUsageRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("networkUsageRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("notificationsBlockSettingsModification", m.GetNotificationsBlockSettingsModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passcodeBlockFingerprintModification", m.GetPasscodeBlockFingerprintModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passcodeBlockFingerprintUnlock", m.GetPasscodeBlockFingerprintUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passcodeBlockModification", m.GetPasscodeBlockModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passcodeBlockSimple", m.GetPasscodeBlockSimple())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeExpirationDays", m.GetPasscodeExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinimumCharacterSetCount", m.GetPasscodeMinimumCharacterSetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinimumLength", m.GetPasscodeMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinutesOfInactivityBeforeLock", m.GetPasscodeMinutesOfInactivityBeforeLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinutesOfInactivityBeforeScreenTimeout", m.GetPasscodeMinutesOfInactivityBeforeScreenTimeout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodePreviousPasscodeBlockCount", m.GetPasscodePreviousPasscodeBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passcodeRequired", m.GetPasscodeRequired())
        if err != nil {
            return err
        }
    }
    if m.GetPasscodeRequiredType() != nil {
        cast := (*m.GetPasscodeRequiredType()).String()
        err = writer.WriteStringValue("passcodeRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeSignInFailureCountBeforeWipe", m.GetPasscodeSignInFailureCountBeforeWipe())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("podcastsBlocked", m.GetPodcastsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("safariBlockAutofill", m.GetSafariBlockAutofill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("safariBlocked", m.GetSafariBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("safariBlockJavaScript", m.GetSafariBlockJavaScript())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("safariBlockPopups", m.GetSafariBlockPopups())
        if err != nil {
            return err
        }
    }
    if m.GetSafariCookieSettings() != nil {
        cast := (*m.GetSafariCookieSettings()).String()
        err = writer.WriteStringValue("safariCookieSettings", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSafariManagedDomains() != nil {
        err = writer.WriteCollectionOfStringValues("safariManagedDomains", m.GetSafariManagedDomains())
        if err != nil {
            return err
        }
    }
    if m.GetSafariPasswordAutoFillDomains() != nil {
        err = writer.WriteCollectionOfStringValues("safariPasswordAutoFillDomains", m.GetSafariPasswordAutoFillDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("safariRequireFraudWarning", m.GetSafariRequireFraudWarning())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("screenCaptureBlocked", m.GetScreenCaptureBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("siriBlocked", m.GetSiriBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("siriBlockedWhenLocked", m.GetSiriBlockedWhenLocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("siriBlockUserGeneratedContent", m.GetSiriBlockUserGeneratedContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("siriRequireProfanityFilter", m.GetSiriRequireProfanityFilter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("spotlightBlockInternetResults", m.GetSpotlightBlockInternetResults())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("voiceDialingBlocked", m.GetVoiceDialingBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wallpaperBlockModification", m.GetWallpaperBlockModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wiFiConnectOnlyToConfiguredNetworks", m.GetWiFiConnectOnlyToConfiguredNetworks())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountBlockModification sets the accountBlockModification property value. Indicates whether or not to allow account modification when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetAccountBlockModification(value *bool)() {
    if m != nil {
        m.accountBlockModification = value
    }
}
// SetActivationLockAllowWhenSupervised sets the activationLockAllowWhenSupervised property value. Indicates whether or not to allow activation lock when the device is in the supervised mode.
func (m *IosGeneralDeviceConfiguration) SetActivationLockAllowWhenSupervised(value *bool)() {
    if m != nil {
        m.activationLockAllowWhenSupervised = value
    }
}
// SetAirDropBlocked sets the airDropBlocked property value. Indicates whether or not to allow AirDrop when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetAirDropBlocked(value *bool)() {
    if m != nil {
        m.airDropBlocked = value
    }
}
// SetAirDropForceUnmanagedDropTarget sets the airDropForceUnmanagedDropTarget property value. Indicates whether or not to cause AirDrop to be considered an unmanaged drop target (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAirDropForceUnmanagedDropTarget(value *bool)() {
    if m != nil {
        m.airDropForceUnmanagedDropTarget = value
    }
}
// SetAirPlayForcePairingPasswordForOutgoingRequests sets the airPlayForcePairingPasswordForOutgoingRequests property value. Indicates whether or not to enforce all devices receiving AirPlay requests from this device to use a pairing password.
func (m *IosGeneralDeviceConfiguration) SetAirPlayForcePairingPasswordForOutgoingRequests(value *bool)() {
    if m != nil {
        m.airPlayForcePairingPasswordForOutgoingRequests = value
    }
}
// SetAppleNewsBlocked sets the appleNewsBlocked property value. Indicates whether or not to block the user from using News when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAppleNewsBlocked(value *bool)() {
    if m != nil {
        m.appleNewsBlocked = value
    }
}
// SetAppleWatchBlockPairing sets the appleWatchBlockPairing property value. Indicates whether or not to allow Apple Watch pairing when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAppleWatchBlockPairing(value *bool)() {
    if m != nil {
        m.appleWatchBlockPairing = value
    }
}
// SetAppleWatchForceWristDetection sets the appleWatchForceWristDetection property value. Indicates whether or not to force a paired Apple Watch to use Wrist Detection (iOS 8.2 and later).
func (m *IosGeneralDeviceConfiguration) SetAppleWatchForceWristDetection(value *bool)() {
    if m != nil {
        m.appleWatchForceWristDetection = value
    }
}
// SetAppsSingleAppModeList sets the appsSingleAppModeList property value. Gets or sets the list of iOS apps allowed to autonomously enter Single App Mode. Supervised only. iOS 7.0 and later. This collection can contain a maximum of 500 elements.
func (m *IosGeneralDeviceConfiguration) SetAppsSingleAppModeList(value []AppListItemable)() {
    if m != nil {
        m.appsSingleAppModeList = value
    }
}
// SetAppStoreBlockAutomaticDownloads sets the appStoreBlockAutomaticDownloads property value. Indicates whether or not to block the automatic downloading of apps purchased on other devices when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAppStoreBlockAutomaticDownloads(value *bool)() {
    if m != nil {
        m.appStoreBlockAutomaticDownloads = value
    }
}
// SetAppStoreBlocked sets the appStoreBlocked property value. Indicates whether or not to block the user from using the App Store. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetAppStoreBlocked(value *bool)() {
    if m != nil {
        m.appStoreBlocked = value
    }
}
// SetAppStoreBlockInAppPurchases sets the appStoreBlockInAppPurchases property value. Indicates whether or not to block the user from making in app purchases.
func (m *IosGeneralDeviceConfiguration) SetAppStoreBlockInAppPurchases(value *bool)() {
    if m != nil {
        m.appStoreBlockInAppPurchases = value
    }
}
// SetAppStoreBlockUIAppInstallation sets the appStoreBlockUIAppInstallation property value. Indicates whether or not to block the App Store app, not restricting installation through Host apps. Applies to supervised mode only (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAppStoreBlockUIAppInstallation(value *bool)() {
    if m != nil {
        m.appStoreBlockUIAppInstallation = value
    }
}
// SetAppStoreRequirePassword sets the appStoreRequirePassword property value. Indicates whether or not to require a password when using the app store.
func (m *IosGeneralDeviceConfiguration) SetAppStoreRequirePassword(value *bool)() {
    if m != nil {
        m.appStoreRequirePassword = value
    }
}
// SetAppsVisibilityList sets the appsVisibilityList property value. List of apps in the visibility list (either visible/launchable apps list or hidden/unlaunchable apps list, controlled by AppsVisibilityListType) (iOS 9.3 and later). This collection can contain a maximum of 10000 elements.
func (m *IosGeneralDeviceConfiguration) SetAppsVisibilityList(value []AppListItemable)() {
    if m != nil {
        m.appsVisibilityList = value
    }
}
// SetAppsVisibilityListType sets the appsVisibilityListType property value. Type of list that is in the AppsVisibilityList. Possible values are: none, appsInListCompliant, appsNotInListCompliant.
func (m *IosGeneralDeviceConfiguration) SetAppsVisibilityListType(value *AppListType)() {
    if m != nil {
        m.appsVisibilityListType = value
    }
}
// SetBluetoothBlockModification sets the bluetoothBlockModification property value. Indicates whether or not to allow modification of Bluetooth settings when the device is in supervised mode (iOS 10.0 and later).
func (m *IosGeneralDeviceConfiguration) SetBluetoothBlockModification(value *bool)() {
    if m != nil {
        m.bluetoothBlockModification = value
    }
}
// SetCameraBlocked sets the cameraBlocked property value. Indicates whether or not to block the user from accessing the camera of the device. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetCameraBlocked(value *bool)() {
    if m != nil {
        m.cameraBlocked = value
    }
}
// SetCellularBlockDataRoaming sets the cellularBlockDataRoaming property value. Indicates whether or not to block data roaming.
func (m *IosGeneralDeviceConfiguration) SetCellularBlockDataRoaming(value *bool)() {
    if m != nil {
        m.cellularBlockDataRoaming = value
    }
}
// SetCellularBlockGlobalBackgroundFetchWhileRoaming sets the cellularBlockGlobalBackgroundFetchWhileRoaming property value. Indicates whether or not to block global background fetch while roaming.
func (m *IosGeneralDeviceConfiguration) SetCellularBlockGlobalBackgroundFetchWhileRoaming(value *bool)() {
    if m != nil {
        m.cellularBlockGlobalBackgroundFetchWhileRoaming = value
    }
}
// SetCellularBlockPerAppDataModification sets the cellularBlockPerAppDataModification property value. Indicates whether or not to allow changes to cellular app data usage settings when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetCellularBlockPerAppDataModification(value *bool)() {
    if m != nil {
        m.cellularBlockPerAppDataModification = value
    }
}
// SetCellularBlockPersonalHotspot sets the cellularBlockPersonalHotspot property value. Indicates whether or not to block Personal Hotspot.
func (m *IosGeneralDeviceConfiguration) SetCellularBlockPersonalHotspot(value *bool)() {
    if m != nil {
        m.cellularBlockPersonalHotspot = value
    }
}
// SetCellularBlockVoiceRoaming sets the cellularBlockVoiceRoaming property value. Indicates whether or not to block voice roaming.
func (m *IosGeneralDeviceConfiguration) SetCellularBlockVoiceRoaming(value *bool)() {
    if m != nil {
        m.cellularBlockVoiceRoaming = value
    }
}
// SetCertificatesBlockUntrustedTlsCertificates sets the certificatesBlockUntrustedTlsCertificates property value. Indicates whether or not to block untrusted TLS certificates.
func (m *IosGeneralDeviceConfiguration) SetCertificatesBlockUntrustedTlsCertificates(value *bool)() {
    if m != nil {
        m.certificatesBlockUntrustedTlsCertificates = value
    }
}
// SetClassroomAppBlockRemoteScreenObservation sets the classroomAppBlockRemoteScreenObservation property value. Indicates whether or not to allow remote screen observation by Classroom app when the device is in supervised mode (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) SetClassroomAppBlockRemoteScreenObservation(value *bool)() {
    if m != nil {
        m.classroomAppBlockRemoteScreenObservation = value
    }
}
// SetClassroomAppForceUnpromptedScreenObservation sets the classroomAppForceUnpromptedScreenObservation property value. Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to view a student's screen without prompting when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetClassroomAppForceUnpromptedScreenObservation(value *bool)() {
    if m != nil {
        m.classroomAppForceUnpromptedScreenObservation = value
    }
}
// SetCompliantAppListType sets the compliantAppListType property value. List that is in the AppComplianceList. Possible values are: none, appsInListCompliant, appsNotInListCompliant.
func (m *IosGeneralDeviceConfiguration) SetCompliantAppListType(value *AppListType)() {
    if m != nil {
        m.compliantAppListType = value
    }
}
// SetCompliantAppsList sets the compliantAppsList property value. List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
func (m *IosGeneralDeviceConfiguration) SetCompliantAppsList(value []AppListItemable)() {
    if m != nil {
        m.compliantAppsList = value
    }
}
// SetConfigurationProfileBlockChanges sets the configurationProfileBlockChanges property value. Indicates whether or not to block the user from installing configuration profiles and certificates interactively when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetConfigurationProfileBlockChanges(value *bool)() {
    if m != nil {
        m.configurationProfileBlockChanges = value
    }
}
// SetDefinitionLookupBlocked sets the definitionLookupBlocked property value. Indicates whether or not to block definition lookup when the device is in supervised mode (iOS 8.1.3 and later ).
func (m *IosGeneralDeviceConfiguration) SetDefinitionLookupBlocked(value *bool)() {
    if m != nil {
        m.definitionLookupBlocked = value
    }
}
// SetDeviceBlockEnableRestrictions sets the deviceBlockEnableRestrictions property value. Indicates whether or not to allow the user to enables restrictions in the device settings when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetDeviceBlockEnableRestrictions(value *bool)() {
    if m != nil {
        m.deviceBlockEnableRestrictions = value
    }
}
// SetDeviceBlockEraseContentAndSettings sets the deviceBlockEraseContentAndSettings property value. Indicates whether or not to allow the use of the 'Erase all content and settings' option on the device when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetDeviceBlockEraseContentAndSettings(value *bool)() {
    if m != nil {
        m.deviceBlockEraseContentAndSettings = value
    }
}
// SetDeviceBlockNameModification sets the deviceBlockNameModification property value. Indicates whether or not to allow device name modification when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetDeviceBlockNameModification(value *bool)() {
    if m != nil {
        m.deviceBlockNameModification = value
    }
}
// SetDiagnosticDataBlockSubmission sets the diagnosticDataBlockSubmission property value. Indicates whether or not to block diagnostic data submission.
func (m *IosGeneralDeviceConfiguration) SetDiagnosticDataBlockSubmission(value *bool)() {
    if m != nil {
        m.diagnosticDataBlockSubmission = value
    }
}
// SetDiagnosticDataBlockSubmissionModification sets the diagnosticDataBlockSubmissionModification property value. Indicates whether or not to allow diagnostics submission settings modification when the device is in supervised mode (iOS 9.3.2 and later).
func (m *IosGeneralDeviceConfiguration) SetDiagnosticDataBlockSubmissionModification(value *bool)() {
    if m != nil {
        m.diagnosticDataBlockSubmissionModification = value
    }
}
// SetDocumentsBlockManagedDocumentsInUnmanagedApps sets the documentsBlockManagedDocumentsInUnmanagedApps property value. Indicates whether or not to block the user from viewing managed documents in unmanaged apps.
func (m *IosGeneralDeviceConfiguration) SetDocumentsBlockManagedDocumentsInUnmanagedApps(value *bool)() {
    if m != nil {
        m.documentsBlockManagedDocumentsInUnmanagedApps = value
    }
}
// SetDocumentsBlockUnmanagedDocumentsInManagedApps sets the documentsBlockUnmanagedDocumentsInManagedApps property value. Indicates whether or not to block the user from viewing unmanaged documents in managed apps.
func (m *IosGeneralDeviceConfiguration) SetDocumentsBlockUnmanagedDocumentsInManagedApps(value *bool)() {
    if m != nil {
        m.documentsBlockUnmanagedDocumentsInManagedApps = value
    }
}
// SetEmailInDomainSuffixes sets the emailInDomainSuffixes property value. An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
func (m *IosGeneralDeviceConfiguration) SetEmailInDomainSuffixes(value []string)() {
    if m != nil {
        m.emailInDomainSuffixes = value
    }
}
// SetEnterpriseAppBlockTrust sets the enterpriseAppBlockTrust property value. Indicates whether or not to block the user from trusting an enterprise app.
func (m *IosGeneralDeviceConfiguration) SetEnterpriseAppBlockTrust(value *bool)() {
    if m != nil {
        m.enterpriseAppBlockTrust = value
    }
}
// SetEnterpriseAppBlockTrustModification sets the enterpriseAppBlockTrustModification property value. [Deprecated] Configuring this setting and setting the value to 'true' has no effect on the device.
func (m *IosGeneralDeviceConfiguration) SetEnterpriseAppBlockTrustModification(value *bool)() {
    if m != nil {
        m.enterpriseAppBlockTrustModification = value
    }
}
// SetFaceTimeBlocked sets the faceTimeBlocked property value. Indicates whether or not to block the user from using FaceTime. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetFaceTimeBlocked(value *bool)() {
    if m != nil {
        m.faceTimeBlocked = value
    }
}
// SetFindMyFriendsBlocked sets the findMyFriendsBlocked property value. Indicates whether or not to block changes to Find My Friends when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetFindMyFriendsBlocked(value *bool)() {
    if m != nil {
        m.findMyFriendsBlocked = value
    }
}
// SetGameCenterBlocked sets the gameCenterBlocked property value. Indicates whether or not to block the user from using Game Center when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetGameCenterBlocked(value *bool)() {
    if m != nil {
        m.gameCenterBlocked = value
    }
}
// SetGamingBlockGameCenterFriends sets the gamingBlockGameCenterFriends property value. Indicates whether or not to block the user from having friends in Game Center. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetGamingBlockGameCenterFriends(value *bool)() {
    if m != nil {
        m.gamingBlockGameCenterFriends = value
    }
}
// SetGamingBlockMultiplayer sets the gamingBlockMultiplayer property value. Indicates whether or not to block the user from using multiplayer gaming. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetGamingBlockMultiplayer(value *bool)() {
    if m != nil {
        m.gamingBlockMultiplayer = value
    }
}
// SetHostPairingBlocked sets the hostPairingBlocked property value. indicates whether or not to allow host pairing to control the devices an iOS device can pair with when the iOS device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetHostPairingBlocked(value *bool)() {
    if m != nil {
        m.hostPairingBlocked = value
    }
}
// SetIBooksStoreBlocked sets the iBooksStoreBlocked property value. Indicates whether or not to block the user from using the iBooks Store when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetIBooksStoreBlocked(value *bool)() {
    if m != nil {
        m.iBooksStoreBlocked = value
    }
}
// SetIBooksStoreBlockErotica sets the iBooksStoreBlockErotica property value. Indicates whether or not to block the user from downloading media from the iBookstore that has been tagged as erotica.
func (m *IosGeneralDeviceConfiguration) SetIBooksStoreBlockErotica(value *bool)() {
    if m != nil {
        m.iBooksStoreBlockErotica = value
    }
}
// SetICloudBlockActivityContinuation sets the iCloudBlockActivityContinuation property value. Indicates whether or not to block the user from continuing work they started on iOS device to another iOS or macOS device.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockActivityContinuation(value *bool)() {
    if m != nil {
        m.iCloudBlockActivityContinuation = value
    }
}
// SetICloudBlockBackup sets the iCloudBlockBackup property value. Indicates whether or not to block iCloud backup. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockBackup(value *bool)() {
    if m != nil {
        m.iCloudBlockBackup = value
    }
}
// SetICloudBlockDocumentSync sets the iCloudBlockDocumentSync property value. Indicates whether or not to block iCloud document sync. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockDocumentSync(value *bool)() {
    if m != nil {
        m.iCloudBlockDocumentSync = value
    }
}
// SetICloudBlockManagedAppsSync sets the iCloudBlockManagedAppsSync property value. Indicates whether or not to block Managed Apps Cloud Sync.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockManagedAppsSync(value *bool)() {
    if m != nil {
        m.iCloudBlockManagedAppsSync = value
    }
}
// SetICloudBlockPhotoLibrary sets the iCloudBlockPhotoLibrary property value. Indicates whether or not to block iCloud Photo Library.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockPhotoLibrary(value *bool)() {
    if m != nil {
        m.iCloudBlockPhotoLibrary = value
    }
}
// SetICloudBlockPhotoStreamSync sets the iCloudBlockPhotoStreamSync property value. Indicates whether or not to block iCloud Photo Stream Sync.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockPhotoStreamSync(value *bool)() {
    if m != nil {
        m.iCloudBlockPhotoStreamSync = value
    }
}
// SetICloudBlockSharedPhotoStream sets the iCloudBlockSharedPhotoStream property value. Indicates whether or not to block Shared Photo Stream.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockSharedPhotoStream(value *bool)() {
    if m != nil {
        m.iCloudBlockSharedPhotoStream = value
    }
}
// SetICloudRequireEncryptedBackup sets the iCloudRequireEncryptedBackup property value. Indicates whether or not to require backups to iCloud be encrypted.
func (m *IosGeneralDeviceConfiguration) SetICloudRequireEncryptedBackup(value *bool)() {
    if m != nil {
        m.iCloudRequireEncryptedBackup = value
    }
}
// SetITunesBlockExplicitContent sets the iTunesBlockExplicitContent property value. Indicates whether or not to block the user from accessing explicit content in iTunes and the App Store. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetITunesBlockExplicitContent(value *bool)() {
    if m != nil {
        m.iTunesBlockExplicitContent = value
    }
}
// SetITunesBlockMusicService sets the iTunesBlockMusicService property value. Indicates whether or not to block Music service and revert Music app to classic mode when the device is in supervised mode (iOS 9.3 and later and macOS 10.12 and later).
func (m *IosGeneralDeviceConfiguration) SetITunesBlockMusicService(value *bool)() {
    if m != nil {
        m.iTunesBlockMusicService = value
    }
}
// SetITunesBlockRadio sets the iTunesBlockRadio property value. Indicates whether or not to block the user from using iTunes Radio when the device is in supervised mode (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) SetITunesBlockRadio(value *bool)() {
    if m != nil {
        m.iTunesBlockRadio = value
    }
}
// SetKeyboardBlockAutoCorrect sets the keyboardBlockAutoCorrect property value. Indicates whether or not to block keyboard auto-correction when the device is in supervised mode (iOS 8.1.3 and later).
func (m *IosGeneralDeviceConfiguration) SetKeyboardBlockAutoCorrect(value *bool)() {
    if m != nil {
        m.keyboardBlockAutoCorrect = value
    }
}
// SetKeyboardBlockDictation sets the keyboardBlockDictation property value. Indicates whether or not to block the user from using dictation input when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetKeyboardBlockDictation(value *bool)() {
    if m != nil {
        m.keyboardBlockDictation = value
    }
}
// SetKeyboardBlockPredictive sets the keyboardBlockPredictive property value. Indicates whether or not to block predictive keyboards when device is in supervised mode (iOS 8.1.3 and later).
func (m *IosGeneralDeviceConfiguration) SetKeyboardBlockPredictive(value *bool)() {
    if m != nil {
        m.keyboardBlockPredictive = value
    }
}
// SetKeyboardBlockShortcuts sets the keyboardBlockShortcuts property value. Indicates whether or not to block keyboard shortcuts when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetKeyboardBlockShortcuts(value *bool)() {
    if m != nil {
        m.keyboardBlockShortcuts = value
    }
}
// SetKeyboardBlockSpellCheck sets the keyboardBlockSpellCheck property value. Indicates whether or not to block keyboard spell-checking when the device is in supervised mode (iOS 8.1.3 and later).
func (m *IosGeneralDeviceConfiguration) SetKeyboardBlockSpellCheck(value *bool)() {
    if m != nil {
        m.keyboardBlockSpellCheck = value
    }
}
// SetKioskModeAllowAssistiveSpeak sets the kioskModeAllowAssistiveSpeak property value. Indicates whether or not to allow assistive speak while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowAssistiveSpeak(value *bool)() {
    if m != nil {
        m.kioskModeAllowAssistiveSpeak = value
    }
}
// SetKioskModeAllowAssistiveTouchSettings sets the kioskModeAllowAssistiveTouchSettings property value. Indicates whether or not to allow access to the Assistive Touch Settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowAssistiveTouchSettings(value *bool)() {
    if m != nil {
        m.kioskModeAllowAssistiveTouchSettings = value
    }
}
// SetKioskModeAllowAutoLock sets the kioskModeAllowAutoLock property value. Indicates whether or not to allow device auto lock while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockAutoLock instead.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowAutoLock(value *bool)() {
    if m != nil {
        m.kioskModeAllowAutoLock = value
    }
}
// SetKioskModeAllowColorInversionSettings sets the kioskModeAllowColorInversionSettings property value. Indicates whether or not to allow access to the Color Inversion Settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowColorInversionSettings(value *bool)() {
    if m != nil {
        m.kioskModeAllowColorInversionSettings = value
    }
}
// SetKioskModeAllowRingerSwitch sets the kioskModeAllowRingerSwitch property value. Indicates whether or not to allow use of the ringer switch while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockRingerSwitch instead.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowRingerSwitch(value *bool)() {
    if m != nil {
        m.kioskModeAllowRingerSwitch = value
    }
}
// SetKioskModeAllowScreenRotation sets the kioskModeAllowScreenRotation property value. Indicates whether or not to allow screen rotation while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockScreenRotation instead.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowScreenRotation(value *bool)() {
    if m != nil {
        m.kioskModeAllowScreenRotation = value
    }
}
// SetKioskModeAllowSleepButton sets the kioskModeAllowSleepButton property value. Indicates whether or not to allow use of the sleep button while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockSleepButton instead.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowSleepButton(value *bool)() {
    if m != nil {
        m.kioskModeAllowSleepButton = value
    }
}
// SetKioskModeAllowTouchscreen sets the kioskModeAllowTouchscreen property value. Indicates whether or not to allow use of the touchscreen while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockTouchscreen instead.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowTouchscreen(value *bool)() {
    if m != nil {
        m.kioskModeAllowTouchscreen = value
    }
}
// SetKioskModeAllowVoiceOverSettings sets the kioskModeAllowVoiceOverSettings property value. Indicates whether or not to allow access to the voice over settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowVoiceOverSettings(value *bool)() {
    if m != nil {
        m.kioskModeAllowVoiceOverSettings = value
    }
}
// SetKioskModeAllowVolumeButtons sets the kioskModeAllowVolumeButtons property value. Indicates whether or not to allow use of the volume buttons while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockVolumeButtons instead.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowVolumeButtons(value *bool)() {
    if m != nil {
        m.kioskModeAllowVolumeButtons = value
    }
}
// SetKioskModeAllowZoomSettings sets the kioskModeAllowZoomSettings property value. Indicates whether or not to allow access to the zoom settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowZoomSettings(value *bool)() {
    if m != nil {
        m.kioskModeAllowZoomSettings = value
    }
}
// SetKioskModeAppStoreUrl sets the kioskModeAppStoreUrl property value. URL in the app store to the app to use for kiosk mode. Use if KioskModeManagedAppId is not known.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAppStoreUrl(value *string)() {
    if m != nil {
        m.kioskModeAppStoreUrl = value
    }
}
// SetKioskModeBuiltInAppId sets the kioskModeBuiltInAppId property value. ID for built-in apps to use for kiosk mode. Used when KioskModeManagedAppId and KioskModeAppStoreUrl are not set.
func (m *IosGeneralDeviceConfiguration) SetKioskModeBuiltInAppId(value *string)() {
    if m != nil {
        m.kioskModeBuiltInAppId = value
    }
}
// SetKioskModeManagedAppId sets the kioskModeManagedAppId property value. Managed app id of the app to use for kiosk mode. If KioskModeManagedAppId is specified then KioskModeAppStoreUrl will be ignored.
func (m *IosGeneralDeviceConfiguration) SetKioskModeManagedAppId(value *string)() {
    if m != nil {
        m.kioskModeManagedAppId = value
    }
}
// SetKioskModeRequireAssistiveTouch sets the kioskModeRequireAssistiveTouch property value. Indicates whether or not to require assistive touch while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeRequireAssistiveTouch(value *bool)() {
    if m != nil {
        m.kioskModeRequireAssistiveTouch = value
    }
}
// SetKioskModeRequireColorInversion sets the kioskModeRequireColorInversion property value. Indicates whether or not to require color inversion while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeRequireColorInversion(value *bool)() {
    if m != nil {
        m.kioskModeRequireColorInversion = value
    }
}
// SetKioskModeRequireMonoAudio sets the kioskModeRequireMonoAudio property value. Indicates whether or not to require mono audio while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeRequireMonoAudio(value *bool)() {
    if m != nil {
        m.kioskModeRequireMonoAudio = value
    }
}
// SetKioskModeRequireVoiceOver sets the kioskModeRequireVoiceOver property value. Indicates whether or not to require voice over while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeRequireVoiceOver(value *bool)() {
    if m != nil {
        m.kioskModeRequireVoiceOver = value
    }
}
// SetKioskModeRequireZoom sets the kioskModeRequireZoom property value. Indicates whether or not to require zoom while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeRequireZoom(value *bool)() {
    if m != nil {
        m.kioskModeRequireZoom = value
    }
}
// SetLockScreenBlockControlCenter sets the lockScreenBlockControlCenter property value. Indicates whether or not to block the user from using control center on the lock screen.
func (m *IosGeneralDeviceConfiguration) SetLockScreenBlockControlCenter(value *bool)() {
    if m != nil {
        m.lockScreenBlockControlCenter = value
    }
}
// SetLockScreenBlockNotificationView sets the lockScreenBlockNotificationView property value. Indicates whether or not to block the user from using the notification view on the lock screen.
func (m *IosGeneralDeviceConfiguration) SetLockScreenBlockNotificationView(value *bool)() {
    if m != nil {
        m.lockScreenBlockNotificationView = value
    }
}
// SetLockScreenBlockPassbook sets the lockScreenBlockPassbook property value. Indicates whether or not to block the user from using passbook when the device is locked.
func (m *IosGeneralDeviceConfiguration) SetLockScreenBlockPassbook(value *bool)() {
    if m != nil {
        m.lockScreenBlockPassbook = value
    }
}
// SetLockScreenBlockTodayView sets the lockScreenBlockTodayView property value. Indicates whether or not to block the user from using the Today View on the lock screen.
func (m *IosGeneralDeviceConfiguration) SetLockScreenBlockTodayView(value *bool)() {
    if m != nil {
        m.lockScreenBlockTodayView = value
    }
}
// SetMediaContentRatingApps sets the mediaContentRatingApps property value. Media content rating settings for Apps. Possible values are: allAllowed, allBlocked, agesAbove4, agesAbove9, agesAbove12, agesAbove17.
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingApps(value *RatingAppsType)() {
    if m != nil {
        m.mediaContentRatingApps = value
    }
}
// SetMediaContentRatingAustralia sets the mediaContentRatingAustralia property value. Media content rating settings for Australia
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingAustralia(value MediaContentRatingAustraliaable)() {
    if m != nil {
        m.mediaContentRatingAustralia = value
    }
}
// SetMediaContentRatingCanada sets the mediaContentRatingCanada property value. Media content rating settings for Canada
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingCanada(value MediaContentRatingCanadaable)() {
    if m != nil {
        m.mediaContentRatingCanada = value
    }
}
// SetMediaContentRatingFrance sets the mediaContentRatingFrance property value. Media content rating settings for France
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingFrance(value MediaContentRatingFranceable)() {
    if m != nil {
        m.mediaContentRatingFrance = value
    }
}
// SetMediaContentRatingGermany sets the mediaContentRatingGermany property value. Media content rating settings for Germany
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingGermany(value MediaContentRatingGermanyable)() {
    if m != nil {
        m.mediaContentRatingGermany = value
    }
}
// SetMediaContentRatingIreland sets the mediaContentRatingIreland property value. Media content rating settings for Ireland
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingIreland(value MediaContentRatingIrelandable)() {
    if m != nil {
        m.mediaContentRatingIreland = value
    }
}
// SetMediaContentRatingJapan sets the mediaContentRatingJapan property value. Media content rating settings for Japan
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingJapan(value MediaContentRatingJapanable)() {
    if m != nil {
        m.mediaContentRatingJapan = value
    }
}
// SetMediaContentRatingNewZealand sets the mediaContentRatingNewZealand property value. Media content rating settings for New Zealand
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingNewZealand(value MediaContentRatingNewZealandable)() {
    if m != nil {
        m.mediaContentRatingNewZealand = value
    }
}
// SetMediaContentRatingUnitedKingdom sets the mediaContentRatingUnitedKingdom property value. Media content rating settings for United Kingdom
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingUnitedKingdom(value MediaContentRatingUnitedKingdomable)() {
    if m != nil {
        m.mediaContentRatingUnitedKingdom = value
    }
}
// SetMediaContentRatingUnitedStates sets the mediaContentRatingUnitedStates property value. Media content rating settings for United States
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingUnitedStates(value MediaContentRatingUnitedStatesable)() {
    if m != nil {
        m.mediaContentRatingUnitedStates = value
    }
}
// SetMessagesBlocked sets the messagesBlocked property value. Indicates whether or not to block the user from using the Messages app on the supervised device.
func (m *IosGeneralDeviceConfiguration) SetMessagesBlocked(value *bool)() {
    if m != nil {
        m.messagesBlocked = value
    }
}
// SetNetworkUsageRules sets the networkUsageRules property value. List of managed apps and the network rules that applies to them. This collection can contain a maximum of 1000 elements.
func (m *IosGeneralDeviceConfiguration) SetNetworkUsageRules(value []IosNetworkUsageRuleable)() {
    if m != nil {
        m.networkUsageRules = value
    }
}
// SetNotificationsBlockSettingsModification sets the notificationsBlockSettingsModification property value. Indicates whether or not to allow notifications settings modification (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) SetNotificationsBlockSettingsModification(value *bool)() {
    if m != nil {
        m.notificationsBlockSettingsModification = value
    }
}
// SetPasscodeBlockFingerprintModification sets the passcodeBlockFingerprintModification property value. Block modification of registered Touch ID fingerprints when in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetPasscodeBlockFingerprintModification(value *bool)() {
    if m != nil {
        m.passcodeBlockFingerprintModification = value
    }
}
// SetPasscodeBlockFingerprintUnlock sets the passcodeBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *IosGeneralDeviceConfiguration) SetPasscodeBlockFingerprintUnlock(value *bool)() {
    if m != nil {
        m.passcodeBlockFingerprintUnlock = value
    }
}
// SetPasscodeBlockModification sets the passcodeBlockModification property value. Indicates whether or not to allow passcode modification on the supervised device (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetPasscodeBlockModification(value *bool)() {
    if m != nil {
        m.passcodeBlockModification = value
    }
}
// SetPasscodeBlockSimple sets the passcodeBlockSimple property value. Indicates whether or not to block simple passcodes.
func (m *IosGeneralDeviceConfiguration) SetPasscodeBlockSimple(value *bool)() {
    if m != nil {
        m.passcodeBlockSimple = value
    }
}
// SetPasscodeExpirationDays sets the passcodeExpirationDays property value. Number of days before the passcode expires. Valid values 1 to 65535
func (m *IosGeneralDeviceConfiguration) SetPasscodeExpirationDays(value *int32)() {
    if m != nil {
        m.passcodeExpirationDays = value
    }
}
// SetPasscodeMinimumCharacterSetCount sets the passcodeMinimumCharacterSetCount property value. Number of character sets a passcode must contain. Valid values 0 to 4
func (m *IosGeneralDeviceConfiguration) SetPasscodeMinimumCharacterSetCount(value *int32)() {
    if m != nil {
        m.passcodeMinimumCharacterSetCount = value
    }
}
// SetPasscodeMinimumLength sets the passcodeMinimumLength property value. Minimum length of passcode. Valid values 4 to 14
func (m *IosGeneralDeviceConfiguration) SetPasscodeMinimumLength(value *int32)() {
    if m != nil {
        m.passcodeMinimumLength = value
    }
}
// SetPasscodeMinutesOfInactivityBeforeLock sets the passcodeMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a passcode is required.
func (m *IosGeneralDeviceConfiguration) SetPasscodeMinutesOfInactivityBeforeLock(value *int32)() {
    if m != nil {
        m.passcodeMinutesOfInactivityBeforeLock = value
    }
}
// SetPasscodeMinutesOfInactivityBeforeScreenTimeout sets the passcodeMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *IosGeneralDeviceConfiguration) SetPasscodeMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    if m != nil {
        m.passcodeMinutesOfInactivityBeforeScreenTimeout = value
    }
}
// SetPasscodePreviousPasscodeBlockCount sets the passcodePreviousPasscodeBlockCount property value. Number of previous passcodes to block. Valid values 1 to 24
func (m *IosGeneralDeviceConfiguration) SetPasscodePreviousPasscodeBlockCount(value *int32)() {
    if m != nil {
        m.passcodePreviousPasscodeBlockCount = value
    }
}
// SetPasscodeRequired sets the passcodeRequired property value. Indicates whether or not to require a passcode.
func (m *IosGeneralDeviceConfiguration) SetPasscodeRequired(value *bool)() {
    if m != nil {
        m.passcodeRequired = value
    }
}
// SetPasscodeRequiredType sets the passcodeRequiredType property value. Type of passcode that is required. Possible values are: deviceDefault, alphanumeric, numeric.
func (m *IosGeneralDeviceConfiguration) SetPasscodeRequiredType(value *RequiredPasswordType)() {
    if m != nil {
        m.passcodeRequiredType = value
    }
}
// SetPasscodeSignInFailureCountBeforeWipe sets the passcodeSignInFailureCountBeforeWipe property value. Number of sign in failures allowed before wiping the device. Valid values 2 to 11
func (m *IosGeneralDeviceConfiguration) SetPasscodeSignInFailureCountBeforeWipe(value *int32)() {
    if m != nil {
        m.passcodeSignInFailureCountBeforeWipe = value
    }
}
// SetPodcastsBlocked sets the podcastsBlocked property value. Indicates whether or not to block the user from using podcasts on the supervised device (iOS 8.0 and later).
func (m *IosGeneralDeviceConfiguration) SetPodcastsBlocked(value *bool)() {
    if m != nil {
        m.podcastsBlocked = value
    }
}
// SetSafariBlockAutofill sets the safariBlockAutofill property value. Indicates whether or not to block the user from using Auto fill in Safari. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetSafariBlockAutofill(value *bool)() {
    if m != nil {
        m.safariBlockAutofill = value
    }
}
// SetSafariBlocked sets the safariBlocked property value. Indicates whether or not to block the user from using Safari. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetSafariBlocked(value *bool)() {
    if m != nil {
        m.safariBlocked = value
    }
}
// SetSafariBlockJavaScript sets the safariBlockJavaScript property value. Indicates whether or not to block JavaScript in Safari.
func (m *IosGeneralDeviceConfiguration) SetSafariBlockJavaScript(value *bool)() {
    if m != nil {
        m.safariBlockJavaScript = value
    }
}
// SetSafariBlockPopups sets the safariBlockPopups property value. Indicates whether or not to block popups in Safari.
func (m *IosGeneralDeviceConfiguration) SetSafariBlockPopups(value *bool)() {
    if m != nil {
        m.safariBlockPopups = value
    }
}
// SetSafariCookieSettings sets the safariCookieSettings property value. Cookie settings for Safari. Possible values are: browserDefault, blockAlways, allowCurrentWebSite, allowFromWebsitesVisited, allowAlways.
func (m *IosGeneralDeviceConfiguration) SetSafariCookieSettings(value *WebBrowserCookieSettings)() {
    if m != nil {
        m.safariCookieSettings = value
    }
}
// SetSafariManagedDomains sets the safariManagedDomains property value. URLs matching the patterns listed here will be considered managed.
func (m *IosGeneralDeviceConfiguration) SetSafariManagedDomains(value []string)() {
    if m != nil {
        m.safariManagedDomains = value
    }
}
// SetSafariPasswordAutoFillDomains sets the safariPasswordAutoFillDomains property value. Users can save passwords in Safari only from URLs matching the patterns listed here. Applies to devices in supervised mode (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) SetSafariPasswordAutoFillDomains(value []string)() {
    if m != nil {
        m.safariPasswordAutoFillDomains = value
    }
}
// SetSafariRequireFraudWarning sets the safariRequireFraudWarning property value. Indicates whether or not to require fraud warning in Safari.
func (m *IosGeneralDeviceConfiguration) SetSafariRequireFraudWarning(value *bool)() {
    if m != nil {
        m.safariRequireFraudWarning = value
    }
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether or not to block the user from taking Screenshots.
func (m *IosGeneralDeviceConfiguration) SetScreenCaptureBlocked(value *bool)() {
    if m != nil {
        m.screenCaptureBlocked = value
    }
}
// SetSiriBlocked sets the siriBlocked property value. Indicates whether or not to block the user from using Siri.
func (m *IosGeneralDeviceConfiguration) SetSiriBlocked(value *bool)() {
    if m != nil {
        m.siriBlocked = value
    }
}
// SetSiriBlockedWhenLocked sets the siriBlockedWhenLocked property value. Indicates whether or not to block the user from using Siri when locked.
func (m *IosGeneralDeviceConfiguration) SetSiriBlockedWhenLocked(value *bool)() {
    if m != nil {
        m.siriBlockedWhenLocked = value
    }
}
// SetSiriBlockUserGeneratedContent sets the siriBlockUserGeneratedContent property value. Indicates whether or not to block Siri from querying user-generated content when used on a supervised device.
func (m *IosGeneralDeviceConfiguration) SetSiriBlockUserGeneratedContent(value *bool)() {
    if m != nil {
        m.siriBlockUserGeneratedContent = value
    }
}
// SetSiriRequireProfanityFilter sets the siriRequireProfanityFilter property value. Indicates whether or not to prevent Siri from dictating, or speaking profane language on supervised device.
func (m *IosGeneralDeviceConfiguration) SetSiriRequireProfanityFilter(value *bool)() {
    if m != nil {
        m.siriRequireProfanityFilter = value
    }
}
// SetSpotlightBlockInternetResults sets the spotlightBlockInternetResults property value. Indicates whether or not to block Spotlight search from returning internet results on supervised device.
func (m *IosGeneralDeviceConfiguration) SetSpotlightBlockInternetResults(value *bool)() {
    if m != nil {
        m.spotlightBlockInternetResults = value
    }
}
// SetVoiceDialingBlocked sets the voiceDialingBlocked property value. Indicates whether or not to block voice dialing.
func (m *IosGeneralDeviceConfiguration) SetVoiceDialingBlocked(value *bool)() {
    if m != nil {
        m.voiceDialingBlocked = value
    }
}
// SetWallpaperBlockModification sets the wallpaperBlockModification property value. Indicates whether or not to allow wallpaper modification on supervised device (iOS 9.0 and later) .
func (m *IosGeneralDeviceConfiguration) SetWallpaperBlockModification(value *bool)() {
    if m != nil {
        m.wallpaperBlockModification = value
    }
}
// SetWiFiConnectOnlyToConfiguredNetworks sets the wiFiConnectOnlyToConfiguredNetworks property value. Indicates whether or not to force the device to use only Wi-Fi networks from configuration profiles when the device is in supervised mode. Available for devices running iOS and iPadOS versions 14.4 and earlier. Devices running 14.5+ should use the setting, 'WiFiConnectToAllowedNetworksOnlyForced.
func (m *IosGeneralDeviceConfiguration) SetWiFiConnectOnlyToConfiguredNetworks(value *bool)() {
    if m != nil {
        m.wiFiConnectOnlyToConfiguredNetworks = value
    }
}
