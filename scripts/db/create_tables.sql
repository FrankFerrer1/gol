-- create database
CREATE DATABASE assets;

-- connect to assets db
\c assets

-- create table
CREATE TABLE IF NOT EXISTS links (
    id serial PRIMARY KEY,
    assetname character varying(50) NOT NULL,
    url character varying(255) NOT NULL
);

INSERT INTO links (assetname, url)
VALUES
    ('Link 1', 'https://example.com/link1'),
    ('Link 2', 'https://example.com/link2'),
    ('Link 3', 'https://example.com/link3'),
    ('Link 4', 'https://example.com/link4'),
    ('Link 5', 'https://example.com/link5'),
    ('Link 6', 'https://example.com/link6'),
    ('Link 7', 'https://example.com/link7'),
    ('Link 8', 'https://example.com/link8'),
    ('Link 9', 'https://example.com/link9'),
    ('Link 10', 'https://example.com/link10');