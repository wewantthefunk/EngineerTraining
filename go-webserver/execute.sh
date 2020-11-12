#!/bin/bash
psql -U postgres << EOF
  CREATE DATABASE mytestdb;
  \c mytestdb
  CREATE TABLE IF NOT EXISTS public.mytable (id INT, username VARCHAR NOT NULL, password VARCHAR NOT NULL);
  INSERT INTO public.mytable VALUES (1, 'user 1', 'password1');
EOF
