Write-Host "`n--- health (200) ---"
curl.exe -s localhost:8080/health

Write-Host "`n--- create (201) ---"
$yeni = curl.exe -s -X POST localhost:8080/appointments -H "Content-Type: application/json" -d '{\"patient\":\"TestHasta\",\"phone\":\"5550001122\",\"date\":\"2026-12-05\",\"clinic_id\":1}' | ConvertFrom-Json
$id = $yeni.id
Write-Host "olusan id: $id"

Write-Host "`n--- create, bos hasta adi (400) ---"
curl.exe -s -X POST localhost:8080/appointments -H "Content-Type: application/json" -d '{\"patient\":\"\",\"phone\":\"555\",\"date\":\"2026-12-05\",\"clinic_id\":1}'

Write-Host "`n--- create, bozuk JSON (400) ---"
curl.exe -s -X POST localhost:8080/appointments -H "Content-Type: application/json" -d '{bozuk}'

Write-Host "`n--- list (200) ---"
curl.exe -s localhost:8080/appointments

Write-Host "`n--- get by id (200) ---"
curl.exe -s localhost:8080/appointments/$id

Write-Host "`n--- get, olmayan kayit (404) ---"
curl.exe -s -i localhost:8080/appointments/999999 | Select-Object -First 1

Write-Host "`n--- get, bozuk id (400) ---"
curl.exe -s -i localhost:8080/appointments/abc | Select-Object -First 1

Write-Host "`n--- delete (204) ---"
curl.exe -s -i -X DELETE localhost:8080/appointments/$id | Select-Object -First 1

Write-Host "`n--- delete, ayni id tekrar (404) ---"
curl.exe -s -i -X DELETE localhost:8080/appointments/$id | Select-Object -First 1