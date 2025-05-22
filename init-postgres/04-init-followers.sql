\connect followers;

CREATE TABLE IF NOT EXISTS public.follows (
                                              follower_id varchar(255) not null,
    followed_id varchar(255) not null,
    created_at  timestamp default CURRENT_TIMESTAMP,
    primary key (follower_id, followed_id)
    );

CREATE INDEX idx_follower_id ON public.follows (follower_id);
CREATE INDEX idx_followed_id ON public.follows (followed_id);

ALTER TABLE public.follows OWNER TO eduardopena;
