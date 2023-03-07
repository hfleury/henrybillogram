CREATE TABLE IF NOT EXISTS public.discount (
    discount_id SERIAL PRIMARY KEY,
    brand_id INT NOT NULL,
    discount_code VARCHAR(180) NOT NULL,
    discount_created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    discount_used_at TIMESTAMPTZ,
    CONSTRAINT fk_brand FOREIGN KEY(brand_id) REFERENCES brand(brand_id)
);