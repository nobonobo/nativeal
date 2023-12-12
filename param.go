package nativeal

type AudioParam interface {
	DefaultValue() float64
	MaxValue() float64
	MinValue() float64
	Value() float64
	SetValue(value float64)

	SetValueAtTime(value, startTime float64)
	LinearRampToValueAtTime(value, endTime float64)
	ExponentialRampToValueAtTime(value, endTime float64)
	SetTargetAtTime(target, startTime, timeConstant float64)
	SetValueCurveAtTime(values []float64, startTime, duration float64)
	CancelScheduledValues(startTime float64)
	CancelAndHoldAtTime(cancelTime float64)
}
