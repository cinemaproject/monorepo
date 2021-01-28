create table films
(
	id integer not null
		constraint films_pk
			AUTO_INCREMENT
			primary key,
	title varchar(255),
	poster_url varchar(255),
	type varchar(255),
	start_year integer,
	end_year integer,
	runtime_minutes integer,
	imdb_id varchar(255)
);

alter table films owner to postgres;

create unique index films_id_uindex
	on films (id);

create index films_title_index
	on films (title);

