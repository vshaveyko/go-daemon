package process

func Entrypoint() {

	// toProcess := getRequiredProcessing()
	toProcess := []Pipeline{
		{
			Connectors: []Connector{
				{
					Source: EndPoint{Name: "First-connector", Id: 1},
					Target: EndPoint{Name: "Second-connector", Id: 2},
				},
				{
					Source: EndPoint{Name: "Third-connector", Id: 3},
					Target: EndPoint{Name: "Second-connector", Id: 2},
				},
				{
					Source: EndPoint{Name: "Fourth-connector", Id: 4},
					Target: EndPoint{Name: "First-connector", Id: 1},
				},
			},
		},
	}

	scheduleProcessee(toProcess)

}
