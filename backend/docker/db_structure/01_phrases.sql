create table phrases
(
	film_id char(10) not null
		constraint phrases_films_id_fk
			references films,
	phrase text not null,
	time_offset integer
);

alter table phrases owner to postgres;

