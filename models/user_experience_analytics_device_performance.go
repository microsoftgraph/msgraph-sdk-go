package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDevicePerformance 
type UserExperienceAnalyticsDevicePerformance struct {
    Entity
    // Average (mean) number of Blue Screens per device in the last 14 days. Valid values 0 to 9999999
    averageBlueScreens *float64
    // Average (mean) number of Restarts per device in the last 14 days. Valid values 0 to 9999999
    averageRestarts *float64
    // Number of Blue Screens in the last 14 days. Valid values 0 to 9999999
    blueScreenCount *int32
    // The user experience analytics device boot score.
    bootScore *int32
    // The user experience analytics device core boot time in milliseconds.
    coreBootTimeInMs *int32
    // The user experience analytics device core login time in milliseconds.
    coreLoginTimeInMs *int32
    // User experience analytics summarized device count.
    deviceCount *int64
    // The user experience analytics device name.
    deviceName *string
    // The diskType property
    diskType *DiskType
    // The user experience analytics device group policy boot time in milliseconds.
    groupPolicyBootTimeInMs *int32
    // The user experience analytics device group policy login time in milliseconds.
    groupPolicyLoginTimeInMs *int32
    // The healthStatus property
    healthStatus *UserExperienceAnalyticsHealthState
    // The user experience analytics device login score.
    loginScore *int32
    // The user experience analytics device manufacturer.
    manufacturer *string
    // The user experience analytics device model.
    model *string
    // The user experience analytics model level startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    modelStartupPerformanceScore *float64
    // The user experience analytics device Operating System version.
    operatingSystemVersion *string
    // The user experience analytics responsive desktop time in milliseconds.
    responsiveDesktopTimeInMs *int32
    // Number of Restarts in the last 14 days. Valid values 0 to 9999999
    restartCount *int32
    // The user experience analytics device startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    startupPerformanceScore *float64
}
// NewUserExperienceAnalyticsDevicePerformance instantiates a new UserExperienceAnalyticsDevicePerformance and sets the default values.
func NewUserExperienceAnalyticsDevicePerformance()(*UserExperienceAnalyticsDevicePerformance) {
    m := &UserExperienceAnalyticsDevicePerformance{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.userExperienceAnalyticsDevicePerformance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserExperienceAnalyticsDevicePerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDevicePerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDevicePerformance(), nil
}
// GetAverageBlueScreens gets the averageBlueScreens property value. Average (mean) number of Blue Screens per device in the last 14 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) GetAverageBlueScreens()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageBlueScreens
    }
}
// GetAverageRestarts gets the averageRestarts property value. Average (mean) number of Restarts per device in the last 14 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) GetAverageRestarts()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageRestarts
    }
}
// GetBlueScreenCount gets the blueScreenCount property value. Number of Blue Screens in the last 14 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) GetBlueScreenCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.blueScreenCount
    }
}
// GetBootScore gets the bootScore property value. The user experience analytics device boot score.
func (m *UserExperienceAnalyticsDevicePerformance) GetBootScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.bootScore
    }
}
// GetCoreBootTimeInMs gets the coreBootTimeInMs property value. The user experience analytics device core boot time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) GetCoreBootTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.coreBootTimeInMs
    }
}
// GetCoreLoginTimeInMs gets the coreLoginTimeInMs property value. The user experience analytics device core login time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) GetCoreLoginTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.coreLoginTimeInMs
    }
}
// GetDeviceCount gets the deviceCount property value. User experience analytics summarized device count.
func (m *UserExperienceAnalyticsDevicePerformance) GetDeviceCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
// GetDeviceName gets the deviceName property value. The user experience analytics device name.
func (m *UserExperienceAnalyticsDevicePerformance) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetDiskType gets the diskType property value. The diskType property
func (m *UserExperienceAnalyticsDevicePerformance) GetDiskType()(*DiskType) {
    if m == nil {
        return nil
    } else {
        return m.diskType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDevicePerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["averageBlueScreens"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageBlueScreens(val)
        }
        return nil
    }
    res["averageRestarts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageRestarts(val)
        }
        return nil
    }
    res["blueScreenCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlueScreenCount(val)
        }
        return nil
    }
    res["bootScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootScore(val)
        }
        return nil
    }
    res["coreBootTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreBootTimeInMs(val)
        }
        return nil
    }
    res["coreLoginTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreLoginTimeInMs(val)
        }
        return nil
    }
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["diskType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiskType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiskType(val.(*DiskType))
        }
        return nil
    }
    res["groupPolicyBootTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyBootTimeInMs(val)
        }
        return nil
    }
    res["groupPolicyLoginTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyLoginTimeInMs(val)
        }
        return nil
    }
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*UserExperienceAnalyticsHealthState))
        }
        return nil
    }
    res["loginScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginScore(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["modelStartupPerformanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelStartupPerformanceScore(val)
        }
        return nil
    }
    res["operatingSystemVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemVersion(val)
        }
        return nil
    }
    res["responsiveDesktopTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponsiveDesktopTimeInMs(val)
        }
        return nil
    }
    res["restartCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartCount(val)
        }
        return nil
    }
    res["startupPerformanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupPerformanceScore(val)
        }
        return nil
    }
    return res
}
// GetGroupPolicyBootTimeInMs gets the groupPolicyBootTimeInMs property value. The user experience analytics device group policy boot time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) GetGroupPolicyBootTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyBootTimeInMs
    }
}
// GetGroupPolicyLoginTimeInMs gets the groupPolicyLoginTimeInMs property value. The user experience analytics device group policy login time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) GetGroupPolicyLoginTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyLoginTimeInMs
    }
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsDevicePerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    if m == nil {
        return nil
    } else {
        return m.healthStatus
    }
}
// GetLoginScore gets the loginScore property value. The user experience analytics device login score.
func (m *UserExperienceAnalyticsDevicePerformance) GetLoginScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.loginScore
    }
}
// GetManufacturer gets the manufacturer property value. The user experience analytics device manufacturer.
func (m *UserExperienceAnalyticsDevicePerformance) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetModel gets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsDevicePerformance) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetModelStartupPerformanceScore gets the modelStartupPerformanceScore property value. The user experience analytics model level startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDevicePerformance) GetModelStartupPerformanceScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.modelStartupPerformanceScore
    }
}
// GetOperatingSystemVersion gets the operatingSystemVersion property value. The user experience analytics device Operating System version.
func (m *UserExperienceAnalyticsDevicePerformance) GetOperatingSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemVersion
    }
}
// GetResponsiveDesktopTimeInMs gets the responsiveDesktopTimeInMs property value. The user experience analytics responsive desktop time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) GetResponsiveDesktopTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.responsiveDesktopTimeInMs
    }
}
// GetRestartCount gets the restartCount property value. Number of Restarts in the last 14 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) GetRestartCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.restartCount
    }
}
// GetStartupPerformanceScore gets the startupPerformanceScore property value. The user experience analytics device startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDevicePerformance) GetStartupPerformanceScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.startupPerformanceScore
    }
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDevicePerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("averageBlueScreens", m.GetAverageBlueScreens())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("averageRestarts", m.GetAverageRestarts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("blueScreenCount", m.GetBlueScreenCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("bootScore", m.GetBootScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("coreBootTimeInMs", m.GetCoreBootTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("coreLoginTimeInMs", m.GetCoreLoginTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("deviceCount", m.GetDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetDiskType() != nil {
        cast := (*m.GetDiskType()).String()
        err = writer.WriteStringValue("diskType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("groupPolicyBootTimeInMs", m.GetGroupPolicyBootTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("groupPolicyLoginTimeInMs", m.GetGroupPolicyLoginTimeInMs())
        if err != nil {
            return err
        }
    }
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("loginScore", m.GetLoginScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("modelStartupPerformanceScore", m.GetModelStartupPerformanceScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystemVersion", m.GetOperatingSystemVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("responsiveDesktopTimeInMs", m.GetResponsiveDesktopTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("restartCount", m.GetRestartCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("startupPerformanceScore", m.GetStartupPerformanceScore())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAverageBlueScreens sets the averageBlueScreens property value. Average (mean) number of Blue Screens per device in the last 14 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) SetAverageBlueScreens(value *float64)() {
    if m != nil {
        m.averageBlueScreens = value
    }
}
// SetAverageRestarts sets the averageRestarts property value. Average (mean) number of Restarts per device in the last 14 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) SetAverageRestarts(value *float64)() {
    if m != nil {
        m.averageRestarts = value
    }
}
// SetBlueScreenCount sets the blueScreenCount property value. Number of Blue Screens in the last 14 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) SetBlueScreenCount(value *int32)() {
    if m != nil {
        m.blueScreenCount = value
    }
}
// SetBootScore sets the bootScore property value. The user experience analytics device boot score.
func (m *UserExperienceAnalyticsDevicePerformance) SetBootScore(value *int32)() {
    if m != nil {
        m.bootScore = value
    }
}
// SetCoreBootTimeInMs sets the coreBootTimeInMs property value. The user experience analytics device core boot time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) SetCoreBootTimeInMs(value *int32)() {
    if m != nil {
        m.coreBootTimeInMs = value
    }
}
// SetCoreLoginTimeInMs sets the coreLoginTimeInMs property value. The user experience analytics device core login time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) SetCoreLoginTimeInMs(value *int32)() {
    if m != nil {
        m.coreLoginTimeInMs = value
    }
}
// SetDeviceCount sets the deviceCount property value. User experience analytics summarized device count.
func (m *UserExperienceAnalyticsDevicePerformance) SetDeviceCount(value *int64)() {
    if m != nil {
        m.deviceCount = value
    }
}
// SetDeviceName sets the deviceName property value. The user experience analytics device name.
func (m *UserExperienceAnalyticsDevicePerformance) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetDiskType sets the diskType property value. The diskType property
func (m *UserExperienceAnalyticsDevicePerformance) SetDiskType(value *DiskType)() {
    if m != nil {
        m.diskType = value
    }
}
// SetGroupPolicyBootTimeInMs sets the groupPolicyBootTimeInMs property value. The user experience analytics device group policy boot time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) SetGroupPolicyBootTimeInMs(value *int32)() {
    if m != nil {
        m.groupPolicyBootTimeInMs = value
    }
}
// SetGroupPolicyLoginTimeInMs sets the groupPolicyLoginTimeInMs property value. The user experience analytics device group policy login time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) SetGroupPolicyLoginTimeInMs(value *int32)() {
    if m != nil {
        m.groupPolicyLoginTimeInMs = value
    }
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsDevicePerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    if m != nil {
        m.healthStatus = value
    }
}
// SetLoginScore sets the loginScore property value. The user experience analytics device login score.
func (m *UserExperienceAnalyticsDevicePerformance) SetLoginScore(value *int32)() {
    if m != nil {
        m.loginScore = value
    }
}
// SetManufacturer sets the manufacturer property value. The user experience analytics device manufacturer.
func (m *UserExperienceAnalyticsDevicePerformance) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetModel sets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsDevicePerformance) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetModelStartupPerformanceScore sets the modelStartupPerformanceScore property value. The user experience analytics model level startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDevicePerformance) SetModelStartupPerformanceScore(value *float64)() {
    if m != nil {
        m.modelStartupPerformanceScore = value
    }
}
// SetOperatingSystemVersion sets the operatingSystemVersion property value. The user experience analytics device Operating System version.
func (m *UserExperienceAnalyticsDevicePerformance) SetOperatingSystemVersion(value *string)() {
    if m != nil {
        m.operatingSystemVersion = value
    }
}
// SetResponsiveDesktopTimeInMs sets the responsiveDesktopTimeInMs property value. The user experience analytics responsive desktop time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) SetResponsiveDesktopTimeInMs(value *int32)() {
    if m != nil {
        m.responsiveDesktopTimeInMs = value
    }
}
// SetRestartCount sets the restartCount property value. Number of Restarts in the last 14 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) SetRestartCount(value *int32)() {
    if m != nil {
        m.restartCount = value
    }
}
// SetStartupPerformanceScore sets the startupPerformanceScore property value. The user experience analytics device startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDevicePerformance) SetStartupPerformanceScore(value *float64)() {
    if m != nil {
        m.startupPerformanceScore = value
    }
}
