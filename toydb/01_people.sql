create table people
(
	id integer not null
		constraint people_pk
			AUTO_INCREMENT
			primary key,
	primary_name varchar(255) not null,
	photo_url varchar(255),
	birth_year integer,
	death_year integer,
	imdb_id varchar(255)
);

alter table people owner to postgres;

create unique index people_id_uindex
	on people (id);

create index people_primary_name_index
	on people (primary_name);

