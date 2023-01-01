number = 600851475143
prime_factors = []

while true
    for l in 2:number
        if (number % l == 0)
            push!(prime_factors, Int(l))
            global number = number / l
            break
        end
    end

    if number == 1
        break
    end
end

answer = maximum(prime_factors)

println(answer)
