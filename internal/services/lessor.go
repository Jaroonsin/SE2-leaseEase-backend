package services

type LessorService interface {
	AcceptReservation(reservationID uint) error
	DeclineReservation(reservationID uint) error
}
