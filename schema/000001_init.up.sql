CREATE TABLE "users" (
    "id" serial PRIMARY KEY NOT NULL,
    "name" varchar(200) NOT NULL,
    "key" varchar(100) NOT NULL unique,
    "secret" varchar(100) NOT NULL 
);

CREATE TABLE "shelf" (
    "id" serial PRIMARY KEY NOT NULL,
    "isbn" varchar(13) NOT NULL,
    "status" smallint NOT NULL,
    "user_key" varchar(100) NOT NULL references "users"("key"), 
    UNIQUE("isbn", "user_key")
);
