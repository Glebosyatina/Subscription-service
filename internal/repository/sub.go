package repository

import (
	"database/sql"
	"glebosyatina/test_project/internal/domain"
)

type SubRepo struct {
	db *sql.DB
}

func NewSubRepo(dbconn *sql.DB) *SubRepo {
	return &SubRepo{
		db: dbconn,
	}
}

func (r *SubRepo) CreateSub(userId uint64, nameService string, price uint64, start string, end string) (*domain.Sub, error) {
	var subId uint64
	err := r.db.QueryRow("INSERT INTO subscriptions (user_id, service_name, price, start_date, end_date) VALUES ($1,$2,$3,$4,$5) RETURNING id",
		userId, nameService, price, start, end).Scan(&subId)
	if err != nil {
		return nil, err
	}

	return &domain.Sub{
		Id:          subId,
		UserId:      userId,
		NameService: nameService,
		Price:       price,
		Start:       start,
		End:         end,
	}, nil
}

func (r *SubRepo) GetSubByID(idSub uint64) (*domain.Sub, error) {
	var sub domain.Sub

	err := r.db.QueryRow("SELECT * FROM subscriptions WHERE id = $1", idSub).Scan(&sub.Id, &sub.UserId, &sub.NameService, &sub.Price, &sub.Start, &sub.End)
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (r *SubRepo) GetSubscriptions() ([]*domain.Sub, error) {
	subs := make([]*domain.Sub, 0)

	rows, err := r.db.Query("SELECT * FROM subscriptions")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s domain.Sub
		err := rows.Scan(&s.Id, &s.UserId, &s.NameService, &s.Price, &s.Start, &s.End)
		if err != nil {
			return nil, err
		}
		subs = append(subs, &s)
	}
	return subs, nil
}

func (r *SubRepo) DeleteSubByID(idSub uint64) error {
	if _, err := r.db.Exec("DELETE FROM subscriptions WHERE id=$1 RETURNING id", idSub); err != nil {
		return err
	}
	return nil
}

func (r *SubRepo) UpdateSub(idSub uint64, userId uint64, nameService string, price uint64, start string, end string) (*domain.Sub, error) {
	var sub domain.Sub
	err := r.db.QueryRow("UPDATE subscriptions SET user_id=$1, service_name=$2, price=$3, start_date=$4, end_date=$5 WHERE id=$6 RETURNING id, user_id, service_name, price, start_date, end_date", userId, nameService, price, start, end, idSub).Scan(
		&sub.Id, &sub.UserId, &sub.NameService, &sub.Price, &sub.Start, &sub.End)
	if err != nil {
		return nil, err
	}
	return &sub, nil
}
