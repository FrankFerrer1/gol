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