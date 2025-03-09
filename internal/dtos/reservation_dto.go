package dtos

type CreateReservationDTO struct {
	Purpose            string `json:"purpose" example:"Lease agreement for residential property"`
	ProposedMessage    string `json:"proposedMessage" example:"I am very interested in leasing this property."`
	Question           string `json:"question" example:"Could you provide more details about the lease duration?"`
	InterestedProperty uint   `json:"interestedProperty" example:"1"`
}

type ReservationResponseDTO struct {
	ID uint `json:"reservation_id" example:"1"`
}

type UpdateReservationDTO struct {
	Purpose         string `json:"purpose" example:"Updated lease purpose"`
	ProposedMessage string `json:"proposedMessage" example:"I would like to update my earlier message."`
	Question        string `json:"question" example:"Is there any flexibility in the lease terms?"`
}

type ApprovalReservationDTO struct {
	LesseeEmail  string `json:"lesseeEmail" example:"lessee@example.com"`
	PropertyName string `json:"propertyName" example:"Example Property"`
}

type GetReservationDTO struct {
	ID              uint   `json:"id" example:"1"`
	Purpose         string `json:"purpose" example:"Lease agreement for residential property"`
	ProposedMessage string `json:"proposedMessage" example:"I am very interested in leasing this property."`
	Question        string `json:"question" example:"Could you provide more details about the lease duration?"`
	Status          string `json:"status" example:"pending"`
	PropertyID      uint   `json:"interestedProperty" example:"1"`
	LesseeID        uint   `json:"lesseeID" example:"1"`
	PropertyName    string `json:"propertyName" example:"Example Property"`
	LastModified    string `json:"lastModified" example:"2022-01-01T00:00:00Z"`
}
