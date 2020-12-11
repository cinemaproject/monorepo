create table films
(
	id char(10) not null
		constraint films_pk
			primary key,
	title varchar(255),
	type varchar(255),
	start_year integer,
	end_year integer,
	runtime_minutes integer
);

alter table films owner to postgres;

create unique index films_id_uindex
	on films (id);

create index films_title_index
	on films (title);

