# Introduction - Yalnızca GO - 1 Videoda

[Youtube Link](https://youtu.be/AD2zUrhL6GI?si=_rSn1M_LhkSRKgKP)

[Github repo](https://github.com/mustafamuratcoskun/golang-product-app)


Merhabalar, bu videoda 1 Videoda Go programlama dilini detaylı bir şekilde öğreneceğiz. Channellar, Go routineler ve diğer dil özellikleriyle ilgili  detaylı bir bakış açısına sahip olacağız. Ardından, Docker ve PostgreSQL kullanarak bir REST API geliştirme sürecine adım atacağız. Her katmanı entegrasyon ve birim testleriyle güçlendirerek katmanlı mimariyi baştan sona inceleyeceğiz. Herkese iyi seyirler!

## Table of Contents - Yalnızca Go - 1 Videoda - Dakikalar  

- 00:00:00 Eğitime Giriş
- 00:00:52  Golang Kurulumu 
- 00:03:29 Hello World, Modüller ve Bağımlılıklar
- 00:36:30 Arrayler ve Slicelar
- 00:43:40 Golang Karar Yapıları (if-else conditions)
- 00:49:58 For Döngüleri
- 01:02:37 Maps Veri Tipi
- 01:08:37 For Each Döngüleri
- 01:15:54 Fonksiyonlar
- 01:30:53 Pointerlar
- 01:46:37 Pointerlar ve Arrayler
- 01:50:30 Structlar
- 02:08:30 Interfaceler
- 02:42:25 Defer Kullanımı ve Panic
- 02:55:36 Variable Scopes
- 03:03:28 Paketler ve Modüller
- 03:22:23 Function Literals
- 03:30:37 Error Handling
- 03:56:22 Go Routineler ile Asenkron İşlemler
- 04:33:05 Unbuffered Channellar
- 04:53:30 Buffered Channellar ve Select
- 05:21:53 Golang Context Package
- 05:38:38 Product App Projesi

Tutor : Mustafa Murat Coşkun

## Notes: Product App

```
go mod init product-app

go get github.com/labstack/echo/v4
go get github.com/jackc/pgx/v4
go get github.com/labstack/gommon
```

```
sh test_db.sh   // ../test/scripts
go test -v      // ../test/infrastructure
```