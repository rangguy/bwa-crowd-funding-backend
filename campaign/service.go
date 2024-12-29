package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repo.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repo.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
