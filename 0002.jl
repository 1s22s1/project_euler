m = 1
n = 1
answer = 0

while n < 4000000
    if n % 2 == 0
        global answer +=n
    end

    global n, m = n + m, n
end

@show answer
