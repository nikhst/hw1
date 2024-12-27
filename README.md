# hw1
Ладно, я сдаюсь, curl меня победил(
Большими буквами я выделю те пункты, за которые можно получить хоть какие-то баллы (почему какие-то? чтобы узнать ответ на этот вопрос, читайте дальше!), чтобы было легче их выставлять.

ЧТО ЭТО ЗА ПРОЕКТ
Этот проект содержит в себе калькулятор rpn, который вычисляет выражения с помощью обратной польской нотации. Пользователь может отправить арифметическое выражение по HTTP и получить в ответ результат.
Программа выдает либо выражение, либо ошибку. 

ЕСТЬ РАЗБИЕНИЕ ПРОГРАММЫ НА ПАКЕТЫ
Программа состоит из файла запуска main, файла работы с сервером answer и файла с калькулятором rpn.

КАК ИМ ПОЛЬЗОВАТЬСЯ 
Сначала запускаем из папки с файлом main.go команду
go run main.go

Затем с помощью команды 
curl --location "localhost:8080/api/v1/calculate" --header 'Content-Type: application/json' --data "{\"expression\": \"2+2*2\"}"

или команды
curl -X POST -d 'input=\"2+2"' http://localhost:8080/v1/calculate
посылаем программе данные и радуемся результату)

ТЕСТЫ
А результату мы не радуемся, потому что что-то не работает. Если попробуете запустить программу с кодом, то она в случае первой команды выдаст ошибку Expression is not valid и разнообразные ошибки curl (URL rejected: bad hostname, еще один bad hostname одновременно с первым, 
Port number is not a decimal number between 0 and 65535, третий bad hostname, Could not resolve host, unmatched close brackets). Судя по гуглу, все эти ошибки значат проблему с экранированием выражения в data, но я перепробовала все советы оттуда и ничего не помогает, ни другие скобки, ни ставить \, 
ни отправлять команду в три строки, ни разным образом экранировать текст, ничего. Наверное, надо ставить линукс.

ПРИМЕР ИСПОЛЬЗОВАНИЯ

curl -X POST -d 'input="\2+2"' http://localhost:8080/api/v1/calculate
Expression is not valid 
(если заменить текст ошибки на другой текст, например, what, то будет показываться what, а это значит, что в expression попадает не исходная строка, а какое-то другое выражение)

curl -X POST -d 'input=\"2+2"\' http://localhost:8080/api/v1/calculate
curl: (2) no URL specified

curl --location "localhost:8080/api/v1/calculate" --header 'Content-Type: application/json' --data "{\"expression\": \"2+2*2\"}"
Expression is not valid
curl: (6) Could not resolve host: application

curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
curl: (3) URL rejected: Bad hostname
curl: (3) URL rejected: Bad hostname
curl: (3) URL rejected: Port number was not a decimal number between 0 and 65535
curl: (6) Could not resolve host: application
curl: (3) URL rejected: Bad hostname
curl: (3) unmatched brace in URL position 2:
'{
 ^


Как видите, в лучшем случае удалось добиться ответа Expression is not valid. Причем я пыталась вывести значение expression на экран, но оно тоже не выводится.
Expression is not valid выводится в любом случае (введена буква, цифра, выражение типа 2+2). Я думала, программа падает, но тогда бы log.Fatal показала бы ошибку по идее (она стоит в main именно для этого).
Так что код рабочий, но с ним что-то не так. Надеюсь, смогу еще подумать над ним в каникулы. Извините за задержку. Я пыталась что-то с ним сделать несколько дней подряд, но это бесполезно. Надеюсь, меня не отчислят. 
Надо наконец вступить в чат уже, а то я сначала не заметила ссылку, а потом было уже неловко(


