# loadtester

go mod init loadtester

Расширить диапазон доступных динамических портов до 50000.
Конечный максимальный 65535.

netsh int ipv4 set dynamicport tcp start=15535 num=50000