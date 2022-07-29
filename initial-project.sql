-- Table: public.assets

-- DROP TABLE IF EXISTS public.assets;
CREATE SEQUENCE assets_id_seq;
CREATE TABLE IF NOT EXISTS public.assets
(
    id integer NOT NULL DEFAULT nextval('assets_id_seq'),
    date_created timestamp with time zone,
    date_updated timestamp with time zone,
    symbol character varying(255) COLLATE pg_catalog."default",
    logo character varying(255) COLLATE pg_catalog."default",
    name character varying(255) COLLATE pg_catalog."default",
    last_price character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT assets_pkey PRIMARY KEY (id)
);

ALTER SEQUENCE assets_id_seq OWNED BY assets.id;

ALTER TABLE IF EXISTS public.assets
    OWNER to root;

INSERT INTO public.assets(
	id, date_created, date_updated, symbol, logo, name, last_price)
	VALUES (1,current_timestamp, null, 'BTC', 'https://cdn.pixabay.com/photo/2013/12/08/12/12/bitcoin-225079_960_720.png', 'Bitcoin', 10000);
INSERT INTO public.assets(
	id, date_created, date_updated, symbol, logo, name, last_price)
	VALUES (2,current_timestamp, null, 'ETH', 'https://w7.pngwing.com/pngs/268/1013/png-transparent-ethereum-eth-hd-logo-thumbnail.png', 'Ethereum', 500);
INSERT INTO public.assets(
	id, date_created, date_updated, symbol, logo, name, last_price)
	VALUES (3,current_timestamp, null, 'THB', 'https://w7.pngwing.com/pngs/73/573/png-transparent-thai-baht-ripple-omisego-coin-money-money-baht-text-trademark-investment-thumbnail.png', 'Thai bath', 32);


-- Table: public.users

-- DROP TABLE IF EXISTS public.users;
CREATE SEQUENCE users_id_seq;
CREATE TABLE IF NOT EXISTS public.users
(
    id integer NOT NULL DEFAULT nextval('users_id_seq'),
    status character varying(255) COLLATE pg_catalog."default" NOT NULL DEFAULT 'draft'::character varying,
    date_created timestamp with time zone,
    date_updated timestamp with time zone,
    first_name character varying(255) COLLATE pg_catalog."default",
    last_name character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

ALTER SEQUENCE users_id_seq OWNED BY users.id;

ALTER TABLE IF EXISTS public.users
    OWNER to root;

INSERT INTO public.users(
	id, status, date_created, date_updated, first_name, last_name)
	VALUES (1,'publish', current_timestamp, null, 'test', 'system');


-- Table: public.wallets

-- DROP TABLE IF EXISTS public.wallets;
CREATE SEQUENCE wallets_id_seq;
CREATE TABLE IF NOT EXISTS public.wallets
(
    id integer NOT NULL DEFAULT nextval('wallets_id_seq'),
    date_created timestamp with time zone,
    date_updated timestamp with time zone,
    asset_id integer,
    user_id integer,
    amount bigint,
    CONSTRAINT wallets_pkey PRIMARY KEY (id),
    CONSTRAINT wallets_asset_id_foreign FOREIGN KEY (asset_id)
        REFERENCES public.assets (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE SET NULL,
    CONSTRAINT wallets_user_id_foreign FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE SET NULL
);

ALTER SEQUENCE wallets_id_seq OWNED BY wallets.id;


ALTER TABLE IF EXISTS public.wallets
    OWNER to root;
	
INSERT INTO public.wallets(
	 date_created, date_updated, asset_id, user_id, amount)
	VALUES ( current_timestamp, null, 3, 1, 100000);


-- Table: public.orders

-- DROP TABLE IF EXISTS public.orders;

CREATE SEQUENCE orders_id_seq;
CREATE TABLE IF NOT EXISTS public.orders
(
    id integer NOT NULL DEFAULT nextval('orders_id_seq'),
    date_created timestamp with time zone,
    date_updated timestamp with time zone,
    order_type character varying(255) COLLATE pg_catalog."default",
    price_action numeric(10,5),
    amount numeric(10,5),
    user_id integer,
    asset_id integer,
    CONSTRAINT orders_pkey PRIMARY KEY (id),
    CONSTRAINT orders_asset_id_foreign FOREIGN KEY (asset_id)
        REFERENCES public.assets (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE SET NULL,
    CONSTRAINT orders_user_id_foreign FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE SET NULL
);

ALTER SEQUENCE orders_id_seq OWNED BY orders.id;

ALTER TABLE IF EXISTS public.orders
    OWNER to root;