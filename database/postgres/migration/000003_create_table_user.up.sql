CREATE TABLE IF NOT EXISTS public.user (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(80) NOT NULL
);

INSERT INTO public.user(user_name) 
VALUES('user1'),('user2'),('user2'),('user2');

ALTER TABLE public.discount ADD COLUMN user_id INT
CONSTRAINT fk_user REFERENCES public.user (user_id); 