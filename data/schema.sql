--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4.5
-- Dumped by pg_dump version 9.4.0
-- Started on 2016-01-10 23:31:19 EST

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- TOC entry 178 (class 3079 OID 11861)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2031 (class 0 OID 0)
-- Dependencies: 178
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 173 (class 1259 OID 16388)
-- Name: actor; Type: TABLE; Schema: public; Owner: demo; Tablespace: 
--

CREATE TABLE actor (
    id integer NOT NULL,
    name text NOT NULL,
    create_dt timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);


ALTER TABLE actor OWNER TO demo;

--
-- TOC entry 172 (class 1259 OID 16386)
-- Name: actor_id_seq; Type: SEQUENCE; Schema: public; Owner: demo
--

CREATE SEQUENCE actor_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE actor_id_seq OWNER TO demo;

--
-- TOC entry 2032 (class 0 OID 0)
-- Dependencies: 172
-- Name: actor_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: demo
--

ALTER SEQUENCE actor_id_seq OWNED BY actor.id;


--
-- TOC entry 175 (class 1259 OID 16406)
-- Name: movie; Type: TABLE; Schema: public; Owner: demo; Tablespace: 
--

CREATE TABLE movie (
    id integer NOT NULL,
    title text NOT NULL,
    create_dt timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);


ALTER TABLE movie OWNER TO demo;

--
-- TOC entry 174 (class 1259 OID 16404)
-- Name: movie_id_seq; Type: SEQUENCE; Schema: public; Owner: demo
--

CREATE SEQUENCE movie_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE movie_id_seq OWNER TO demo;

--
-- TOC entry 2033 (class 0 OID 0)
-- Dependencies: 174
-- Name: movie_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: demo
--

ALTER SEQUENCE movie_id_seq OWNED BY movie.id;


--
-- TOC entry 177 (class 1259 OID 16418)
-- Name: movieactor; Type: TABLE; Schema: public; Owner: demo; Tablespace: 
--

CREATE TABLE movieactor (
    id integer NOT NULL,
    movie_id integer NOT NULL,
    actor_id integer NOT NULL,
    create_dt timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);


ALTER TABLE movieactor OWNER TO demo;

--
-- TOC entry 176 (class 1259 OID 16416)
-- Name: movieactor_id_seq; Type: SEQUENCE; Schema: public; Owner: demo
--

CREATE SEQUENCE movieactor_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE movieactor_id_seq OWNER TO demo;

--
-- TOC entry 2034 (class 0 OID 0)
-- Dependencies: 176
-- Name: movieactor_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: demo
--

ALTER SEQUENCE movieactor_id_seq OWNED BY movieactor.id;


--
-- TOC entry 1899 (class 2604 OID 16391)
-- Name: id; Type: DEFAULT; Schema: public; Owner: demo
--

ALTER TABLE ONLY actor ALTER COLUMN id SET DEFAULT nextval('actor_id_seq'::regclass);


--
-- TOC entry 1901 (class 2604 OID 16409)
-- Name: id; Type: DEFAULT; Schema: public; Owner: demo
--

ALTER TABLE ONLY movie ALTER COLUMN id SET DEFAULT nextval('movie_id_seq'::regclass);


--
-- TOC entry 1903 (class 2604 OID 16421)
-- Name: id; Type: DEFAULT; Schema: public; Owner: demo
--

ALTER TABLE ONLY movieactor ALTER COLUMN id SET DEFAULT nextval('movieactor_id_seq'::regclass);


--
-- TOC entry 1906 (class 2606 OID 16403)
-- Name: pk_actor; Type: CONSTRAINT; Schema: public; Owner: demo; Tablespace: 
--

ALTER TABLE ONLY actor
    ADD CONSTRAINT pk_actor PRIMARY KEY (id);


--
-- TOC entry 1908 (class 2606 OID 16415)
-- Name: pk_movie; Type: CONSTRAINT; Schema: public; Owner: demo; Tablespace: 
--

ALTER TABLE ONLY movie
    ADD CONSTRAINT pk_movie PRIMARY KEY (id);


--
-- TOC entry 1912 (class 2606 OID 16424)
-- Name: pk_movieactor; Type: CONSTRAINT; Schema: public; Owner: demo; Tablespace: 
--

ALTER TABLE ONLY movieactor
    ADD CONSTRAINT pk_movieactor PRIMARY KEY (id);


--
-- TOC entry 1909 (class 1259 OID 16436)
-- Name: fki_actor_id; Type: INDEX; Schema: public; Owner: demo; Tablespace: 
--

CREATE INDEX fki_actor_id ON movieactor USING btree (actor_id);


--
-- TOC entry 1910 (class 1259 OID 16430)
-- Name: fki_movie_id; Type: INDEX; Schema: public; Owner: demo; Tablespace: 
--

CREATE INDEX fki_movie_id ON movieactor USING btree (movie_id);


--
-- TOC entry 1914 (class 2606 OID 16431)
-- Name: fk_actor_id; Type: FK CONSTRAINT; Schema: public; Owner: demo
--

ALTER TABLE ONLY movieactor
    ADD CONSTRAINT fk_actor_id FOREIGN KEY (actor_id) REFERENCES actor(id);


--
-- TOC entry 1913 (class 2606 OID 16425)
-- Name: fk_movie_id; Type: FK CONSTRAINT; Schema: public; Owner: demo
--

ALTER TABLE ONLY movieactor
    ADD CONSTRAINT fk_movie_id FOREIGN KEY (movie_id) REFERENCES movie(id);


--
-- TOC entry 2030 (class 0 OID 0)
-- Dependencies: 5
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2016-01-10 23:31:19 EST

--
-- PostgreSQL database dump complete
--

