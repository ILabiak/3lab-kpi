
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
-- Name: forum_service; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA forum_service;


ALTER SCHEMA forum_service OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: forums; Type: TABLE; Schema: forum_service; Owner: postgres
--

CREATE TABLE forum_service.forums (
    id integer NOT NULL,
    name text NOT NULL,
    "topickeyword" text,
    users text[]
);


ALTER TABLE forum_service.forums OWNER TO postgres;

--
-- Name: forums_id_seq; Type: SEQUENCE; Schema: forum_service; Owner: postgres
--

ALTER TABLE forum_service.forums ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME forum_service.forums_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users; Type: TABLE; Schema: forum_service; Owner: postgres
--

CREATE TABLE forum_service.users (
    id integer NOT NULL,
    username text NOT NULL,
    interests text[]
);


ALTER TABLE forum_service.users OWNER TO postgres;

--
-- Name: users_id_seq1; Type: SEQUENCE; Schema: forum_service; Owner: postgres
--

ALTER TABLE forum_service.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME forum_service.users_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


-- Insert test data 
INSERT INTO forum_service.forums(
	name, "topickeyword", users)
	VALUES ('Політика в Україні', 'ukraine-politics', '{"user1", "user2"}'),
	('Футбол', 'football', '{}'),
	('Література', 'literature', '{"user2"}');
	
	INSERT INTO forum_service.users(
	username, interests)
	VALUES ('user1', '{"politics"}'),
	('user2', '{"ukraine-politics","politics"}');


--
-- Name: forums_id_seq; Type: SEQUENCE SET; Schema: forum_service; Owner: postgres
--

SELECT pg_catalog.setval('forum_service.forums_id_seq', 3, true);


--
-- Name: users_id_seq1; Type: SEQUENCE SET; Schema: forum_service; Owner: postgres
--

SELECT pg_catalog.setval('forum_service.users_id_seq1', 2, true);


--
-- Name: forums forumId; Type: CONSTRAINT; Schema: forum_service; Owner: postgres
--

ALTER TABLE ONLY forum_service.forums
    ADD CONSTRAINT "forumId" PRIMARY KEY (id);


--
-- Name: users userId; Type: CONSTRAINT; Schema: forum_service; Owner: postgres
--

ALTER TABLE ONLY forum_service.users
    ADD CONSTRAINT "userId" PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

