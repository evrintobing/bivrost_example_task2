CREATE ROLE evrin WITH LOGIN SUPERUSER PASSWORD 'root';
CREATE DATABASE example;
GRANT ALL PRIVILEGES ON DATABASE example TO evrin;

\c example evrin

CREATE SEQUENCE IF NOT EXISTS items_id_seq;
CREATE SEQUENCE IF NOT EXISTS order_id_seq;

CREATE TABLE "public"."items" (
    "id" int4 NOT NULL DEFAULT nextval('items_id_seq'::regclass),
    "nama_produk" varchar(255),
    "deskripsi_produk" varchar(255),
    "harga" int4,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."orders" (
    "id_order" int4 NOT NULL DEFAULT nextval('order_id_seq'::regclass),
    "id_produk" int4,
    "jumlah_produk" int4,
    PRIMARY KEY ("id_order")
);


ALTER TABLE "public"."orders" ADD FOREIGN KEY ("id_produk") REFERENCES "public"."items"("id") ON DELETE CASCADE;

INSERT INTO "public"."items" ("id","nama_produk","deskripsi_produk","harga") VALUES ('1','Mac','Bagus','22000000');
INSERT INTO "public"."items" ("id","nama_produk","deskripsi_produk","harga") VALUES ('2','Kalender','Baru','10000');
INSERT INTO "public"."items" ("id","nama_produk","deskripsi_produk","harga") VALUES ('3','Mac','komik','22000');

