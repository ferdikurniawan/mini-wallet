--
-- PostgreSQL database dump
--

-- Dumped from database version 14.0
-- Dumped by pg_dump version 14.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: account_status; Type: TYPE; Schema: public; Owner: ferdinandkurniawan
--

CREATE TYPE public.account_status AS ENUM (
    'disabled',
    'enabled'
);


ALTER TYPE public.account_status OWNER TO ferdinandkurniawan;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: mw_accounts; Type: TABLE; Schema: public; Owner: ferdinandkurniawan
--

CREATE TABLE public.mw_accounts (
    id integer NOT NULL,
    customer_xid character varying(36) NOT NULL,
    token character varying(42) NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    salt character varying(10)
);


ALTER TABLE public.mw_accounts OWNER TO ferdinandkurniawan;

--
-- Name: mw_accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: ferdinandkurniawan
--

CREATE SEQUENCE public.mw_accounts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mw_accounts_id_seq OWNER TO ferdinandkurniawan;

--
-- Name: mw_accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ferdinandkurniawan
--

ALTER SEQUENCE public.mw_accounts_id_seq OWNED BY public.mw_accounts.id;


--
-- Name: mw_activity_logs; Type: TABLE; Schema: public; Owner: ferdinandkurniawan
--

CREATE TABLE public.mw_activity_logs (
    id integer NOT NULL,
    account_id integer,
    activity character varying(255) NOT NULL,
    activity_time timestamp with time zone DEFAULT now()
);


ALTER TABLE public.mw_activity_logs OWNER TO ferdinandkurniawan;

--
-- Name: mw_activity_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: ferdinandkurniawan
--

CREATE SEQUENCE public.mw_activity_logs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mw_activity_logs_id_seq OWNER TO ferdinandkurniawan;

--
-- Name: mw_activity_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ferdinandkurniawan
--

ALTER SEQUENCE public.mw_activity_logs_id_seq OWNED BY public.mw_activity_logs.id;


--
-- Name: mw_transaction_logs; Type: TABLE; Schema: public; Owner: ferdinandkurniawan
--

CREATE TABLE public.mw_transaction_logs (
    id integer NOT NULL,
    wallet_id integer,
    balance_before bigint,
    deposit_amt bigint DEFAULT '0'::bigint,
    withdraw_amt bigint DEFAULT '0'::bigint,
    balance_after bigint,
    created_at timestamp with time zone,
    reference_id character varying(36) NOT NULL,
    transaction_id character varying(36) NOT NULL,
    created_by integer NOT NULL,
    status character varying(10) NOT NULL
);


ALTER TABLE public.mw_transaction_logs OWNER TO ferdinandkurniawan;

--
-- Name: mw_balance_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: ferdinandkurniawan
--

CREATE SEQUENCE public.mw_balance_logs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mw_balance_logs_id_seq OWNER TO ferdinandkurniawan;

--
-- Name: mw_balance_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ferdinandkurniawan
--

ALTER SEQUENCE public.mw_balance_logs_id_seq OWNED BY public.mw_transaction_logs.id;


--
-- Name: mw_wallets; Type: TABLE; Schema: public; Owner: ferdinandkurniawan
--

CREATE TABLE public.mw_wallets (
    id integer NOT NULL,
    wallet_id character varying(36) NOT NULL,
    account_id integer,
    status public.account_status DEFAULT 'disabled'::public.account_status,
    balance bigint DEFAULT '0'::bigint,
    enabled_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone
);


ALTER TABLE public.mw_wallets OWNER TO ferdinandkurniawan;

--
-- Name: mw_wallets_id_seq; Type: SEQUENCE; Schema: public; Owner: ferdinandkurniawan
--

CREATE SEQUENCE public.mw_wallets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mw_wallets_id_seq OWNER TO ferdinandkurniawan;

--
-- Name: mw_wallets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ferdinandkurniawan
--

ALTER SEQUENCE public.mw_wallets_id_seq OWNED BY public.mw_wallets.id;


--
-- Name: mw_accounts id; Type: DEFAULT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_accounts ALTER COLUMN id SET DEFAULT nextval('public.mw_accounts_id_seq'::regclass);


--
-- Name: mw_activity_logs id; Type: DEFAULT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_activity_logs ALTER COLUMN id SET DEFAULT nextval('public.mw_activity_logs_id_seq'::regclass);


--
-- Name: mw_transaction_logs id; Type: DEFAULT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_transaction_logs ALTER COLUMN id SET DEFAULT nextval('public.mw_balance_logs_id_seq'::regclass);


--
-- Name: mw_wallets id; Type: DEFAULT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_wallets ALTER COLUMN id SET DEFAULT nextval('public.mw_wallets_id_seq'::regclass);


--
-- Data for Name: mw_accounts; Type: TABLE DATA; Schema: public; Owner: ferdinandkurniawan
--

COPY public.mw_accounts (id, customer_xid, token, created_at, updated_at, salt) FROM stdin;
9	ea0212d3-abd6-406f-8c67-868e814a2438	04ab518fa60e5a00c1d923a5aa32aa28d873580f	2022-09-12 17:43:48.320998+07	\N	54GDMG5UMJ
\.


--
-- Data for Name: mw_activity_logs; Type: TABLE DATA; Schema: public; Owner: ferdinandkurniawan
--

COPY public.mw_activity_logs (id, account_id, activity, activity_time) FROM stdin;
14	9	enable	2022-09-12 17:46:34.473084+07
15	9	disabled	2022-09-12 17:48:33.795203+07
\.


--
-- Data for Name: mw_transaction_logs; Type: TABLE DATA; Schema: public; Owner: ferdinandkurniawan
--

COPY public.mw_transaction_logs (id, wallet_id, balance_before, deposit_amt, withdraw_amt, balance_after, created_at, reference_id, transaction_id, created_by, status) FROM stdin;
13	14	0	100000	0	100000	2022-09-12 17:47:28.422305+07	1234	9fcfafd4-7d84-4fe0-bb51-72a0d220e80d	9	success
15	14	100000	0	50000	50000	2022-09-12 17:48:01.592232+07	abcdef	a57ae9d8-c4cd-46f6-9626-fce5a974b52a	9	success
\.


--
-- Data for Name: mw_wallets; Type: TABLE DATA; Schema: public; Owner: ferdinandkurniawan
--

COPY public.mw_wallets (id, wallet_id, account_id, status, balance, enabled_at, created_at, updated_at) FROM stdin;
14	cf119b6d-97c4-4e48-8043-525a96cdfde4	9	disabled	50000	2022-09-12 17:46:34.473084+07	2022-09-12 17:46:34.473084+07	2022-09-12 17:48:33.795203+07
\.


--
-- Name: mw_accounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ferdinandkurniawan
--

SELECT pg_catalog.setval('public.mw_accounts_id_seq', 12, true);


--
-- Name: mw_activity_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ferdinandkurniawan
--

SELECT pg_catalog.setval('public.mw_activity_logs_id_seq', 16, true);


--
-- Name: mw_balance_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ferdinandkurniawan
--

SELECT pg_catalog.setval('public.mw_balance_logs_id_seq', 15, true);


--
-- Name: mw_wallets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ferdinandkurniawan
--

SELECT pg_catalog.setval('public.mw_wallets_id_seq', 15, true);


--
-- Name: mw_accounts mw_accounts_customer_xid_key; Type: CONSTRAINT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_accounts
    ADD CONSTRAINT mw_accounts_customer_xid_key UNIQUE (customer_xid);


--
-- Name: mw_accounts mw_accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_accounts
    ADD CONSTRAINT mw_accounts_pkey PRIMARY KEY (id);


--
-- Name: mw_accounts mw_accounts_token_key; Type: CONSTRAINT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_accounts
    ADD CONSTRAINT mw_accounts_token_key UNIQUE (token);


--
-- Name: mw_activity_logs mw_activity_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_activity_logs
    ADD CONSTRAINT mw_activity_logs_pkey PRIMARY KEY (id);


--
-- Name: mw_transaction_logs mw_balance_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_transaction_logs
    ADD CONSTRAINT mw_balance_logs_pkey PRIMARY KEY (id);


--
-- Name: mw_transaction_logs mw_transaction_logs_reference_id_key; Type: CONSTRAINT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_transaction_logs
    ADD CONSTRAINT mw_transaction_logs_reference_id_key UNIQUE (reference_id);


--
-- Name: mw_transaction_logs mw_transaction_logs_transaction_id_key; Type: CONSTRAINT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_transaction_logs
    ADD CONSTRAINT mw_transaction_logs_transaction_id_key UNIQUE (transaction_id);


--
-- Name: mw_wallets mw_wallets_pkey; Type: CONSTRAINT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_wallets
    ADD CONSTRAINT mw_wallets_pkey PRIMARY KEY (id);


--
-- Name: mw_wallets mw_wallets_wallet_id_key; Type: CONSTRAINT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_wallets
    ADD CONSTRAINT mw_wallets_wallet_id_key UNIQUE (wallet_id);


--
-- Name: mw_activity_logs mw_activity_logs_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_activity_logs
    ADD CONSTRAINT mw_activity_logs_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.mw_accounts(id);


--
-- Name: mw_transaction_logs mw_balance_logs_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_transaction_logs
    ADD CONSTRAINT mw_balance_logs_wallet_id_fkey FOREIGN KEY (wallet_id) REFERENCES public.mw_wallets(id);


--
-- Name: mw_wallets mw_wallets_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ferdinandkurniawan
--

ALTER TABLE ONLY public.mw_wallets
    ADD CONSTRAINT mw_wallets_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.mw_accounts(id);


--
-- PostgreSQL database dump complete
--

