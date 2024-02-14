# Решение задачи ZIP


Достать hash из архива 
```bash
zip2john Flag.zip > hash 
```

Брут hash 
```bash
john --wordlist=rockyou.txt  hash
```

пароль от архива `linuxx`

Флаг `StepikCTF{Brutus,-but-not-Caesar's.}`


Подробнее про инструмент можно посмотреть тут https://www.kali.org/tools/john/
