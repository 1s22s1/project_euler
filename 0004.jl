products = []

for i = 100:999
    for j = 100:999
        product = i * j
        product_sring = string(product)

        if product_sring == reverse(product_sring)
            push!(products, product)
        end
    end
end

answer = maximum(products)
println(answer)