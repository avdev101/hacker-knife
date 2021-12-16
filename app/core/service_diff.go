package core

type serviceDiff struct {
	existing  []Service
	collected []ServiceCollectItem
}

func (s *serviceDiff) getNew() []ServiceCollectItem {
	result := make([]ServiceCollectItem, 0)
	existingMap := s.getExistingMap()

	for _, found := range s.collected {
		_, ok := existingMap[found.ServiceType]
		if !ok {
			result = append(result, found)
		}
	}

	return result
}

func (s *serviceDiff) getDeleted() []Service {
	result := make([]Service, 0)
	collectedMap := s.getCollectedMap()

	for _, service := range s.existing {
		_, ok := collectedMap[service.ServiceType]
		if !ok {
			result = append(result, service)
		}
	}

	return result
}

func (s *serviceDiff) getExistingMap() map[string]Service {
	result := make(map[string]Service)

	for _, s := range s.existing {
		result[s.ServiceType] = s
	}

	return result
}

func (s *serviceDiff) getCollectedMap() map[string]ServiceCollectItem {
	result := make(map[string]ServiceCollectItem)

	for _, found := range s.collected {
		result[found.ServiceType] = found
	}

	return result
}
