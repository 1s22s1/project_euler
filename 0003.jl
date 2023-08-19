number = 600851475143
prime_factors = []

while true
    if number == 1
        break
    end

    for l in 2:number
        if (number % l == 0)
            push!(prime_factors, l)
            global number = number / l
            break
        end
    end
end

@show Int(maximum(prime_factors))
