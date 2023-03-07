CREATE TABLE IF NOT EXISTS public.brand (
    brand_id SERIAL PRIMARY KEY,
    brand_name varchar NOT NULL
);

INSERT INTO public.brand(brand_name) 
VALUES('brand1'),('brand2'),('brand3'),('brand4');