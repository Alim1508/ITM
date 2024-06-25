package mypackage

type producer interface {
    Produce() ([]string, error)
}

type presenter interface {
    Present([]string) error
}

type Service struct {
    prod producer
    pres presenter
}

func NewService(prod producer, pres presenter) *Service {
    return &Service{
        prod: prod,
        pres: pres,
    }
}

func (s *Service) MaskLinks(data []string) []string {
    maskedData := make([]string, len(data))
    for i, message := range data {
        maskedData[i] = string(maskLinks(message))
    }
    return maskedData
}

func (s *Service) Run() error {
    data, err := s.prod.Produce()
    if err != nil {
        return err
    }

    maskedData := s.MaskLinks(data)

    if err := s.pres.Present(maskedData); err != nil {
        return err
    }

    return nil
}
