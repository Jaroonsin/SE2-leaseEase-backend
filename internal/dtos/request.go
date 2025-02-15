package dtos

type CreateRequest struct {
	Purpose            	string `json:"purpose" example:"Lease agreement for residential property"`
	ProposedMessage   	string `json:"proposedMessage" example:"I am very interested in leasing this property."`
	Question          	string `json:"question" example:"Could you provide more details about the lease duration?"`
	InterestedProperty 	uint   `json:"interestedProperty" example:"1"`
	LesseeID           	uint   `json:"lesseeID" example:"1"`
}

type UpdateRequest struct {
	Purpose         	string `json:"purpose" example:"Updated lease purpose"`
	ProposedMessage 	string `json:"proposedMessage" example:"I would like to update my earlier message."`
	Question        	string `json:"question" example:"Is there any flexibility in the lease terms?"`
}