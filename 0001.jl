sequence = 1:999

filtered_sequence = filter(n -> ((n % 3) == 0 || (n % 5) == 0), sequence)
answer = sum(filtered_sequence)

println(answer)