db.Query -> sorguyu gönderir , satırları okumaya hazır bir rows döner 

defer rows.close() -> bağlantıyı havuza geri verir , yazmazsan  havuz tükenir.

rows.next() + rows.Scan(...) -> satır satır ilerle , sütunları değişkenlere yaz.

rows.Err() -> dongü bittikten sonra kontrol , ağ koparsa buradan anlarsın.



kullanılmayan eski değişkenleri sil; Encode(appointments) ile Encode(list) karışması bir saat kaybettirdi