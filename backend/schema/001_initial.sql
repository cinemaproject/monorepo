-- +migrate Up
create table if not exists films
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

create unique index if not exists films_id_uindex
	on films (id);

create index if not exists films_title_index
	on films (title);

create table if not exists people
(
	id serial not null
		constraint people_pk
			primary key,
	primary_name varchar(255) not null,
	photo_url varchar(255),
	birth_year integer not null,
	death_year integer not null,
	imdb_id varchar(255)
);

alter table people owner to postgres;

create unique index if not exists people_id_uindex
	on people (id);

create index if not exists people_primary_name_index
	on people (primary_name);

create table if not exists relations
(
	film_id integer not null
		constraint relations_films_id_fk
			references films
				on update cascade on delete restrict,
	person_id integer not null
		constraint relations_people_id_fk
			references people
				on update cascade on delete restrict,
	relation integer not null
);

alter table relations owner to postgres;

create index if not exists relations_person_id_index
	on relations (person_id);

create index if not exists relations_film_id_index
	on relations (film_id);

-- +migrate Down
drop table relations;
drop table films;
drop table people;
