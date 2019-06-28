package di

// DI struct for DI encapsulation the di components
type DI struct {
	// add di component types here
	Sample func() string
}

// InitDI - initilalize dependency injection struct
func InitDI() DI {
	var di DI

	// add di components here
	di.Sample = SampleDI
	return di
}
