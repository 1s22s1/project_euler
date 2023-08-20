products = []

for i = 100:999
    for j = 100:999
        if string(i * j) == reverse(string(i * j))
            push!(products, i * j)
        end
    end
end

@show maximum(products)
