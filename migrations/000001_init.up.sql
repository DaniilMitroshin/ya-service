CREATE TABLE books(
	id bigint generated always as identity primary key,
	title text,
	author text,
	num_pages integer,
	rating double precision 
	);
	