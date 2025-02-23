package repositories

type LessorRepository interface {
	AcceptReservation(reservationID uint) error
	DeclineReservation(reservationID uint) error
}
