m = 1
n = 1
sequence = []

while n < 4000000
    push!(sequence, n)

    swap = n
    global n = n + m
    global m = swap
end

filterd_sequence = filter(iseven , sequence)
answer = sum(filterd_sequence)

println(answer)