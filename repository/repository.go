package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
	"nwdaf-otel/clients/prometheus"
)

type Repository interface {
	Setup() error
	InsertBatch(data []prometheus.MetricResults) error
	Debug() error
}

type sqlLiteRepo struct {
	*sql.DB
}

func NewSQLiteRepo() (Repository, error) {
	db, err := sql.Open("sqlite", "/data/series.db")
	if err != nil {
		return nil, err
	}
	return sqlLiteRepo{
		db,
	}, nil
}

func (r sqlLiteRepo) Setup() error {
	_, err := r.Exec(`CREATE TABLE IF NOT EXISTS series (
        id INTEGER PRIMARY KEY,
        ts INTEGER,
        service TEXT,
        cpu_usage REAL,
        memory_usage REAL,
        total_bytes_sent REAL,
        total_packets_sent REAL,
        total_bytes_received REAL,
        total_packets_received REAL,
        avg_trace_duration REAL
    );`)
	if err != nil {
		return fmt.Errorf("error creating series table: %v", err)
	}
	return nil
}

func (r sqlLiteRepo) InsertBatch(metrics []prometheus.MetricResults) error {
	tx, err := r.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(
		"INSERT INTO series(ts, service, cpu_usage, memory_usage, total_bytes_sent, total_packets_sent, total_bytes_received, total_packets_received, avg_trace_duration) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer func() {
		err := stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	for _, m := range metrics {
		_, err := stmt.Exec(m.Timestamp, m.Service, m.CpuTotalSeconds, m.MemoryTotalBytes,
			m.NetworkTransmitBytesTotal, m.NetworkTransmitPacketsTotal,
			m.NetworkReceiveBytesTotal, m.NetworkReceivePacketsTotal,
			m.AvgTraceDuration)
		if err != nil {
			rollBackErr := tx.Rollback()
			if rollBackErr != nil {
				return errors.Join(err, rollBackErr)
			}
			return err
		}
	}
	return tx.Commit()
}

func (r sqlLiteRepo) Debug() error {
	rows, err := r.Query("SELECT * FROM series")
	if err != nil {
		return err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	for rows.Next() {
		m := prometheus.MetricResults{}
		err := rows.Scan(&m.Id, &m.Timestamp, &m.Service, &m.CpuTotalSeconds, &m.MemoryTotalBytes,
			&m.NetworkTransmitBytesTotal, &m.NetworkTransmitPacketsTotal,
			&m.NetworkReceiveBytesTotal, &m.NetworkReceivePacketsTotal,
			&m.AvgTraceDuration)
		if err != nil {
			return err
		}
		log.Println(m)
	}
	return nil
}
