\connect posts;

CREATE TABLE IF NOT EXISTS public.posts (
                                            id           varchar(36)              not null primary key,
    content      jsonb                    not null,
    author_id    varchar(36)              not null,
    created_at   timestamp with time zone not null,
    updated_at   timestamp with time zone not null,
    deleted_at   timestamp with time zone,
    published_at timestamp with time zone not null
                               );

CREATE INDEX idx_posts_author_id ON public.posts (author_id);
CREATE INDEX idx_posts_created_at ON public.posts (created_at);
CREATE INDEX idx_posts_deleted_at ON public.posts (deleted_at);

ALTER TABLE public.posts OWNER TO eduardopena;
