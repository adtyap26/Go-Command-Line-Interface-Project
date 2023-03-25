#  Jadwal Imsakiyah CLI

Program command line interface (CLI) yang menyajikan informasi jadwal-imsakiyah di seluruh wilayah Indonesia.

## Instalasi

Clone Repositori ini dan jalankan perintah berikut:
```Golang
go build .

```

## Penggunaan
Jalankan binary file yang telah dibuat dengan opsi setelah perintah `cari` :

```Golang
jadwal-imsakiyah cari --provinsi <nama-provinsi> --kabKota <nama-kabupaten/kota>

```
Jalankan opsi `--help` untuk manual.


## Pustaka yang digunakan
This app uses the following third-party libraries:

- [goquery](https://github.com/PuerkitoBio/goquery) 
- [simpletable](https://github.com/alexeyco/simpletable) 
- [cobra](https://github.com/spf13/cobra) 


## Credits
Program ini menggunakan website [viva.co.id](https://www.viva.co.id/jadwal-imsakiyah/) sebagai tempat untuk scraping data.
