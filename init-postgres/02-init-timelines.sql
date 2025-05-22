CREATE TABLE IF NOT EXISTS public.timelines (
                                                user_id      varchar   not null,
                                                post_id      varchar   not null,
                                                created_at   timestamp not null,
                                                published_at timestamp not null,
                                                primary key (user_id, post_id)
    );

ALTER TABLE public.timelines OWNER TO eduardopena;
