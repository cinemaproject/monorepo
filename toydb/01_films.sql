create table films
(
	id serial not null
		constraint films_pk
			primary key,
	title varchar(255),
	poster_url varchar(255),
	type varchar(255),
	start_year integer not null,
	end_year integer not null,
	runtime_minutes integer not null,
	imdb_id varchar(255)
);

alter table films owner to postgres;

create unique index films_id_uindex
	on films (id);

create index films_title_index
	on films (title);

