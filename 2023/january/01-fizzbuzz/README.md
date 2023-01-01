# January 01, 2023 - Fizzbuzz

For the first project I'm going to go for a classic project: Fizzbuzz. Fizzbuzz is a simple challenge in which you iterate through a range of numbers from 1 to *n* (inclusive) and print `Fizz` for numbers divisible by 3, `Buzz` for numbers divisible by 5, and `FizzBuzz` for numbers divisible by both 3 and 5. For example, if *n* is 15, the output would be:

```txt
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
```

For bonus points I'm going to make it extensible. A map of `int:str` will be able to be provided to the function which will change its behaviour, for example:

```go
Fizzbuzz(15, []mapping{
    {2, "Fizz"},
    {3, "Buzz"},
    {4, "Bazz"},
})
```

Would output:

```txt
1
Fizz
Buzz
FizzBazz
5
FizzBuzz
7
FizzBazz
Buzz
Fizz
11
FizzBuzzBazz
13
Fizz
Buzz
```
