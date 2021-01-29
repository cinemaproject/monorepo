create table people
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

create unique index people_id_uindex
	on people (id);

create index people_primary_name_index
	on people (primary_name);

