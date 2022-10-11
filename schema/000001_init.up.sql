CREATE TABLE "users" (
    "id" serial PRIMARY KEY NOT NULL,
    "name" varchar(200) NOT NULL,
    "key" varchar(100) NOT NULL,
    "secret" varchar(100) NOT NULL 
);

CREATE TABLE "books" (
    "id" serial PRIMARY KEY NOT NULL,
    "isbn" varchar(13) NOT NULL,
    "title" varchar(200) NOT NULL,
    "author" varchar(100) NOT NULL, 
    "published" smallint NOT NULL, 
    "pages" smallint NOT NULL, 
    "status" smallint NOT NULL 
);
