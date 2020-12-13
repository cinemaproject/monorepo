create table people
(
	id char(14) not null
		constraint people_pk
			primary key,
	primary_name varchar(255) not null,
	photo_url varchar(255),
	birth_year integer,
	death_year integer
);

alter table people owner to postgres;

create unique index people_id_uindex
	on people (id);

create index people_primary_name_index
	on people (primary_name);

