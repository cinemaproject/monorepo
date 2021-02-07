-- +migrate Up
create table if not exists phrases
(
	film_id integer not null
		constraint phrases_films_id_fk
			references films
				on update cascade on delete restrict,
	phrase tsvector not null,
	time_offset integer
);

alter table phrases owner to postgres;

create index if not exists idx_gin_phrases on phrases using gin ("phrase");

-- +migrate Down
drop table phrases;
