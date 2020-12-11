CREATE TYPE relation_type AS ENUM ('actor', 'director', 'writer');

create table relations
(
	film_id char(10) not null
		constraint relations_films_id_fk
			references films
				on update cascade on delete restrict,
	person_id char(14) not null
		constraint relations_people_id_fk
			references people
				on update cascade on delete restrict,
	relation relation_type not null
);

alter table relations owner to postgres;

create index relations_person_id_index
	on relations (person_id);

create index relations_film_id_index
	on relations (film_id);

