# BeeShifts-Server
## Description
**BeeShifts** — система для автоматизации планирования графиков сменности сотрудников, предназначенная для организаций, чтобы упростить создание и управление рабочими сменами. Система позволит менеджерам создавать шаблоны смен и выбирать различные временные промежутки для составления графика, а сотрудники смогут указывать свои предпочтения по поводу смен в выбранном графике (возможность/желание выйти на выбранную смену), после этого в системе будет задействован алгоритм, который на основе предпочтений сотрудников и норм трудового кодекса распределяет сотрудников по сменам, что в свою очередь минимизирует конфликты.

Данный проект - серверная часть этой системы.

Алгоритм, используемый для распределения сотрудников по сменам, в данный момент разрабатывается мной самостоятельно, как выпускная квалификационная работа.

Архитектура кода - **слоистая**. То есть в центре — бизнес-логика со всеми её сущностями, которые занимаются прикладными задачами. К бизнес-логике относятся `use cases` (код, который выполняет какой-либо бизнес-процесс, в качестве действий использует атомарные методы сервисов), `services` (группа методов, которые группируются в сервисы по смысловой нагрузке и необходимы для изоляции сценария от внешних зависимостей) и `entities` (являются представлением данных из БД). Вокруг бизнес-логики - драйверы, которые могут использоваться как для вызова бизнес-логики, так и для предоставления данных для бизнес-логики (это `handlers` и `repositories`).

Реализованные на данный момент группы эндпоинтов - `users`, `organizations` и `positions`. В процессе реализация следующих - `shifts_type`, `plans`, `schedule`, `preferences`, `shifts` (смены заполняются автоматически алгоритмом, но может возникнуть необходимость в ручной корректировке - поэтому будет такой эндпоинт).

В тех эндпоинтах, где в этом есть необходимость - повешены middlewares авторизации и аутентификации (на основе jwt-токена).

На данный момент реализован `restfull api` с использованием фреймворка `gin`. В перспективе - выделение алгоритма распределения сотрудников по сменам в отдельный сервис, а общение с ним реализовать через grpc.

Используемая база данных - `postgres`, но благодаря реализованной архитектуре, можно легко перейти на другую.

## API docs
Интерактивная swagger документация актуальной версии сервиса - https://api.beeshifts.tech/docs/index.html