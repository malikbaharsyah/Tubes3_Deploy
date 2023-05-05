# Penerapan String Matching dan Regular Expression dalam Pembuatan ChatGPT Sederhana
> Tugas Besar III IF2211 Strategi Algoritma Semester II Tahun 2022/2023
.
> Live demo [_here_](https://www.youtube.com).

## Table of Contents
* [General Info](#general-information)
* [Technologies Used](#technologies-used)
* [Features](#features)
* [Setup](#setup)
* [Usage](#usage)
* [Project Status](#project-status)
* [Acknowledgements](#acknowledgements)
* [Authors](#authors)
<!-- * [License](#license) -->


## General Information
Sebuah aplikasi ChatGPT sederhana dengan mengaplikasikan pendekatan QA yang paling sederhana tersebut. Pencarian pertanyaan yang paling mirip dengan pertanyaan yang diberikan pengguna dilakukan dengan algoritma pencocokan string Knuth-Morris-Pratt (KMP) dan Boyer-Moore (BM).
<!-- You don't have to answer all the questions - just the ones relevant to your project. -->


## Technologies Used
- Backend -> Golang using Gin, Gorm, and Axios
- Frontend -> Next.js


## Features
- Fitur pertanyaan teks
- Fitur kalkulator
- Fitur tanggal
- Tambah pertanyaan dan jawaban ke database
- Hapus pertanyaan dari database


## Setup
- Pastikan sudah install dependencies (Golang, npm, IDE yang dipilih)


## Usage

1. Clone\unzip terlebih dahulu repository ini
2. Clone repo
   ```sh
   git clone https://github.com/archmans/Tubes3_13521010
   ```
3. Pastikan terminal berada pada directory program, jika belum:
   ```sh
   cd <path repo>
   cd src
   ```
4. Local Database setup
    ```sh
    pada setup.go di folder backend/model
    gorm.Open(mysql.Open("root:(isi dengan pasword)@tcp(localhost:3306)/gpt"))
    buka mariaDB pada cmd
    create database gpt;
   ```

5. Jalankan backend
   ```sh
   cd services
   cd backend
   go run main.go
   ```
6. Install npm packages
   ```sh
   npm install
   ```
7. Jalankan frontend
   ```js
   npm run dev
   ```
8. Jalankan web pada http://localhost:3000


## Project Status
- InsyaAllah complete


## Acknowledgements
- Project ini ditujukan untuk memenuhi Tugas Besar III IF2211 Strategi Algoritma
- Many thanks to chatGPT and our team


## Authors
   [13521010 - Muhamad Salman Hakim Alfarisi](https://github.com/archmans)
   
   [13521029 - M. Malik I. Baharsyah](https://github.com/malikbaharsyah)

   [13521030 - Jauza Lathifah Annassalafi](https://github.com/lostgirrlll)



<!-- Optional -->
<!-- ## License -->
<!-- This project is open source and available under the [... License](). -->

<!-- You don't have to include all sections - just the one's relevant to your project -->
