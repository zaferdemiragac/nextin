# NextIn

Go ile yazılmış, randevu ve lead takibi için basit bir REST API. Klinikler ve randevuyla çalışan küçük işletmeler için hasta kaydı ekleme, listeleme, tek kayıt getirme ve silme işlemlerini sunar.

Veriler şu an uygulama belleğinde tutulur; sunucu kapandığında kayıtlar silinir. Kalıcı depolama (PostgreSQL) geliştirme planında.

## Gereksinimler

- Go 1.22 veya üstü (yol kalıplarında metot eşleştirme bu sürümle geldi)

## Çalıştırma

```bash
git clone https://github.com/zaferdemiragac/nextin.git
cd nextin
go run .
```

Sunucu `http://localhost:8080` adresinde çalışır.

## Proje yapısı

```
.
├── main.go       sunucu kurulumu ve yol kayıtları
├── handlers.go   HTTP isteklerini karşılayan fonksiyonlar
├── store.go      bellekteki randevu listesi ve eşzamanlılık kilidi
├── models.go     Appointment veri tipi
└── go.mod
```

## Veri modeli

```json
{
  "id": 1,
  "patient": "Ali Yılmaz",
  "phone": "5551234567",
  "date": "2026-11-10",
  "status": "beklemede"
}
```

`id` alanını sunucu üretir; istek gövdesinde gönderilmesine gerek yoktur.

## Endpoint'ler

| Metot | Yol | Açıklama | Durum kodları |
|---|---|---|---|
| GET | `/health` | Sunucunun ayakta olduğunu bildirir | 200 |
| GET | `/appointments` | Tüm randevuları listeler | 200 |
| POST | `/appointments` | Yeni randevu ekler, eklenen kaydı döner | 201, 400 |
| GET | `/appointments/{id}` | Tek bir randevuyu getirir | 200, 400, 404 |
| DELETE | `/appointments/{id}` | Randevuyu siler | 204, 400, 404 |

Hata durumları:

- **400** — istek hatalı: gövde geçerli JSON değil, veya `{id}` bir sayı değil
- **404** — istek geçerli ama o `id`'ye sahip bir kayıt yok

## Örnek istekler

Sunucu durumu:

```bash
curl localhost:8080/health
```

Yeni randevu ekleme:

```bash
curl -X POST localhost:8080/appointments \
  -H "Content-Type: application/json" \
  -d '{"patient":"Ali","phone":"555","date":"2026-11-10","status":"beklemede"}'
```

Tüm randevuları listeleme:

```bash
curl localhost:8080/appointments
```

Tek randevu getirme:

```bash
curl localhost:8080/appointments/1
```

Randevu silme:

```bash
curl -X DELETE localhost:8080/appointments/1
```

## Teknik notlar

Paylaşılan randevu listesine erişim `sync.Mutex` ile korunur. Her HTTP isteği ayrı bir goroutine'de çalıştığı için, kilit olmadan aynı anda gelen istekler listeyi bozabilir veya aynı `id`'yi iki kayda verebilir.

## Yol haritası

- [ ] PostgreSQL ile kalıcı depolama
- [ ] Kullanıcı kaydı ve