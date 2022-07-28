#!/bin/bash
psql -U admin -d mutantApp <<-EOSQL
CREATE TABLE IF NOT EXISTS dna (
    id VARCHAR(255) PRIMARY KEY UNIQUE,
    is_mutant BOOLEAN
);
CREATE TABLE IF NOT EXISTS stats (
    id VARCHAR(255) PRIMARY KEY UNIQUE,
    count_mutant_dna INTEGER DEFAULT 0,
    count_human_dna INTEGER DEFAULT 0
);
INSERT INTO stats(id) VALUES('main');
EOSQL