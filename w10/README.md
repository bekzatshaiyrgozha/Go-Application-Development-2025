1 pkg -> рассказать про экспортируемые пакет 

2 pkg/domain -> рассказать про домен

3 pkg/graceful -> пройтись про грейсфул 

4 pkg/middleware -> рассказать про WaitContextCancel 

5 pkg/reqresp -> пакет для запросов и ответов для сервиса

6 internal - > рассказать про пакет интернал, что он не экспортируется через gomod

7 cmd/app/main.go -> точка входа

8 internal/app/app.go -> пройтись по запуску 

9 internal/app/config

10 internal/app/connections

11 internal/app/start

12 internal/app/store

13 internal/data -> рассказать про DTO объекты (разница pkg/reqresp и internal/data)

14 internal/deliveries -> слой доставки, внутри разбиваем по доменам и технологиям

15 internal/services -> слой сервисов (контроллеров), внутри разбиваем по доменам и оборачиваем в декораторы

16 internal/repositories -> слой репозиториев, внутри разбиваем по доменам и технологиям

17 internal/usecases -> рассказать про юзкейс, почему важно описывать свой репозиторий