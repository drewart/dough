package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func getDBFilePath() *string {
	userDir, _ := os.UserHomeDir()
	userDataDir := path.Join(userDir, ".dough")
	dbFile := path.Join(userDataDir, "dough.db")
	return &dbFile
}

func InitSchema(dbFile *string) {

	//os.Remove("./dough.db")

	if dbFile == nil || *dbFile == "" {
		dbFile = getDBFilePath()
	}
	parentDir := path.Dir(*dbFile)

	log.Printf("Setting up db: %s", *dbFile)

	fi, err := os.Stat(parentDir)
	if err == os.ErrNotExist || err != nil {
		log.Println("Creating DB")
		e := os.Mkdir(parentDir, 0755)
		if e != nil {
			log.Printf("Error creating %s\n%s", parentDir, e)
		}
	}
	log.Printf("Fileinfo bar fi %v err:%s", fi, err)

	db, err := sql.Open("sqlite3", *dbFile)
	if err != nil {
		log.Fatalf("sqlite3 load failure: %s", err)
	}
	defer db.Close()

	rows, err := db.Query(`select name from sqlite_master where type='table'`)
	if err != nil {
		log.Fatalf("error getting tablename data: %s", err)
	}
	var tableNames []string
	var name string
	for rows.Next() {
		rows.Scan(&name)
		tableNames = append(tableNames, name)
	}
	for _, n := range tableNames {

		dropStmt := fmt.Sprintf("drop table %s", n)

		log.Println("Running :" + dropStmt)

		_, err := db.Exec(dropStmt)
		if err != nil {
			log.Printf("error %s", err)
		}
	}

	sqlStmt := ` 
	create table Account (
		id integer not null primary key,
		name text,
		account_type text,
		on_budget integer
		);

	create table Ledger (
		id integer not null primary key,
		account_id integer,
		date numeric,
		tran_type text,
		cat_id integer
		payee text,
		memo text,
		check_number text,
		amount integer,
		verified integer
		);

	create table Balance (
		id integer not null primary key,
		account_id integer,
		date integer,
		balance integer
		);

	create table Category (
		id integer not null primary key,
		name text,
		parent_id integer,
		code text,
		tags text,
		pos integer,
		active integer
		);
	
	create table CategoryMatch (
		id integer 	not null primary key, 
		cat_id integer,
		is_check integer,
		key_term text,
		terms text,
		replace text,
		ammount_match int
	);

	create table Budget(
		id integer not null primary key,
		name text,
		created integer
	);
	
	create table BudgetCategory(
		id integer non null primary key,
		budget_month integer not null,
		cat_id integer non null,
		amount integer,
		notes text,
		active integer
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	log.Printf("dough db created")
}

type DoughStorage struct {
	db       *sql.DB
	Accounts []Account
}

func NewDoughStorage() *DoughStorage {
	dbFile := getDBFilePath()
	db, err := sql.Open("sqlite3", *dbFile)
	if err != nil {
		log.Fatal(err)
	}
	return &DoughStorage{db: db}
}

func (d *DoughStorage) Close() {
	d.db.Close()
}

func (d *DoughStorage) InsertAccount(id int, name string) {

	tx, err := d.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into Account(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, name)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DoughStorage) GetAccounts() []Account {
	var accounts []Account

	rows, err := d.db.Query("select id, name, account_type,on_budget from Account")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, onBudgetInt int
		var name, accountType string
		var onBudget bool

		err = rows.Scan(&id, &name, &accountType, onBudgetInt)
		if err != nil {
			log.Fatal(err)
		} else {
			onBudget = (onBudgetInt > 0)
			accounts = append(accounts, Account{id, name, accountType, onBudget})
		}
		//fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return accounts
}

func (d *DoughStorage) GetLedgerEnteries(start int, end int) []LedgerEntry {
	var entries []LedgerEntry

	rows, err := d.db.Query(fmt.Sprintf(`select id,
								account_id, 
								date, 
								transaction_type, 
								payee, 
								memo, 
								check_number, 
								amount, 
								verified 
								FROM Ledger 
								where date > %d and date < %d`, start, end))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var entry LedgerEntry
		var id, account_id, dateInt, amount, verified int
		var transType, payee, memo, checkNum string

		rows.Scan(&id, &account_id, &dateInt, &transType, &payee, &memo, &checkNum, &amount, &verified)

		entry.ID = id
		entry.Amount = amount
		entry.Payee = payee
		entry.Date = time.Unix(int64(dateInt), 0)
		entry.TransType = transType
		entry.Memo = memo
		entry.Check = checkNum
		entry.Verified = (verified > 0)
		entries = append(entries, entry)
	}
	return entries
}

func (d *DoughStorage) InsertLeger(entry LedgerEntry) {

	tx, err := d.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(`insert into Ledger(id, account_id, date, tran_type, cat_id, payee, memo, check_number, amount, verified ) 
								values(?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(entry.ID, entry.Account.ID, entry.Date.Unix())
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DoughStorage) InsertCatagory(cat Catagory) {

	tx, err := d.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(`insert into Category(id, name, parent_id, code, tags, pos, active) 
								values(?, ?, ?, ?, ?, ?, ? )`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	tags := ""
	parent_id := 0
	if cat.Parent != nil {
		parent_id = cat.Parent.ID
	}
	if len(cat.Tags) > 0 {
		tags = strings.Join(cat.Tags, "|")
	}

	active := 0
	if cat.Active {
		active = 1
	}
	_, err = stmt.Exec(cat.ID, cat.Name, parent_id, cat.Code, tags, cat.Pos, active)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DoughStorage) GetCategories() []*Catagory {
	var cats []*Catagory

	rows, err := d.db.Query("select id, name, parent_id, code, tags, pos, active from Category where active != 0")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, active, parent_id, pos int
		var name, tags, code string

		err = rows.Scan(&id, &name, &parent_id, &code, &tags, &pos, &active)
		if err != nil {
			log.Fatal(err)
		} else {
			parent := GetCatById(parent_id)

			tagSplit := strings.Split(tags, "|")

			bActive := (active > 0)

			cat, err := NewCatagory(id, name, code, parent, tagSplit, pos, bActive)
			if err != nil {
				log.Printf("error add cat %s", err)
			} else {
				cats = append(cats, cat)
			}
		}
		//fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return cats
}

func Misc() {
	dbFile := getDBFilePath()
	db, err := sql.Open("sqlite3", *dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into Account(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("account_%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select id, name from Account")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("select name from Account where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	_, err = db.Exec("delete from Account")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into Account(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, name from Account")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
