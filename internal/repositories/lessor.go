package repositories

type LessorRepository interface {
	AcceptReservation(reservationID uint, lessorID uint) error
	DeclineReservation(reservationID uint, lessorID uint) error
}
