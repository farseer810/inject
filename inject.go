package inject

var defaultContainer = New()

func Provide(constructor interface{}, opts ...ProvideOption) error {
	return defaultContainer.Provide(constructor, opts...)
}

func ProvideValue(value interface{}, opts ...ProvideOption) error {
	return defaultContainer.ProvideWithValue(value, opts...)
}

func ProvideValues(values ...interface{}) error {
	return defaultContainer.ProvideWithValues(values...)
}

func Invoke(function interface{}, opts ...InvokeOption) error {
	return defaultContainer.Invoke(function, opts...)
}

func GetValues(pointers ...interface{}) error {
	return defaultContainer.GetValues(pointers...)
}

func Populate() error {
	return defaultContainer.Populate()
}

func PopulateValues(pointers ...interface{}) error {
	oldRequiredValuePointers := defaultContainer.requiredValuePointers
	err := defaultContainer.GetValues(pointers...)
	if err != nil {
		return err
	}
	err = defaultContainer.Populate()
	if err != nil {
		defaultContainer.requiredValuePointers = oldRequiredValuePointers
		return err
	}
	return nil
}
