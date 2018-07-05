def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("DNA strands with different size")

    return len(
        list(filter(lambda x: x[0] != x[1], zip(strand_a, strand_b)))
        )