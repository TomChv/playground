# Use the output builtin.
%builtins output

# Import the serialize_word() function.
from starkware.cairo.common.alloc import alloc
from starkware.cairo.common.serialize import serialize_word

func array_sum(arr : felt*, size) -> (sum):
    if size == 0:
        return (sum=0)
    end

    # size is not zero.
    let (sum_of_rest) = array_sum(arr=arr + 1, size=size - 1)
    return (sum=[arr] + sum_of_rest)
end

func array_product(arr : felt*, size) -> (product):
    if size == 0:
        return (product=1)
    end

    # size is not zero.
    let (product_of_rest) = array_product(arr=arr + 1, size=size - 1)
    return (product=[arr] * product_of_rest)
end

func main{output_ptr : felt*}():
    const ARRAY_SIZE = 3

    let (ptr1) = alloc()
    let (ptr2) = alloc()

    assert [ptr1] = 9
    assert [ptr1 + 1] = 16
    assert [ptr1 + 2] = 25

    assert [ptr2] = 9
    assert [ptr2 + 1] = 16
    assert [ptr2 + 2] = 25

    # let (sum) = array_sum(arr=ptr1, size=ARRAY_SIZE)
    let (product) = array_product(arr=ptr2, size=ARRAY_SIZE)

    serialize_word(product)

    return ()
end
