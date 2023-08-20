function main()
products = []

for i = 999:-1:100
    for j = 999:-1:100
        if string(i * j) == reverse(string(i * j))
            push!(products, i * j)
        end
    end
end

@show maximum(products)
end

main()
