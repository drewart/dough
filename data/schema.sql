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