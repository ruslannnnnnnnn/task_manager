my first app using golang

it's a simple rest api app that allows to execute CRUD operations over tasks

task has id, title, description, created_at fields

no frameworks

uses sqlite3 as a dbms

to run this app you need docker install on your machine
to run create .env file, copy content of env.example into .env
then run docker compose up --build