package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdateForBusinessConfiguration 
type WindowsUpdateForBusinessConfiguration struct {
    DeviceConfiguration
    // Possible values for automatic update mode.
    automaticUpdateMode *AutomaticUpdateMode
    // Which branch devices will receive their updates from
    businessReadyUpdatesOnly *WindowsUpdateType
    // Delivery optimization mode for peer distribution
    deliveryOptimizationMode *WindowsDeliveryOptimizationMode
    // Exclude Windows update Drivers
    driversExcluded *bool
    // Defer Feature Updates by these many days
    featureUpdatesDeferralPeriodInDays *int32
    // Pause Feature Updates
    featureUpdatesPaused *bool
    // Feature Updates Pause Expiry datetime
    featureUpdatesPauseExpiryDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Installation schedule
    installationSchedule WindowsUpdateInstallScheduleTypeable
    // Allow Microsoft Update Service
    microsoftUpdateServiceAllowed *bool
    // Possible values for pre-release features.
    prereleaseFeatures *PrereleaseFeatures
    // Defer Quality Updates by these many days
    qualityUpdatesDeferralPeriodInDays *int32
    // Pause Quality Updates
    qualityUpdatesPaused *bool
    // Quality Updates Pause Expiry datetime
    qualityUpdatesPauseExpiryDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewWindowsUpdateForBusinessConfiguration instantiates a new WindowsUpdateForBusinessConfiguration and sets the default values.
func NewWindowsUpdateForBusinessConfiguration()(*WindowsUpdateForBusinessConfiguration) {
    m := &WindowsUpdateForBusinessConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    return m
}
// CreateWindowsUpdateForBusinessConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdateForBusinessConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdateForBusinessConfiguration(), nil
}
// GetAutomaticUpdateMode gets the automaticUpdateMode property value. Possible values for automatic update mode.
func (m *WindowsUpdateForBusinessConfiguration) GetAutomaticUpdateMode()(*AutomaticUpdateMode) {
    if m == nil {
        return nil
    } else {
        return m.automaticUpdateMode
    }
}
// GetBusinessReadyUpdatesOnly gets the businessReadyUpdatesOnly property value. Which branch devices will receive their updates from
func (m *WindowsUpdateForBusinessConfiguration) GetBusinessReadyUpdatesOnly()(*WindowsUpdateType) {
    if m == nil {
        return nil
    } else {
        return m.businessReadyUpdatesOnly
    }
}
// GetDeliveryOptimizationMode gets the deliveryOptimizationMode property value. Delivery optimization mode for peer distribution
func (m *WindowsUpdateForBusinessConfiguration) GetDeliveryOptimizationMode()(*WindowsDeliveryOptimizationMode) {
    if m == nil {
        return nil
    } else {
        return m.deliveryOptimizationMode
    }
}
// GetDriversExcluded gets the driversExcluded property value. Exclude Windows update Drivers
func (m *WindowsUpdateForBusinessConfiguration) GetDriversExcluded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.driversExcluded
    }
}
// GetFeatureUpdatesDeferralPeriodInDays gets the featureUpdatesDeferralPeriodInDays property value. Defer Feature Updates by these many days
func (m *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesDeferralPeriodInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.featureUpdatesDeferralPeriodInDays
    }
}
// GetFeatureUpdatesPaused gets the featureUpdatesPaused property value. Pause Feature Updates
func (m *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesPaused()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.featureUpdatesPaused
    }
}
// GetFeatureUpdatesPauseExpiryDateTime gets the featureUpdatesPauseExpiryDateTime property value. Feature Updates Pause Expiry datetime
func (m *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesPauseExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.featureUpdatesPauseExpiryDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdateForBusinessConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["automaticUpdateMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAutomaticUpdateMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticUpdateMode(val.(*AutomaticUpdateMode))
        }
        return nil
    }
    res["businessReadyUpdatesOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsUpdateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessReadyUpdatesOnly(val.(*WindowsUpdateType))
        }
        return nil
    }
    res["deliveryOptimizationMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDeliveryOptimizationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryOptimizationMode(val.(*WindowsDeliveryOptimizationMode))
        }
        return nil
    }
    res["driversExcluded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriversExcluded(val)
        }
        return nil
    }
    res["featureUpdatesDeferralPeriodInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdatesDeferralPeriodInDays(val)
        }
        return nil
    }
    res["featureUpdatesPaused"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdatesPaused(val)
        }
        return nil
    }
    res["featureUpdatesPauseExpiryDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdatesPauseExpiryDateTime(val)
        }
        return nil
    }
    res["installationSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsUpdateInstallScheduleTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallationSchedule(val.(WindowsUpdateInstallScheduleTypeable))
        }
        return nil
    }
    res["microsoftUpdateServiceAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftUpdateServiceAllowed(val)
        }
        return nil
    }
    res["prereleaseFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrereleaseFeatures)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrereleaseFeatures(val.(*PrereleaseFeatures))
        }
        return nil
    }
    res["qualityUpdatesDeferralPeriodInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdatesDeferralPeriodInDays(val)
        }
        return nil
    }
    res["qualityUpdatesPaused"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdatesPaused(val)
        }
        return nil
    }
    res["qualityUpdatesPauseExpiryDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdatesPauseExpiryDateTime(val)
        }
        return nil
    }
    return res
}
// GetInstallationSchedule gets the installationSchedule property value. Installation schedule
func (m *WindowsUpdateForBusinessConfiguration) GetInstallationSchedule()(WindowsUpdateInstallScheduleTypeable) {
    if m == nil {
        return nil
    } else {
        return m.installationSchedule
    }
}
// GetMicrosoftUpdateServiceAllowed gets the microsoftUpdateServiceAllowed property value. Allow Microsoft Update Service
func (m *WindowsUpdateForBusinessConfiguration) GetMicrosoftUpdateServiceAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.microsoftUpdateServiceAllowed
    }
}
// GetPrereleaseFeatures gets the prereleaseFeatures property value. Possible values for pre-release features.
func (m *WindowsUpdateForBusinessConfiguration) GetPrereleaseFeatures()(*PrereleaseFeatures) {
    if m == nil {
        return nil
    } else {
        return m.prereleaseFeatures
    }
}
// GetQualityUpdatesDeferralPeriodInDays gets the qualityUpdatesDeferralPeriodInDays property value. Defer Quality Updates by these many days
func (m *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesDeferralPeriodInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.qualityUpdatesDeferralPeriodInDays
    }
}
// GetQualityUpdatesPaused gets the qualityUpdatesPaused property value. Pause Quality Updates
func (m *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesPaused()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.qualityUpdatesPaused
    }
}
// GetQualityUpdatesPauseExpiryDateTime gets the qualityUpdatesPauseExpiryDateTime property value. Quality Updates Pause Expiry datetime
func (m *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesPauseExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.qualityUpdatesPauseExpiryDateTime
    }
}
// Serialize serializes information the current object
func (m *WindowsUpdateForBusinessConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAutomaticUpdateMode() != nil {
        cast := (*m.GetAutomaticUpdateMode()).String()
        err = writer.WriteStringValue("automaticUpdateMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetBusinessReadyUpdatesOnly() != nil {
        cast := (*m.GetBusinessReadyUpdatesOnly()).String()
        err = writer.WriteStringValue("businessReadyUpdatesOnly", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeliveryOptimizationMode() != nil {
        cast := (*m.GetDeliveryOptimizationMode()).String()
        err = writer.WriteStringValue("deliveryOptimizationMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("driversExcluded", m.GetDriversExcluded())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("featureUpdatesDeferralPeriodInDays", m.GetFeatureUpdatesDeferralPeriodInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("featureUpdatesPaused", m.GetFeatureUpdatesPaused())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("featureUpdatesPauseExpiryDateTime", m.GetFeatureUpdatesPauseExpiryDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("installationSchedule", m.GetInstallationSchedule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftUpdateServiceAllowed", m.GetMicrosoftUpdateServiceAllowed())
        if err != nil {
            return err
        }
    }
    if m.GetPrereleaseFeatures() != nil {
        cast := (*m.GetPrereleaseFeatures()).String()
        err = writer.WriteStringValue("prereleaseFeatures", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("qualityUpdatesDeferralPeriodInDays", m.GetQualityUpdatesDeferralPeriodInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("qualityUpdatesPaused", m.GetQualityUpdatesPaused())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("qualityUpdatesPauseExpiryDateTime", m.GetQualityUpdatesPauseExpiryDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAutomaticUpdateMode sets the automaticUpdateMode property value. Possible values for automatic update mode.
func (m *WindowsUpdateForBusinessConfiguration) SetAutomaticUpdateMode(value *AutomaticUpdateMode)() {
    if m != nil {
        m.automaticUpdateMode = value
    }
}
// SetBusinessReadyUpdatesOnly sets the businessReadyUpdatesOnly property value. Which branch devices will receive their updates from
func (m *WindowsUpdateForBusinessConfiguration) SetBusinessReadyUpdatesOnly(value *WindowsUpdateType)() {
    if m != nil {
        m.businessReadyUpdatesOnly = value
    }
}
// SetDeliveryOptimizationMode sets the deliveryOptimizationMode property value. Delivery optimization mode for peer distribution
func (m *WindowsUpdateForBusinessConfiguration) SetDeliveryOptimizationMode(value *WindowsDeliveryOptimizationMode)() {
    if m != nil {
        m.deliveryOptimizationMode = value
    }
}
// SetDriversExcluded sets the driversExcluded property value. Exclude Windows update Drivers
func (m *WindowsUpdateForBusinessConfiguration) SetDriversExcluded(value *bool)() {
    if m != nil {
        m.driversExcluded = value
    }
}
// SetFeatureUpdatesDeferralPeriodInDays sets the featureUpdatesDeferralPeriodInDays property value. Defer Feature Updates by these many days
func (m *WindowsUpdateForBusinessConfiguration) SetFeatureUpdatesDeferralPeriodInDays(value *int32)() {
    if m != nil {
        m.featureUpdatesDeferralPeriodInDays = value
    }
}
// SetFeatureUpdatesPaused sets the featureUpdatesPaused property value. Pause Feature Updates
func (m *WindowsUpdateForBusinessConfiguration) SetFeatureUpdatesPaused(value *bool)() {
    if m != nil {
        m.featureUpdatesPaused = value
    }
}
// SetFeatureUpdatesPauseExpiryDateTime sets the featureUpdatesPauseExpiryDateTime property value. Feature Updates Pause Expiry datetime
func (m *WindowsUpdateForBusinessConfiguration) SetFeatureUpdatesPauseExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.featureUpdatesPauseExpiryDateTime = value
    }
}
// SetInstallationSchedule sets the installationSchedule property value. Installation schedule
func (m *WindowsUpdateForBusinessConfiguration) SetInstallationSchedule(value WindowsUpdateInstallScheduleTypeable)() {
    if m != nil {
        m.installationSchedule = value
    }
}
// SetMicrosoftUpdateServiceAllowed sets the microsoftUpdateServiceAllowed property value. Allow Microsoft Update Service
func (m *WindowsUpdateForBusinessConfiguration) SetMicrosoftUpdateServiceAllowed(value *bool)() {
    if m != nil {
        m.microsoftUpdateServiceAllowed = value
    }
}
// SetPrereleaseFeatures sets the prereleaseFeatures property value. Possible values for pre-release features.
func (m *WindowsUpdateForBusinessConfiguration) SetPrereleaseFeatures(value *PrereleaseFeatures)() {
    if m != nil {
        m.prereleaseFeatures = value
    }
}
// SetQualityUpdatesDeferralPeriodInDays sets the qualityUpdatesDeferralPeriodInDays property value. Defer Quality Updates by these many days
func (m *WindowsUpdateForBusinessConfiguration) SetQualityUpdatesDeferralPeriodInDays(value *int32)() {
    if m != nil {
        m.qualityUpdatesDeferralPeriodInDays = value
    }
}
// SetQualityUpdatesPaused sets the qualityUpdatesPaused property value. Pause Quality Updates
func (m *WindowsUpdateForBusinessConfiguration) SetQualityUpdatesPaused(value *bool)() {
    if m != nil {
        m.qualityUpdatesPaused = value
    }
}
// SetQualityUpdatesPauseExpiryDateTime sets the qualityUpdatesPauseExpiryDateTime property value. Quality Updates Pause Expiry datetime
func (m *WindowsUpdateForBusinessConfiguration) SetQualityUpdatesPauseExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.qualityUpdatesPauseExpiryDateTime = value
    }
}
