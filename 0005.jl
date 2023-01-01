function gcd(m, n)
    while true
        if m < n
            m, n = n, m
        end

        if((rem(m, n) == 0))
            return n
        end

        m, n = n, rem(m, n)
    end
end

function lcm(m, n)
    m * n / gcd(m, n)
end

answer = 1

for i = 1:20
    global answer = lcm(answer, i)
end

println(Int(answer))