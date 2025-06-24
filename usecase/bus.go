package usecase

import (
	"bus-booking/entity"
	"bus-booking/repository"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
)

type BusUsecase struct {
	Mysql *repository.MysqlCon
}

func (b *BusUsecase) CreateBus(req entity.Bus) error {
	if b.Mysql == nil {
		zap.L().Error("Database connection is nil")
		return errors.New("database connection not initialized")
	}

	// Convert facilities []string to JSON
	facilitiesJSON, err := json.Marshal(req.Facilities)
	if err != nil {
		zap.L().Error("Failed to marshal facilities", zap.Error(err))
		return err
	}
	departureTime := req.DepartureTime.Format("2006-01-02 15:04:05")
	arrivalTime := req.ArrivalTime.Format("2006-01-02 15:04:05")

	query := `
		INSERT INTO bus (
			bus_id, operator_name, bus_number, bus_type,
			total_seats, available_seats, facilities,
			from_location, to_location,
			departure_time, arrival_time,
			fare, status
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	err = b.Mysql.Connection.Exec(query,
		req.BusID,
		req.OperatorName,
		req.BusNumber,
		req.BusType,
		req.TotalSeats,
		req.AvailableSeats,
		facilitiesJSON,
		req.From,
		req.To,
		departureTime,
		arrivalTime,
		req.Fare,
		req.Status,
	).Error

	if err != nil {
		zap.L().Error("Failed to insert bus record", zap.Error(err))
		return err
	}

	zap.L().Info("Bus created successfully", zap.String("bus_id", req.BusID))
	return nil
}
