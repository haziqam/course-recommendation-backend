## Deksripsi Program

Program ini merupakan aplikasi berbasis web yang digunakan untuk menemukan pilihan mata kuliah yang dapat menghasilkan IP maksimal. Terdapat tiga data utama yang disimpan pada program ini, yaitu:

- Fakultas, yang memiliki atribut:
  - Nama fakultas
- Jurusan, yang memiliki atribut:
  - Nama jurusan
  - Nama fakultas dari jurusan tersebut
- Matkul, yang memiliki atribut:
  - Nama matkul
  - Nama jurusan dari matkul tersebut
  - Minimum semester pengambilan
  - Jumlah SKS matkul
  - Prediksi indeks matkul

Dengan ketiga data tersebut, pengguna dapat menemukan pilihan mata kuliah yang dapat menghasilkan IP maksimal dengan menginput fakultas pengguna, semester pengguna saat ini, SKS minimal yang dapat diambil, serta SKS maksimal yang dapat diambil. Program akan menampilkan pilihan mata kuliah yang disarankan, IP yang dihasilkan, serta total SKS dari matkul-matkul tersebut. Selain itu, pengguna dapat menambah atau menghapus data entitas Fakultas, Jurusan, maupun Matkul.

## Teknologi yang Digunakan

- Go Fiber, sebagai framework untuk routing dan middleware (mirip express.js)
  github.com/gofiber/fiber/v2
- Air, untuk melakukan live reload ketika terjadi perubahan pada source code (mirip nodemon)
  github.com/cosmtrek/air@latest
- PostgreSQL, database management system untuk basis data relasional
- Docker, untuk container aplikasi

## Penjelasan Algoritma

Misal terdapat n mata kuliah yang disimpan pada array `mataKuliah` dengan indeks (0, 1, 2, ..., n - 1). Masing-masing mata kuliah memiliki data SKS yang disimpan pada array `SKS`, serta prediksi indeks yang disimpan pada array `prediksi` sehingga mata kuliah `i` memiliki SKS sebesar `SKS[i]` dan prediksi indeks sebesar `prediksi[i]`. Objektif dari program ini adalah menentukan matkul mana yang akan dipilih sehingga menghasilkan IP tertinggi, di mana

```
 IP = sum(prediksi[i] * SKS[i]) / sum(SKS[i]) for i = 0 to i = n - 1
```

dengan constraint

```
  minSKS <= sum(SKS_terpilih) <= maxSKS
```

Dalam menyelesaikan permasalahan ini, saya mendekomposisi masalah menjadi dua bagian, yaitu:

- Mencari seluruh kemungkinan pilihan matkul yang total SKS-nya memenuhi constraint.
- Mencari satu kemungkinan pilihan matkul yang menghasilkan IP tertinggi (dan SKS tertinggi) dari seluruh kemungkinan pilihan matkul yang memenuhi constraint

Untuk menyelesaikan subpersoalan yang pertama, digunakan Dynamic Programming. Persoalan ini mirip dengan persoalan Subset Sum, di mana terdapat suatu multiset berisi kumpulan bilangan, dan kita perlu menentukan suatu subset dari multiset tersebut yang jumlah anggotanya sama dengan K. Hal ini mirip dengan persoalan mencari kemungkinan pilihan matkul yang total SKS-nya sama dengan K, di mana array `SKS` dapat dianalogikan dengan multiset berisi bilangan pada persoalan Subset Sum. Untuk menemukan apakah terdapat sebuah subset mata kuliah yang total SKS-nya sama dengan K, kita dapat menyelesaikan persamaan rekursif berikut.

```
isSubsetAvailable(SKS, n, K) = isSubsetAvailable(SKS, n - 1, K) ||
                               isSubsetAvailable(SKS, n - 1, K - SKS[n - 1])

Keterangan: n adalah panjang array SKS
```

Terdapat dua kemungkinan yang dapat menyebabkan persamaan `isSubsetAvailable(SKS, n, K)` bernilai `true`

- Apabila matkul ke n - 1 tidak diambil, dan terdapat subset dari array `SKS[0 ... n - 2]` yang memiliki total SKS sebesar K

  ATAU

- Apabila matkul ke n - 1 diambil, dan terdapat subset dari array `SKS[0 ... n - 2]` yang memiliki total SKS sebesar K - SKS[n - 1]

Fungsi rekursif di atas dievaluasi hingga mencapai kasus basis, yaitu ketika `K = 0`, fungsi selalu bernilai `true`, karena untuk setiap array, pasti ada subset dari array tersebut yang jumlah anggotanya = 0 (subset tersebut adalah himpunan kosong). Basis lainnya adalah ketika `n = 0`, di mana fungsi selalu bernilai `false`, untuk K > 0, karena tidak mungkin mendapatkan total jumlah SKS > 0 dari array kosong.

Untuk menentukan apakah ada subset dari array `SKS` yang total SKS-nya memenuhi constraint, kita dapat melakukan iterasi nilai K dari minSKS hingga maxSKS dan mengevaluasi fungsi rekursif di atas. Pencarian dengan metode rekursif memiliki kompleksitas yang besar (eksponensial), oleh karena itu kita dapat menggunakan teknik Dynamic Programming, di mana kita melakukan pencatatan pada setiap tahap pemilihan mata kuliah. Dengan begitu, pada tahap ke-p, kita dapat melihat catatan tahap ke-(p - 1) yang telah kita buat sebelumnya.

Dalam pendekatan Dynamic Programming, kita perlu menyimpan tabel boolean `dpTable` berukuran `(n + 1) * (maxSKS + 1)`, di mana `dpTable[p][sisaSKS]` berisi jawaban apakah terdapat subset dari array `mataKuliah[0 ... p - 1]` sedemikian sehingga total SKS dari subset tersebut sama dengan `sisaSKS`. Seluruh elemen `dpTable` diinisialisasi dengan false terlebih dahulu, kemudian dilanjutkan dengan pengisian seluruh anggota dengan cara berikut.

- Lakukan pengisian basis terlebih dahulu, yaitu ketika `sisaSKS = 0`

```
dpTable[tahap][0] = true for tahap = 0 to tahap = n

Note:
untuk setiap tahap, selalu ada subset yang total SKS-nya sama dengan 0, yaitu himpunan kosong.
Oleh karena itu, dpTable[tahap][0] bernilai true.
```

- Lakukan pengisian basis lainnya, yaitu ketika `tahap = 0`, yaitu belum ada mata kuliah yang diperiksa hingga saat ini.

```
dpTable[0][sisaSKS] = false for sisaSKS = 0 to sisaSKS = n

Note: karena belum ada mata kuliah yang diperiksa (sejauh ini masih himpunan kosong),
maka tidak mungkin ada subset dari himpunan kosong yang menghasilkan total SKS > 0.
Oleh karena itu, dpTable[0][sisaSKS] bernilai false. Karena di awal kita sudah
inisialisasi seluruh elemen dengan false, kita tidak perlu melakukan pengisian
secara manual, sehingga pada program, tahap ini di skip
```

- Lakukan iterasi dari `tahap = 1` hingga `tahap = n`, dan untuk setiap tahap, lakukan iterasi dari `sisaSKS = 1` hingga `sisaSKS = maxSKS` dan lakukan pengisian tabel berdasarkan rumus berikut.

```
dpTable[tahap][sisaSKS] = dpTable[tahap - 1][sisaSKS] || dpTable[tahap - 1][sisaSKS - SKS[tahap - 1]]

Note:
dpTable[tahap - 1][sisaSKS] artinya mataKuliah[tahap - 1] ga kepilih,
jadi sisaSKS-nya ga berubah untuk tahap berikutnya.

dpTable[tahap - 1][sisaSKS - SKS[tahap - 1]] artinya mataKuliah[tahap - 1] kepilih,
jadi sisaSKS-nya berkurang untuk tahap berikutnya
```

- Setelah pengisian selesai, cek apakah `dpTable[n][K] = true`, untuk `minSKS <= K <= maxSKS`. Jika ya, maka terdapat subset sedemikian sehingga total SKS dari subset tersebut sama dengan K.

- Untuk setiap `dpTable[n][K]` yang bernilai true, tentukan apakah `MataKuliah[K - 1]` terpilih dengan cara mengecek:

  - Jika `dpTable[n - 1][K]` bernilai true, maka tanpa memilih mata kuliah pada tahap tersebut dapat menghasilkan subset yang total SKS nya sama dengan K.
  - Jika `dpTable[n - 1][K - SKS[K - 1]]` bernilai true, maka dengan memilih mata kuliah pada tahap tersebut dapat menghasilkan subset yang total SKS nya sama dengan K (perhatikan bahwa kedua kasus ini tidak mutually exclusive, sehingga untuk setiap tahap tetep bisa ada dua kemungkinan yang menghasilkan solusi)

- Lakukan pengecekan hingga baris ke-0 secara rekursif. Apabila pada baris tersebut sisa SKS nya mencapai 0, maka subset tersebut memenuhi constraint. Simpan seluruh subset yang memenuhi constraint.

Setelah seluruh subset yang memenuhi constraint dicatat, maka kita dapat melanjutkan ke subpersoalan kedua, yaitu mencari subset yang menghasilkan IP tertinggi dari seluruh subset yang memenuhi constraint. Hal ini dapat dilakukan dengan linear search, yaitu traversal setiap subset, kemudian bandingkan IP dari subset tersebut dan update IP tertinggi apabila menemukan IP yang lebih tinggi dibandingkan dnegan IP tertinggi sebelumnya.

## Analisis Algoritma

Berikut adalah kompleksitas waktu algoritma pencarian mata kuliah terbaik dengan pendekatan Dynamic Programming.

- Untuk melakukan pengisian `dpTable`, diperlukan `(n +1)(maxSKS + 1)` iterasi, sehingga kompleksitasnya adalah
  ```
  O(n * maxSKS)
  ```
- Untuk melakukan linear search untuk menemukan subset dengan IP tertinggi, maka kompleksitas waktu akan linear terhadap banyaknya subset yang memenuhi constraint. Misalkan banyaknya subset yang memenuhi constraint adalah v, maka kompleksitas linear search adalah
  ```
  O(v)
  ```
  Pada umumnya, nilai v akan jauh lebih sedikit dari 2^n (total seluruh subset yang dapat dibentuk) karena sebagian besar subset telah difilter pada tahap Dynamic Programming sebelumnya. Oleh karena itu, average case pendekatan ini jauh lebih baik daripada brute force atau pendekatan dengan fungsi rekursif. Namun, batas atas nilai v adalah 2^n.
- Sehingga, kompleksitas keseluruhan adalah
  ```
    O(max(n * maxSKS, v))
  ```

## Cara menjalankan aplikasi

- Clone repository ini (atau download zip) dan masuk ke dalam root directory

  ```
  git clone https://github.com/haziqam/course-scheduler-backend.git
  cd course_scheduler-backend
  ```

- Buat .env file untuk login database. dotenv file disimpan dalam root directory

  ```
  DB_HOST=db
  DB_USER= INSERT_YOUR_USERNAME
  DB_PASSWORD=INSERT_YOUR_PASSOWRD
  DB_NAME=coursedb
  ```

- Jalankan docker compose

  ```
  docker compose up -d
  ```

- Jalankan aplikasi frontend [Frontend repository](https://github.com/haziqam/course-scheduler-frontend)

## Referensi

- https://www.youtube.com/watch?v=34l1kTIQCIA (thanks mas mas india <3)
- https://www.youtube.com/watch?v=p08c0-99SyU&t=1020s
