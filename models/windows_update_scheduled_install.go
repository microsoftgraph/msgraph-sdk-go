package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdateScheduledInstall 
type WindowsUpdateScheduledInstall struct {
    WindowsUpdateInstallScheduleType
    // Scheduled Install Day in week. Possible values are: userDefined, everyday, sunday, monday, tuesday, wednesday, thursday, friday, saturday, noScheduledScan.
    scheduledInstallDay *WeeklySchedule
    // Scheduled Install Time during day
    scheduledInstallTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
}
// NewWindowsUpdateScheduledInstall instantiates a new WindowsUpdateScheduledInstall and sets the default values.
func NewWindowsUpdateScheduledInstall()(*WindowsUpdateScheduledInstall) {
    m := &WindowsUpdateScheduledInstall{
        WindowsUpdateInstallScheduleType: *NewWindowsUpdateInstallScheduleType(),
    }
    return m
}
// CreateWindowsUpdateScheduledInstallFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdateScheduledInstallFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdateScheduledInstall(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdateScheduledInstall) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsUpdateInstallScheduleType.GetFieldDeserializers()
    res["scheduledInstallDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWeeklySchedule)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledInstallDay(val.(*WeeklySchedule))
        }
        return nil
    }
    res["scheduledInstallTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledInstallTime(val)
        }
        return nil
    }
    return res
}
// GetScheduledInstallDay gets the scheduledInstallDay property value. Scheduled Install Day in week. Possible values are: userDefined, everyday, sunday, monday, tuesday, wednesday, thursday, friday, saturday, noScheduledScan.
func (m *WindowsUpdateScheduledInstall) GetScheduledInstallDay()(*WeeklySchedule) {
    if m == nil {
        return nil
    } else {
        return m.scheduledInstallDay
    }
}
// GetScheduledInstallTime gets the scheduledInstallTime property value. Scheduled Install Time during day
func (m *WindowsUpdateScheduledInstall) GetScheduledInstallTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.scheduledInstallTime
    }
}
// Serialize serializes information the current object
func (m *WindowsUpdateScheduledInstall) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsUpdateInstallScheduleType.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetScheduledInstallDay() != nil {
        cast := (*m.GetScheduledInstallDay()).String()
        err = writer.WriteStringValue("scheduledInstallDay", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("scheduledInstallTime", m.GetScheduledInstallTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetScheduledInstallDay sets the scheduledInstallDay property value. Scheduled Install Day in week. Possible values are: userDefined, everyday, sunday, monday, tuesday, wednesday, thursday, friday, saturday, noScheduledScan.
func (m *WindowsUpdateScheduledInstall) SetScheduledInstallDay(value *WeeklySchedule)() {
    if m != nil {
        m.scheduledInstallDay = value
    }
}
// SetScheduledInstallTime sets the scheduledInstallTime property value. Scheduled Install Time during day
func (m *WindowsUpdateScheduledInstall) SetScheduledInstallTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    if m != nil {
        m.scheduledInstallTime = value
    }
}
