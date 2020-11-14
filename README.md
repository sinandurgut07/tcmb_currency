# TCMB Güncel Döviz Kurları 

TCMB webservisini (http://www.tcmb.gov.tr/kurlar/today.xml) kullanarak güncel döviz kurlarını alır ve JSON formatında servise çevirir. Olası bir veri tutarsızlığı durumunda TCMB verilerini dikkate alın! 

## Kurulum
```
git clone git://github.com/sinandurgut07/tcmb_currency
cd tcmb_currency
go run main.go
```
## Kullanım
```
Tüm para birimleri:
http://localhost:3000

Tekil para birimi:
http://localhost:3000/currency/{currency}
Ör:
http://localhost:3000/currency/USD

```
