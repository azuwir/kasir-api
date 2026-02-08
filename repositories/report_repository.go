package repositories

import (
	"database/sql"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetReport(start_date, end_date string) (map[string]interface{}, error) {
	query := "SELECT SUM(total_amount) AS total_revenue, COUNT(*) AS total_transactions FROM transactions"

	args := []interface{}{}
	if start_date != "" {
		query += " WHERE DATE(created_at) >= $1"
		args = append(args, start_date)
	}
	if end_date != "" {
		query += " AND DATE(created_at) <= $2"
		args = append(args, end_date)
	}
	query += " GROUP BY DATE(created_at)"

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	report := make(map[string]interface{})
	if rows.Next() {
		var total_revenue float64
		var total_transactions int
		err := rows.Scan(&total_revenue, &total_transactions)
		if err != nil {
			return nil, err
		}
		report["total_revenue"] = total_revenue
		report["total_transactions"] = total_transactions
		report["bestseller_products"] = make(map[string]interface{})
	} else {
		report["total_revenue"] = 0
		report["total_transactions"] = 0
		report["bestseller_products"] = make(map[string]interface{})
	}

	// Fetch bestseller products
	query_detail := "SELECT p.name, SUM(d.quantity) AS quantity_sold FROM transactions t JOIN transaction_details d ON t.id = d.transaction_id JOIN products p ON d.product_id = p.id"

	args = []interface{}{}
	if start_date != "" {
		query_detail += " WHERE DATE(t.created_at) >= $1"
		args = append(args, start_date)
	}
	if end_date != "" {
		query_detail += " AND DATE(t.created_at) <= $2"
		args = append(args, end_date)
	}
	query_detail += " GROUP BY p.name ORDER BY quantity_sold DESC LIMIT 1"

	rows, err = repo.db.Query(query_detail, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var name string
		var quantity_sold int
		err := rows.Scan(&name, &quantity_sold)
		if err != nil {
			return nil, err
		}

		bestsellers := report["bestseller_products"].(map[string]interface{})
		bestsellers["name"] = name
		bestsellers["quantity_sold"] = quantity_sold
	}

	return report, nil
}

func (repo *ReportRepository) GetReportToday(today string) (map[string]interface{}, error) {
	today = time.Now().Format("2006-01-02")
	query := "SELECT SUM(total_amount) AS total_revenue, COUNT(*) AS total_transactions FROM transactions WHERE DATE(created_at) = $1"

	rows, err := repo.db.Query(query, today)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	report := make(map[string]interface{})
	if rows.Next() {
		var total_revenue float64
		var total_transactions int
		err := rows.Scan(&total_revenue, &total_transactions)
		if err != nil {
			return nil, err
		}
		report["total_revenue"] = total_revenue
		report["total_transactions"] = total_transactions
		report["bestseller_products"] = make(map[string]interface{})
	} else {
		report["total_revenue"] = 0
		report["total_transactions"] = 0
		report["bestseller_products"] = make(map[string]interface{})
	}

	// Fetch bestseller products
	query_detail := "SELECT p.name, SUM(d.quantity) AS quantity_sold FROM transaction_details d JOIN transactions t ON d.transaction_id = t.id JOIN products p ON d.product_id = p.id"
	query_detail += " WHERE DATE(t.created_at) = $1 GROUP BY p.name ORDER BY quantity_sold DESC LIMIT 1"

	rows, err = repo.db.Query(query_detail, today)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var name string
		var quantity_sold int
		err := rows.Scan(&name, &quantity_sold)
		if err != nil {
			return nil, err
		}

		bestsellers := report["bestseller_products"].(map[string]interface{})
		bestsellers["name"] = name
		bestsellers["quantity_sold"] = quantity_sold
	}

	return report, nil
}
