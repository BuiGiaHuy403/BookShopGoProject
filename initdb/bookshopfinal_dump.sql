--
-- PostgreSQL database dump
--

-- Dumped from database version 15.10
-- Dumped by pg_dump version 15.10

-- Started on 2025-01-14 11:20:27

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
-- TOC entry 858 (class 1247 OID 16972)
-- Name: order_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.order_status AS ENUM (
    'pending',
    'shipped',
    'delivered',
    'cancelled'
);


ALTER TYPE public.order_status OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 217 (class 1259 OID 16937)
-- Name: authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authors (
    author_id integer NOT NULL,
    author_name character varying(100) NOT NULL,
    status boolean DEFAULT true
);


ALTER TABLE public.authors OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16936)
-- Name: authors_author_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.authors ALTER COLUMN author_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.authors_author_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 218 (class 1259 OID 16943)
-- Name: book_authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.book_authors (
    book_id integer NOT NULL,
    author_id integer NOT NULL
);


ALTER TABLE public.book_authors OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16928)
-- Name: books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.books (
    book_id integer NOT NULL,
    book_name character varying(100) NOT NULL,
    description text,
    price numeric(10,2) NOT NULL,
    genre character varying(100) NOT NULL,
    status boolean DEFAULT true
);


ALTER TABLE public.books OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16927)
-- Name: books_book_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.books ALTER COLUMN book_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.books_book_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 223 (class 1259 OID 16988)
-- Name: order_books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_books (
    order_id integer NOT NULL,
    book_id integer NOT NULL,
    quantity integer NOT NULL
);


ALTER TABLE public.order_books OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 16982)
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    order_id integer NOT NULL,
    total_price numeric(10,2),
    status public.order_status NOT NULL,
    order_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16981)
-- Name: orders_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.orders ALTER COLUMN order_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.orders_order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 220 (class 1259 OID 16959)
-- Name: stocks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.stocks (
    stock_id integer NOT NULL,
    book_id integer NOT NULL,
    quantity integer NOT NULL,
    update_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT stocks_quantity_check CHECK ((quantity >= 0))
);


ALTER TABLE public.stocks OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16958)
-- Name: stocks_stock_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.stocks ALTER COLUMN stock_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.stocks_stock_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 3366 (class 0 OID 16937)
-- Dependencies: 217
-- Data for Name: authors; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (1, 'J.K. Rowling2', true) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (2, 'George Orwell', true) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (3, 'J.R.R. Tolkien', true) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (4, 'Agatha Christie', true) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (5, 'Dan Brown', true) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (6, 'Stephen King', true) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (7, 'Harper Lee', true) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (8, 'F. Scott Fitzgerald', true) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (9, 'Isaac Asimov', true) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (10, 'Jane Austen', true) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (11, '', false) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (12, '', false) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (13, '', false) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (14, '', false) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (15, '', false) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (16, '', false) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (17, '', false) ON CONFLICT DO NOTHING;
INSERT INTO public.authors OVERRIDING SYSTEM VALUE VALUES (18, '', false) ON CONFLICT DO NOTHING;


--
-- TOC entry 3367 (class 0 OID 16943)
-- Dependencies: 218
-- Data for Name: book_authors; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.book_authors VALUES (1, 1) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (2, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (3, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (4, 4) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (5, 5) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (6, 6) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (7, 7) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (8, 8) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (9, 9) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (10, 10) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (1, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (2, 5) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (5, 6) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (11, 11) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (12, 12) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (13, 13) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (14, 14) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (15, 15) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (16, 16) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (17, 17) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (18, 18) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (29, 1) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (30, 1) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (31, 1) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (37, 1) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (38, 1) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (39, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (39, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (39, 4) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (40, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (40, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (40, 4) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (41, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (41, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (41, 4) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (42, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (42, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (42, 4) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (43, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (43, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (43, 4) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (44, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (44, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (44, 4) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (45, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (45, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (45, 4) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (46, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (46, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (46, 4) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (47, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (47, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (47, 4) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (49, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (49, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.book_authors VALUES (49, 4) ON CONFLICT DO NOTHING;


--
-- TOC entry 3364 (class 0 OID 16928)
-- Dependencies: 215
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (1, 'Harry Potter and the Sorcerer s Stone', 'A young wizard discovers his magical heritage.', 19.99, 'Fantasy', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (2, '1984', 'A dystopian novel about a totalitarian regime.', 14.99, 'Dystopian', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (3, 'The Hobbit', 'A hobbit goes on an adventure to reclaim a stolen treasure.', 15.99, 'Fantasy', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (4, 'Murder on the Orient Express', 'A detective investigates a murder on a train.', 12.99, 'Mystery', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (5, 'The Da Vinci Code', 'A symbologist uncovers religious secrets hidden in art.', 18.99, 'Thriller', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (6, 'The Shining', 'A family isolated in a haunted hotel struggles to survive.', 16.99, 'Horror', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (7, 'To Kill a Mockingbird', 'A young girl learns about racial injustice in the South.', 13.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (8, 'The Great Gatsby', 'A wealthy man tries to rekindle a past romance in the roaring twenties.', 17.99, 'Classic', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (9, 'Foundation', 'A scientist predicts the fall of an empire and builds a foundation to preserve knowledge.', 22.99, 'Sci-Fi', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (10, 'Pride and Prejudice', 'A classic tale of love and societal expectations in the 19th century.', 11.99, 'Romance', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (11, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', false) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (12, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', false) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (13, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (14, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (15, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (16, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (17, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (18, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (23, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (24, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (25, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (26, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (27, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (28, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (29, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (30, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (31, 'The Great Book', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (37, 'The Great BOOK2', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (38, 'The Great BOOK3', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (39, 'The Great BOOK4', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (40, 'The Great BOOK5', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (41, 'The Great BOOK6', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (42, 'The Great BOOK7', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (43, 'The Great BOOK8', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (44, 'The Great BOOK9', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (45, 'The Great BOOK10', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (46, 'The Great BOOK11', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (47, 'The Great BOOK12', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;
INSERT INTO public.books OVERRIDING SYSTEM VALUE VALUES (49, 'The Great BOOK13', 'A fascinating story about the world', 19.99, 'Fiction', true) ON CONFLICT DO NOTHING;


--
-- TOC entry 3372 (class 0 OID 16988)
-- Dependencies: 223
-- Data for Name: order_books; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.order_books VALUES (1, 1, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.order_books VALUES (1, 2, 1) ON CONFLICT DO NOTHING;
INSERT INTO public.order_books VALUES (2, 3, 3) ON CONFLICT DO NOTHING;
INSERT INTO public.order_books VALUES (2, 4, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.order_books VALUES (3, 5, 1) ON CONFLICT DO NOTHING;
INSERT INTO public.order_books VALUES (3, 6, 2) ON CONFLICT DO NOTHING;
INSERT INTO public.order_books VALUES (4, 7, 5) ON CONFLICT DO NOTHING;


--
-- TOC entry 3371 (class 0 OID 16982)
-- Dependencies: 222
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.orders OVERRIDING SYSTEM VALUE VALUES (1, 200.00, 'pending', '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;
INSERT INTO public.orders OVERRIDING SYSTEM VALUE VALUES (2, 250.00, 'shipped', '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;
INSERT INTO public.orders OVERRIDING SYSTEM VALUE VALUES (3, 300.00, 'delivered', '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;
INSERT INTO public.orders OVERRIDING SYSTEM VALUE VALUES (4, 150.00, 'cancelled', '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;


--
-- TOC entry 3369 (class 0 OID 16959)
-- Dependencies: 220
-- Data for Name: stocks; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (1, 1, 50, '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (2, 2, 30, '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (3, 3, 40, '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (4, 4, 20, '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (5, 5, 60, '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (6, 6, 70, '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (7, 7, 80, '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (8, 8, 90, '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (9, 9, 50, '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (10, 10, 120, '2025-01-02 11:08:33.85295') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (11, 37, 100, '0001-01-01 00:00:00') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (12, 38, 100, '2025-01-03 16:29:15.337509') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (13, 39, 100, '2025-01-03 16:30:17.675974') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (14, 40, 100, '2025-01-06 09:41:28.095939') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (15, 41, 100, '2025-01-06 09:43:46.845248') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (16, 42, 100, '2025-01-06 09:44:31.348124') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (17, 43, 100, '2025-01-06 09:44:47.426897') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (18, 44, 100, '2025-01-06 09:44:52.534369') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (19, 45, 100, '2025-01-06 09:45:20.607302') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (20, 46, 100, '2025-01-07 11:19:02.729023') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (21, 47, 100, '2025-01-07 11:20:45.110656') ON CONFLICT DO NOTHING;
INSERT INTO public.stocks OVERRIDING SYSTEM VALUE VALUES (22, 49, 100, '2025-01-07 11:23:56.550037') ON CONFLICT DO NOTHING;


--
-- TOC entry 3378 (class 0 OID 0)
-- Dependencies: 216
-- Name: authors_author_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.authors_author_id_seq', 18, true);


--
-- TOC entry 3379 (class 0 OID 0)
-- Dependencies: 214
-- Name: books_book_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.books_book_id_seq', 49, true);


--
-- TOC entry 3380 (class 0 OID 0)
-- Dependencies: 221
-- Name: orders_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_order_id_seq', 4, true);


--
-- TOC entry 3381 (class 0 OID 0)
-- Dependencies: 219
-- Name: stocks_stock_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.stocks_stock_id_seq', 22, true);


--
-- TOC entry 3207 (class 2606 OID 16942)
-- Name: authors authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors
    ADD CONSTRAINT authors_pkey PRIMARY KEY (author_id);


--
-- TOC entry 3209 (class 2606 OID 16947)
-- Name: book_authors book_authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.book_authors
    ADD CONSTRAINT book_authors_pkey PRIMARY KEY (book_id, author_id);


--
-- TOC entry 3205 (class 2606 OID 16935)
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (book_id);


--
-- TOC entry 3215 (class 2606 OID 16992)
-- Name: order_books order_books_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_books
    ADD CONSTRAINT order_books_pkey PRIMARY KEY (order_id, book_id);


--
-- TOC entry 3213 (class 2606 OID 16987)
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_id);


--
-- TOC entry 3211 (class 2606 OID 16965)
-- Name: stocks stocks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stocks
    ADD CONSTRAINT stocks_pkey PRIMARY KEY (stock_id);


--
-- TOC entry 3216 (class 2606 OID 16953)
-- Name: book_authors book_authors_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.book_authors
    ADD CONSTRAINT book_authors_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.authors(author_id) ON DELETE CASCADE;


--
-- TOC entry 3217 (class 2606 OID 16948)
-- Name: book_authors book_authors_book_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.book_authors
    ADD CONSTRAINT book_authors_book_id_fkey FOREIGN KEY (book_id) REFERENCES public.books(book_id) ON DELETE CASCADE;


--
-- TOC entry 3219 (class 2606 OID 16998)
-- Name: order_books order_books_book_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_books
    ADD CONSTRAINT order_books_book_id_fkey FOREIGN KEY (book_id) REFERENCES public.books(book_id) ON DELETE CASCADE;


--
-- TOC entry 3220 (class 2606 OID 16993)
-- Name: order_books order_books_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_books
    ADD CONSTRAINT order_books_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(order_id) ON DELETE CASCADE;


--
-- TOC entry 3218 (class 2606 OID 16966)
-- Name: stocks stocks_book_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stocks
    ADD CONSTRAINT stocks_book_id_fkey FOREIGN KEY (book_id) REFERENCES public.books(book_id) ON DELETE CASCADE;


-- Completed on 2025-01-14 11:20:27

--
-- PostgreSQL database dump complete
--

