# Тема: поиск данных по XML

Это комбинированное задание по тому, как отправлять запросы, получать ответы, работать с параметрами и хедерами.

У нас есть какой-то поисковый сервис:

- SearchServer - своего рода внешняя система. Непосредственно занимается поиском данных в файле `dataset.xml`. В продакшене бы запускалась в виде отдельного веб-сервиса, но в вашем колде запустится как отдельный хендлер.

SearchServer принимает GET-параметры:

- `query` - что искать. Ищем по полям записи `Name` и `About` просто подстроку, без регулярок. `Name` - это first_name + last_name из xml (вам надо руками пройтись в цикле по записям и сделать такой, автоматом нельзя). Если поле пустое - то возвращаем все записи (поиск пустой подстроки всегда возвращает true), т.е. делаем только логику сортировки
- `order_field` - по какому полю сортировать. Работает по полям `Id`, `Age`, `Name`, если пустой - то сортируем по `Name`, если что-то другое - SearchServer ругается ошибкой.
- `order_by` - направление сортировки (как есть, по убыванию, по возрастанию)
- `limit` - сколько записей вернуть
- `offset` - начиня с какой записи вернуть (сколько пропустить с начала) - нужно для огранизации постраничной навигации

Дополнительно:

- Данные для работы лежат в файле `dataset.xml`

# Результаты:
Запрос: http://127.0.0.1:8080/?query=ipsum&order_field=Id&order_by=-1&limit=8&offset=2
![alt text](src/filter3.png)

Запрос: http://127.0.0.1:8080/?query=Lorem&order_field=Age&order_by=1&limit=8&offset=2
![alt text](src/filter1.png)

Запрос: http://127.0.0.1:8080/?query=ipsum&order_field=Age&order_by=-1&limit=8&offset=2
![alt text](src/filter2.png)
