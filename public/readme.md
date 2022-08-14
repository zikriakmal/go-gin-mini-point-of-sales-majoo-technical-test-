### Majoo Test By Zikri Akmal

#### Used Lib
-  GIN for http web framework
-  GORM for Object Relational Mapping Database
-  Testify for testing
-  Golang-jwt for create auth token 
-  Godotenv for .env file to store config
-  Faker to get fake data

#### Test Case 
**A**. Membuat fungsi login (5 point) <br/>
**B**. Untuk authorization pada point A gunakan JWT (6 point)<br/>
**C**. Laporan nama merchant, omzet per hari dalam pada bulan november mulai tanggal 1
sampai dengan tanggal 30 dengan pagination. Apabila tidak ada transaksi pada tanggal itu
omzet akan bernilai 0 (6 point)<br/>
**D**. API untuk menampilkan laporan nama merchant, nama outlet, omzet per hari pada bulan
november mulai tanggal 1 sampai dengan tanggal 30 dengan pagination. Apabila tidak ada
transaksi pada tanggal itu omzet akan bernilai 0 (6 point)<br/>
**E**. Pada poin C pastikan user tidak bisa melakukan akses pada merchant_id yang bukan
miliknya (10 point)<br/>
**F**. Pada poin D pastikan user tidak bisa melakukan akses laporan pada outlet_id yang bukan
miliknya (5 point)<br/>
**G**. Dari test case pada point C dan point D, apakah struktur ERD yang dibentuk sudah optimal
? berikan penjelasannya (9 point)<br/>
**H**. Dokumen teknis Data Manipulation Language (DML) (3 point)<br/>



Notes :
- No 1 is archive by .zip git archive
- No 1 has deployed on https://majoo.zikriakmal.my.id/
- No 1 downloadable documentation access on https://majoo.zikriakmal.my.id/
- No 1 has openapi documentation named “apispec.json” (openapi 3.0)
- to run no 1 cp .env.example to .env and set database configuration and set jwt_secret
- docker-compose up -d ( to run with docker compose)
- app has default seed set on .env true or false ( if seed setted true every app restart seed run and insert multiple times)
- default user login has seed with credential :
  - username : majoo
  - password : majoo123
