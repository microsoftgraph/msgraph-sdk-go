package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Schedule provides operations to manage the drive singleton.
type Schedule struct {
    Entity
    // Indicates whether the schedule is enabled for the team. Required.
    enabled *bool;
    // 
    offerShiftRequests []OfferShiftRequestable;
    // Indicates whether offer shift requests are enabled for the schedule.
    offerShiftRequestsEnabled *bool;
    // 
    openShiftChangeRequests []OpenShiftChangeRequestable;
    // 
    openShifts []OpenShiftable;
    // Indicates whether open shifts are enabled for the schedule.
    openShiftsEnabled *bool;
    // The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
    provisionStatus *OperationStatus;
    // Additional information about why schedule provisioning failed.
    provisionStatusCode *string;
    // The logical grouping of users in the schedule (usually by role).
    schedulingGroups []SchedulingGroupable;
    // The shifts in the schedule.
    shifts []Shiftable;
    // 
    swapShiftsChangeRequests []SwapShiftsChangeRequestable;
    // Indicates whether swap shifts requests are enabled for the schedule.
    swapShiftsRequestsEnabled *bool;
    // Indicates whether time clock is enabled for the schedule.
    timeClockEnabled *bool;
    // The set of reasons for a time off in the schedule.
    timeOffReasons []TimeOffReasonable;
    // 
    timeOffRequests []TimeOffRequestable;
    // Indicates whether time off requests are enabled for the schedule.
    timeOffRequestsEnabled *bool;
    // The instances of times off in the schedule.
    timesOff []TimeOffable;
    // Indicates the time zone of the schedule team using tz database format. Required.
    timeZone *string;
    // 
    workforceIntegrationIds []string;
}
// NewSchedule instantiates a new schedule and sets the default values.
func NewSchedule()(*Schedule) {
    m := &Schedule{
        Entity: *NewEntity(),
    }
    return m
}
// CreateScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateScheduleFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSchedule(), nil
}
// GetEnabled gets the enabled property value. Indicates whether the schedule is enabled for the team. Required.
func (m *Schedule) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Schedule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["offerShiftRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOfferShiftRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfferShiftRequestable, len(val))
            for i, v := range val {
                res[i] = v.(OfferShiftRequestable)
            }
            m.SetOfferShiftRequests(res)
        }
        return nil
    }
    res["offerShiftRequestsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfferShiftRequestsEnabled(val)
        }
        return nil
    }
    res["openShiftChangeRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOpenShiftChangeRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OpenShiftChangeRequestable, len(val))
            for i, v := range val {
                res[i] = v.(OpenShiftChangeRequestable)
            }
            m.SetOpenShiftChangeRequests(res)
        }
        return nil
    }
    res["openShifts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOpenShiftFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OpenShiftable, len(val))
            for i, v := range val {
                res[i] = v.(OpenShiftable)
            }
            m.SetOpenShifts(res)
        }
        return nil
    }
    res["openShiftsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenShiftsEnabled(val)
        }
        return nil
    }
    res["provisionStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisionStatus(val.(*OperationStatus))
        }
        return nil
    }
    res["provisionStatusCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisionStatusCode(val)
        }
        return nil
    }
    res["schedulingGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSchedulingGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SchedulingGroupable, len(val))
            for i, v := range val {
                res[i] = v.(SchedulingGroupable)
            }
            m.SetSchedulingGroups(res)
        }
        return nil
    }
    res["shifts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateShiftFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Shiftable, len(val))
            for i, v := range val {
                res[i] = v.(Shiftable)
            }
            m.SetShifts(res)
        }
        return nil
    }
    res["swapShiftsChangeRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSwapShiftsChangeRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SwapShiftsChangeRequestable, len(val))
            for i, v := range val {
                res[i] = v.(SwapShiftsChangeRequestable)
            }
            m.SetSwapShiftsChangeRequests(res)
        }
        return nil
    }
    res["swapShiftsRequestsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSwapShiftsRequestsEnabled(val)
        }
        return nil
    }
    res["timeClockEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeClockEnabled(val)
        }
        return nil
    }
    res["timeOffReasons"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeOffReasonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeOffReasonable, len(val))
            for i, v := range val {
                res[i] = v.(TimeOffReasonable)
            }
            m.SetTimeOffReasons(res)
        }
        return nil
    }
    res["timeOffRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeOffRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeOffRequestable, len(val))
            for i, v := range val {
                res[i] = v.(TimeOffRequestable)
            }
            m.SetTimeOffRequests(res)
        }
        return nil
    }
    res["timeOffRequestsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeOffRequestsEnabled(val)
        }
        return nil
    }
    res["timesOff"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeOffFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeOffable, len(val))
            for i, v := range val {
                res[i] = v.(TimeOffable)
            }
            m.SetTimesOff(res)
        }
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    res["workforceIntegrationIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetWorkforceIntegrationIds(res)
        }
        return nil
    }
    return res
}
// GetOfferShiftRequests gets the offerShiftRequests property value. 
func (m *Schedule) GetOfferShiftRequests()([]OfferShiftRequestable) {
    if m == nil {
        return nil
    } else {
        return m.offerShiftRequests
    }
}
// GetOfferShiftRequestsEnabled gets the offerShiftRequestsEnabled property value. Indicates whether offer shift requests are enabled for the schedule.
func (m *Schedule) GetOfferShiftRequestsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.offerShiftRequestsEnabled
    }
}
// GetOpenShiftChangeRequests gets the openShiftChangeRequests property value. 
func (m *Schedule) GetOpenShiftChangeRequests()([]OpenShiftChangeRequestable) {
    if m == nil {
        return nil
    } else {
        return m.openShiftChangeRequests
    }
}
// GetOpenShifts gets the openShifts property value. 
func (m *Schedule) GetOpenShifts()([]OpenShiftable) {
    if m == nil {
        return nil
    } else {
        return m.openShifts
    }
}
// GetOpenShiftsEnabled gets the openShiftsEnabled property value. Indicates whether open shifts are enabled for the schedule.
func (m *Schedule) GetOpenShiftsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.openShiftsEnabled
    }
}
// GetProvisionStatus gets the provisionStatus property value. The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
func (m *Schedule) GetProvisionStatus()(*OperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.provisionStatus
    }
}
// GetProvisionStatusCode gets the provisionStatusCode property value. Additional information about why schedule provisioning failed.
func (m *Schedule) GetProvisionStatusCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisionStatusCode
    }
}
// GetSchedulingGroups gets the schedulingGroups property value. The logical grouping of users in the schedule (usually by role).
func (m *Schedule) GetSchedulingGroups()([]SchedulingGroupable) {
    if m == nil {
        return nil
    } else {
        return m.schedulingGroups
    }
}
// GetShifts gets the shifts property value. The shifts in the schedule.
func (m *Schedule) GetShifts()([]Shiftable) {
    if m == nil {
        return nil
    } else {
        return m.shifts
    }
}
// GetSwapShiftsChangeRequests gets the swapShiftsChangeRequests property value. 
func (m *Schedule) GetSwapShiftsChangeRequests()([]SwapShiftsChangeRequestable) {
    if m == nil {
        return nil
    } else {
        return m.swapShiftsChangeRequests
    }
}
// GetSwapShiftsRequestsEnabled gets the swapShiftsRequestsEnabled property value. Indicates whether swap shifts requests are enabled for the schedule.
func (m *Schedule) GetSwapShiftsRequestsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.swapShiftsRequestsEnabled
    }
}
// GetTimeClockEnabled gets the timeClockEnabled property value. Indicates whether time clock is enabled for the schedule.
func (m *Schedule) GetTimeClockEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.timeClockEnabled
    }
}
// GetTimeOffReasons gets the timeOffReasons property value. The set of reasons for a time off in the schedule.
func (m *Schedule) GetTimeOffReasons()([]TimeOffReasonable) {
    if m == nil {
        return nil
    } else {
        return m.timeOffReasons
    }
}
// GetTimeOffRequests gets the timeOffRequests property value. 
func (m *Schedule) GetTimeOffRequests()([]TimeOffRequestable) {
    if m == nil {
        return nil
    } else {
        return m.timeOffRequests
    }
}
// GetTimeOffRequestsEnabled gets the timeOffRequestsEnabled property value. Indicates whether time off requests are enabled for the schedule.
func (m *Schedule) GetTimeOffRequestsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.timeOffRequestsEnabled
    }
}
// GetTimesOff gets the timesOff property value. The instances of times off in the schedule.
func (m *Schedule) GetTimesOff()([]TimeOffable) {
    if m == nil {
        return nil
    } else {
        return m.timesOff
    }
}
// GetTimeZone gets the timeZone property value. Indicates the time zone of the schedule team using tz database format. Required.
func (m *Schedule) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
// GetWorkforceIntegrationIds gets the workforceIntegrationIds property value. 
func (m *Schedule) GetWorkforceIntegrationIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.workforceIntegrationIds
    }
}
func (m *Schedule) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Schedule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetOfferShiftRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOfferShiftRequests()))
        for i, v := range m.GetOfferShiftRequests() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("offerShiftRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("offerShiftRequestsEnabled", m.GetOfferShiftRequestsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetOpenShiftChangeRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOpenShiftChangeRequests()))
        for i, v := range m.GetOpenShiftChangeRequests() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("openShiftChangeRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOpenShifts() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOpenShifts()))
        for i, v := range m.GetOpenShifts() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("openShifts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("openShiftsEnabled", m.GetOpenShiftsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetProvisionStatus() != nil {
        cast := (*m.GetProvisionStatus()).String()
        err = writer.WriteStringValue("provisionStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("provisionStatusCode", m.GetProvisionStatusCode())
        if err != nil {
            return err
        }
    }
    if m.GetSchedulingGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSchedulingGroups()))
        for i, v := range m.GetSchedulingGroups() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("schedulingGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetShifts() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetShifts()))
        for i, v := range m.GetShifts() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("shifts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSwapShiftsChangeRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSwapShiftsChangeRequests()))
        for i, v := range m.GetSwapShiftsChangeRequests() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("swapShiftsChangeRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("swapShiftsRequestsEnabled", m.GetSwapShiftsRequestsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("timeClockEnabled", m.GetTimeClockEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetTimeOffReasons() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTimeOffReasons()))
        for i, v := range m.GetTimeOffReasons() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("timeOffReasons", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTimeOffRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTimeOffRequests()))
        for i, v := range m.GetTimeOffRequests() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("timeOffRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("timeOffRequestsEnabled", m.GetTimeOffRequestsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetTimesOff() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTimesOff()))
        for i, v := range m.GetTimesOff() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("timesOff", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    if m.GetWorkforceIntegrationIds() != nil {
        err = writer.WriteCollectionOfStringValues("workforceIntegrationIds", m.GetWorkforceIntegrationIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnabled sets the enabled property value. Indicates whether the schedule is enabled for the team. Required.
func (m *Schedule) SetEnabled(value *bool)() {
    if m != nil {
        m.enabled = value
    }
}
// SetOfferShiftRequests sets the offerShiftRequests property value. 
func (m *Schedule) SetOfferShiftRequests(value []OfferShiftRequestable)() {
    if m != nil {
        m.offerShiftRequests = value
    }
}
// SetOfferShiftRequestsEnabled sets the offerShiftRequestsEnabled property value. Indicates whether offer shift requests are enabled for the schedule.
func (m *Schedule) SetOfferShiftRequestsEnabled(value *bool)() {
    if m != nil {
        m.offerShiftRequestsEnabled = value
    }
}
// SetOpenShiftChangeRequests sets the openShiftChangeRequests property value. 
func (m *Schedule) SetOpenShiftChangeRequests(value []OpenShiftChangeRequestable)() {
    if m != nil {
        m.openShiftChangeRequests = value
    }
}
// SetOpenShifts sets the openShifts property value. 
func (m *Schedule) SetOpenShifts(value []OpenShiftable)() {
    if m != nil {
        m.openShifts = value
    }
}
// SetOpenShiftsEnabled sets the openShiftsEnabled property value. Indicates whether open shifts are enabled for the schedule.
func (m *Schedule) SetOpenShiftsEnabled(value *bool)() {
    if m != nil {
        m.openShiftsEnabled = value
    }
}
// SetProvisionStatus sets the provisionStatus property value. The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
func (m *Schedule) SetProvisionStatus(value *OperationStatus)() {
    if m != nil {
        m.provisionStatus = value
    }
}
// SetProvisionStatusCode sets the provisionStatusCode property value. Additional information about why schedule provisioning failed.
func (m *Schedule) SetProvisionStatusCode(value *string)() {
    if m != nil {
        m.provisionStatusCode = value
    }
}
// SetSchedulingGroups sets the schedulingGroups property value. The logical grouping of users in the schedule (usually by role).
func (m *Schedule) SetSchedulingGroups(value []SchedulingGroupable)() {
    if m != nil {
        m.schedulingGroups = value
    }
}
// SetShifts sets the shifts property value. The shifts in the schedule.
func (m *Schedule) SetShifts(value []Shiftable)() {
    if m != nil {
        m.shifts = value
    }
}
// SetSwapShiftsChangeRequests sets the swapShiftsChangeRequests property value. 
func (m *Schedule) SetSwapShiftsChangeRequests(value []SwapShiftsChangeRequestable)() {
    if m != nil {
        m.swapShiftsChangeRequests = value
    }
}
// SetSwapShiftsRequestsEnabled sets the swapShiftsRequestsEnabled property value. Indicates whether swap shifts requests are enabled for the schedule.
func (m *Schedule) SetSwapShiftsRequestsEnabled(value *bool)() {
    if m != nil {
        m.swapShiftsRequestsEnabled = value
    }
}
// SetTimeClockEnabled sets the timeClockEnabled property value. Indicates whether time clock is enabled for the schedule.
func (m *Schedule) SetTimeClockEnabled(value *bool)() {
    if m != nil {
        m.timeClockEnabled = value
    }
}
// SetTimeOffReasons sets the timeOffReasons property value. The set of reasons for a time off in the schedule.
func (m *Schedule) SetTimeOffReasons(value []TimeOffReasonable)() {
    if m != nil {
        m.timeOffReasons = value
    }
}
// SetTimeOffRequests sets the timeOffRequests property value. 
func (m *Schedule) SetTimeOffRequests(value []TimeOffRequestable)() {
    if m != nil {
        m.timeOffRequests = value
    }
}
// SetTimeOffRequestsEnabled sets the timeOffRequestsEnabled property value. Indicates whether time off requests are enabled for the schedule.
func (m *Schedule) SetTimeOffRequestsEnabled(value *bool)() {
    if m != nil {
        m.timeOffRequestsEnabled = value
    }
}
// SetTimesOff sets the timesOff property value. The instances of times off in the schedule.
func (m *Schedule) SetTimesOff(value []TimeOffable)() {
    if m != nil {
        m.timesOff = value
    }
}
// SetTimeZone sets the timeZone property value. Indicates the time zone of the schedule team using tz database format. Required.
func (m *Schedule) SetTimeZone(value *string)() {
    if m != nil {
        m.timeZone = value
    }
}
// SetWorkforceIntegrationIds sets the workforceIntegrationIds property value. 
func (m *Schedule) SetWorkforceIntegrationIds(value []string)() {
    if m != nil {
        m.workforceIntegrationIds = value
    }
}
