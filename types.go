package fnboot

type FnT interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~string | ~float32 | ~float64 | ~bool
}

type FnInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type FnString interface {
	~string
}

type FnFloat interface {
	~float32 | ~float64
}

type FnBool interface {
	~bool
}

type FnNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type FnInterface interface {
	interface{}
}

type FnAny interface {
	any
}
