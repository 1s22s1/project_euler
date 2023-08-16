answer = sum(filter(n -> ((n % 3) == 0 || (n % 5) == 0), 1:999))

@show answer
