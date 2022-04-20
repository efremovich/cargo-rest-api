package seeds

// Init will seeds initial seeder.
func Init() []Seed {
	return initFactory()
}

// All will seeds all defined seeder.
func All() []Seed {
	b := prepare()
	return b
}

// prepare will prepare fake data based on entity's faker struct.
func prepare() []Seed {
	roleFactories := roleFactory()
	userFactories := userFactory()
	sityFactories := sityFactory()
	vehicleFactories := vehicleFactory()
	passengerTypeFactories := passengerTypeFactory()
	priceFactories := priceFactory()
	passengerFactories := passengerFactory()

	var (
		allFactories []Seed
	)
	allFactories = append(allFactories, roleFactories...)
	allFactories = append(allFactories, userFactories...)
	allFactories = append(allFactories, sityFactories...)
	allFactories = append(allFactories, vehicleFactories...)
	allFactories = append(allFactories, passengerTypeFactories...)
	allFactories = append(allFactories, priceFactories...)
	allFactories = append(allFactories, passengerFactories...)

	return allFactories
}
