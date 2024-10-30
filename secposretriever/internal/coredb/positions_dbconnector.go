package coredb

import (
	"context"
	"fmt"
	"log"
	"time"
)

// credit : https://stackoverflow.com/questions/28751402/how-can-i-share-database-connection-between-packages-in-go
//var DBposition *db.DBstruct

type TFSXSP9DDL struct {
	key_part_date  int
	party_bic      string
	sec_acc_id     int
	sec_id         int
	ts_creat_row   string
	busn_dt        time.Time
	isin           string
	pos_qty        int64
	pos_qty_fd     string
	sett_typ       string
	neg_pos_flag   string
	period_evt_ref string
	syst_ent       string
	sec_acc_ext    string
	sec_acc_typ    string
	// cd_end_invstr        string
	// party_id             int
	//	rstr_typ_id    int
	// cntry_cod            string
	// cfi                  string
	// final_matur_expir_dt time.Time
	// issu_dt              time.Time
	// min_sett_unit        int
	// min_sett_unit_fd     int
	// min_sett_amt         int32
	// sett_typ4            string
	// sett_unit_mult       int
	// curncy_cod           string
	// curncy_amt_fd        int
}

func (r RequestPosition) GetPositionsFromDB() (data []Position, err error) {

	var selectClauseSql = "select a.key_part_date,a.syst_ent,a.party_bic,a.sec_acc_ext,a.isin,a.sec_acc_typ,a.sec_acc_id,a.sec_id,a.sec_acc_typ,a.ts_creat_row,a.busn_dt,a.pos_qty,a.pos_qty_fd,a.sett_typ,a.neg_pos_flag,a.period_evt_ref "
	var fromClauseSql = r.GetSQLbeforeFilter()
	var dynamicSql = selectClauseSql + fromClauseSql

	if r.Bic != "" {
		dynamicSql = dynamicSql + ` a.party_bic = '` + r.Bic + `' and `
	}

	if r.Isin != "" {
		dynamicSql = dynamicSql + ` a.isin = '` + r.Isin + `'`
	}

	if r.Account != "" {
		dynamicSql = dynamicSql + ` and a.sec_acc_ext = '` + r.Account + `'`
	}

	if r.Restrictiontype != "" {
		dynamicSql = dynamicSql + ` and a.sec_acc_typ = '` + r.Restrictiontype + `'`
	}

	if r.Filter.Date != "" {
		dynamicSql = dynamicSql + ` and a.busn_dt = '` + r.Filter.Date + `'`
	}

	if r.Filter.Phase != "" && r.Filter.Phase != "LAST" {
		dynamicSql = dynamicSql + ` and a.period_evt_ref = '` + r.Filter.Phase + `'`
	}

	fmt.Printf("\n<debut>" + dynamicSql + "<fin>\n")

	rows, err := DBconnector.ConPool.Query(context.Background(), dynamicSql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var onePosition Position

	//	var rowSlice []TFSXSP9DDL
	for rows.Next() {
		var r TFSXSP9DDL
		err := rows.Scan(
			&r.key_part_date,
			&r.syst_ent,
			&r.party_bic,
			&r.sec_acc_ext,
			&r.isin,
			&r.sec_acc_typ,
			&r.sec_acc_id,
			&r.sec_id,
			&r.sec_acc_typ,
			&r.ts_creat_row,
			&r.busn_dt,
			&r.pos_qty,
			&r.pos_qty_fd,
			&r.sett_typ,
			&r.neg_pos_flag,
			&r.period_evt_ref)

		if err != nil {
			log.Fatal(err)
		}

		onePosition.Isin = r.isin
		onePosition.Account = r.sec_acc_ext
		onePosition.Restrictiontype = r.sec_acc_typ
		onePosition.Quantity = r.pos_qty
		onePosition.QuantityFD = r.pos_qty_fd
		onePosition.LastTimestamp = r.ts_creat_row
		onePosition.Phase = r.period_evt_ref
		onePosition.PartyBic = r.party_bic
		onePosition.BusinessDate = r.busn_dt
		onePosition.SystemEntity = r.syst_ent
		onePosition.AccountType = r.sec_acc_typ

		data = append(data, onePosition)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal("Error while creating ....")
	}

	return data, err

}

func (r RequestPosition) GetSQLbeforeFilter() (fromSql string) {

	var case1 = ` from "tfsxpositions" a, (SELECT key_part_date, sec_acc_ext,isin, sec_acc_typ,max(ts_creat_row) as ts_creat_row FROM "tfsxpositions" group by key_part_date, sec_acc_ext,isin, sec_acc_typ) b where a.key_part_date = b.key_part_date and   a.sec_acc_ext = b.sec_acc_ext and a.isin = b.isin and a.sec_acc_typ = b.sec_acc_typ and a.ts_creat_row = b.ts_creat_row and `
	var case2 = ` from "tfsxpositions" a where `

	if r.Filter.Phase == "LAST" {
		return case1
	} else {

		return case2
	}
}
